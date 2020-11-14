package znet

import (
    "fmt"
    "log"
    "net"
    "qmdx00.cn/zinx/ziface"
)

// IServer implement
type Server struct {
    Name      string
    IPVersion string
    IP        string
    Port      int
    Router    ziface.IRouter
}

func (s *Server) AddRouter(router ziface.IRouter) {
    s.Router = router
    log.Println("Add Router Succeed ...")
}

func (s *Server) Start() {

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
            log.Printf("listen %s error: %v", s.IPVersion, err)
            return
        }
        log.Printf("start zinx server %s succeed, listening ...", s.Name)

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
    return &Server{
        Name:      name,
        IPVersion: "tcp4",
        IP:        "0.0.0.0",
        Port:      2333,
        Router:    nil,
    }
}
