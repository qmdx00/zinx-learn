package znet

import (
    "errors"
    "fmt"
    "net"
    "qmdx00.cn/zinx/ziface"
)

// IServer implement
type Server struct {
    Name      string
    IPVersion string
    IP        string
    Port      int
}

// 客户端绑定的链接处理的业务
func callBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
    fmt.Println("[Conn Handle] callBackToClient ...")
    if _, err := conn.Write(data[:cnt]); err != nil {
        fmt.Printf("write to client error: %v", err)
        return errors.New("BackToClient Error")
    }
    return nil
}

func (s *Server) Start() {

    fmt.Printf("[Start] Server listener at addr: %s:%d ...\n", s.IP, s.Port)

    go func() {
        // 创建socket套接字
        addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
        if err != nil {
            fmt.Println("resolve tcp addr error:", err)
            return
        }

        listener, err := net.ListenTCP(s.IPVersion, addr)
        if err != nil {
            fmt.Printf("listen %s error: %v", s.IPVersion, err)
            return
        }
        fmt.Printf("start zinx server %s succeed, listening ...", s.Name)

        cid := 0
        // 监听TCP连接
        for {
            conn, err := listener.AcceptTCP()
            if err != nil {
                fmt.Printf("Accept error: %v\n", err)
                continue
            }
            
            dealConn := NewConnection(conn, 1, callBackToClient)
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
    }
}
