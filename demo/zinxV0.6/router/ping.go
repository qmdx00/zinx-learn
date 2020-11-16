package router

import (
    "log"
    "qmdx00.cn/zinx/ziface"
    "qmdx00.cn/zinx/znet"
)

type PingRouter struct {
    znet.BaseRouter
}

func (p *PingRouter) PreHandler(req ziface.IRequest) {

}

func (p *PingRouter) Handler(req ziface.IRequest) {
    log.Println("Call Ping Router Handler")
    log.Printf("Received from Client: { msgId = %d, data = %s }\n", req.GetMsgId(), string(req.GetData()))

    _ = req.GetConn().SendMsg(1, []byte("ping ...ping ...ping ..."))
}

func (p *PingRouter) PostHandler(req ziface.IRequest) {

}
