package main

import (
    "fmt"
    "log"
    "zinx-learn/zinx/ziface"
    "zinx-learn/zinx/znet"
)

type PingRouter struct {
    znet.BaseRouter
}

func (p *PingRouter) PreHandler(request ziface.IRequest) {
    fmt.Println("Call Router PreHandler")
    _, err := request.GetConn().GetTCPConn().Write([]byte("before ping\n"))
    if err != nil {
        log.Printf("call back before ping error: %v\n", err)
    }
}
func (p *PingRouter) Handler(request ziface.IRequest) {
    fmt.Println("Call Router Handler")
    _, err := request.GetConn().GetTCPConn().Write([]byte("ping ...ping ... ping ...\n"))
    if err != nil {
        log.Printf("call back ping error: %v\n", err)
    }
}
func (p *PingRouter) PostHandler(request ziface.IRequest) {
    fmt.Println("Call Router PostHandler")
    _, err := request.GetConn().GetTCPConn().Write([]byte("after ping\n"))
    if err != nil {
        log.Printf("call back after ping error: %v\n", err)
    }
}

func main() {
    s := znet.NewServer("[zinx v0.3]")
    s.AddRouter(1, &PingRouter{})
    s.Serve()
}
