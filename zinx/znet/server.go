package znet

import (
    "encoding/json"
    "fmt"
    "log"
    "net"
    "qmdx00.cn/zinx/utils"
    "qmdx00.cn/zinx/ziface"
)

// IServer implement
type Server struct {
    Name      string
    IPVersion string
    IP        string
    Port      uint
    Router    ziface.IRouter
}

func (s *Server) AddRouter(router ziface.IRouter) {
    s.Router = router
    log.Println("Add Router Succeed ...")
}

func (s *Server) Start() {
    jsons, _ := json.Marshal(utils.GlobalObject)
    log.Printf("[Zinx global config] %v\n", string(jsons))

    log.Printf("[Start] Server listener at addr: %s:%d ...\n", s.IP, s.Port)

    go func() {
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

        cid := 0
        // 监听TCP连接
        for {
            conn, err := listener.AcceptTCP()
            if err != nil {
                log.Printf("Accept error: %v\n", err)
                continue
            }

            dealConn := NewConnection(conn, 1, s.Router)
            cid++
            go dealConn.Start()
        }
    }()
}

func (s *Server) Stop() {

}

func (s *Server) Serve() {
    s.Start()
    select {}
}

func NewServer(name string) ziface.IServer {
    s := &Server{
        Name:      utils.GlobalObject.Name,
        IPVersion: "tcp4",
        IP:        utils.GlobalObject.Host,
        Port:      utils.GlobalObject.TcpPort,
        Router:    nil,
    }
    return s
}
