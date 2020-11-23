package router

import (
    "log"
    "zinx-learn/zinx/ziface"
    "zinx-learn/zinx/znet"
)

type HelloRouter struct {
    znet.BaseRouter
}

func (p *HelloRouter) PreHandler(req ziface.IRequest) {

}

func (p *HelloRouter) Handler(req ziface.IRequest) {
    log.Println("Call Hello Router Handler")
    log.Printf("Received from Client: { msgId = %d, data = %s }\n", req.GetMsgId(), string(req.GetData()))

    _ = req.GetConn().SendMsg(1, []byte("hello zinx v0.7"))
}

func (p *HelloRouter) PostHandler(req ziface.IRequest) {

}
