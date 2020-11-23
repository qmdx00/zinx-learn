package znet

import (
    "encoding/json"
    "fmt"
    "log"
    "net"
    "zinx-learn/zinx/utils"
    "zinx-learn/zinx/ziface"
)

// IServer implement
type Server struct {
    Name        string
    IPVersion   string
    IP          string
    Port        uint
    MsgHandler  ziface.IMsgHandler
    ConnManager ziface.IConnManager
    // 链接创建和停止的 Hook 函数
    OnConnStart func(conn ziface.IConnection)
    OnConnStop  func(conn ziface.IConnection)
}

func NewServer(name string) ziface.IServer {
    s := &Server{
        Name:        utils.GlobalObject.Name,
        IPVersion:   "tcp4",
        IP:          utils.GlobalObject.Host,
        Port:        utils.GlobalObject.TcpPort,
        MsgHandler:  NewMsgHandler(),
        ConnManager: NewConnManager(),
    }
    return s
}

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
    s.MsgHandler.AddRouter(msgId, router)
    log.Println("Add Router Succeed ...")
}

func (s *Server) Start() {
    jsons, _ := json.Marshal(utils.GlobalObject)
    log.Printf("[Zinx global config] %v\n", string(jsons))

    log.Printf("[Start] Server listener at addr: %s:%d ...\n", s.IP, s.Port)

    go func() {
        s.MsgHandler.StartWorkerPool()
        // 创建socket套接字
        addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
        if err != nil {
            log.Println("resolve tcp addr error:", err)
            return
        }

        listener, err := net.ListenTCP(s.IPVersion, addr)
        if err != nil {
            log.Printf("listen %s error: %v\n", s.IPVersion, err)
            return
        }
        log.Printf("start zinx server [%s] succeed, listening ...\n", s.Name)

        var cid uint32 = 0
        // 监听TCP连接
        for {
            conn, err := listener.AcceptTCP()
            if err != nil {
                log.Printf("Accept error: %v\n", err)
                continue
            }

            if s.ConnManager.Size() > utils.GlobalObject.MaxConn {
                log.Printf("Too Many Connections, Max Connections = %d\n", utils.GlobalObject.MaxConn)
                _ = conn.Close()
                continue
            }
            dealConn := NewConnection(s, conn, cid, s.MsgHandler)
            cid++
            go dealConn.Start()
        }
    }()
}

func (s *Server) Stop() {
    log.Printf("[Stop] Zinx Server %s\n", s.Name)
    s.ConnManager.Clear()
}

func (s *Server) Serve() {
    s.Start()
    select {}
}

func (s *Server) GetConnManager() ziface.IConnManager {
    return s.ConnManager
}

func (s *Server) SetOnConnStart(hook func(conn ziface.IConnection)) {
    s.OnConnStart = hook
}

func (s *Server) SetOnConnStop(hook func(conn ziface.IConnection)) {
    s.OnConnStop = hook
}

func (s *Server) CallOnConnStart(conn ziface.IConnection) {
    if s.OnConnStart != nil {
        log.Println("Call OnConnStart Func")
        s.OnConnStart(conn)
    }
}

func (s *Server) CallOnConnStop(conn ziface.IConnection) {
    if s.OnConnStop != nil {
        log.Println("Call OnConnStop Func")
        s.OnConnStart(conn)
    }
}
