package main

import (
    "log"
    "zinx-learn/zinx/ziface"
    "zinx-learn/zinx/znet"
)

type PingRouter struct {
    znet.BaseRouter
}

func (p *PingRouter) Handler(req ziface.IRequest) {
    log.Println("Call Router Handler")
    log.Printf("Received from Client: { msgId = %d, data = %s }\n", req.GetMsgId(), string(req.GetData()))

    _ = req.GetConn().SendMsg(1, []byte("ping ...ping ...ping ..."))
}

func main() {
    s := znet.NewServer("[zinx v0.5]")
    s.AddRouter(1, &PingRouter{})
    s.Serve()
}
