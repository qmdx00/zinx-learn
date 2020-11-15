package main

import (
    "log"
    "qmdx00.cn/zinx/ziface"
    "qmdx00.cn/zinx/znet"
)

type PingRouter struct {
    znet.BaseRouter
}

func (p *PingRouter) PreHandler(request ziface.IRequest) {
    log.Println("Call Router PreHandler")
    _, err := request.GetConn().GetTCPConn().Write([]byte("before ping\n"))
    if err != nil {
        log.Printf("call back before ping error: %v\n", err)
    }
}
func (p *PingRouter) Handler(request ziface.IRequest) {
    log.Println("Call Router Handler")
    _, err := request.GetConn().GetTCPConn().Write([]byte("ping ...ping ... ping ...\n"))
    if err != nil {
        log.Printf("call back ping error: %v\n", err)
    }
}
func (p *PingRouter) PostHandler(request ziface.IRequest) {
    log.Println("Call Router PostHandler")
    _, err := request.GetConn().GetTCPConn().Write([]byte("after ping\n"))
    if err != nil {
        log.Printf("call back after ping error: %v\n", err)
    }
}

func main() {
    s := znet.NewServer("[zinx v0.4]")
    s.AddRouter(&PingRouter{})
    s.Serve()
}
