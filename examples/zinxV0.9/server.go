package main

import (
    "log"
    "zinx-learn/examples/zinxV0.7/router"
    "zinx-learn/zinx/ziface"
    "zinx-learn/zinx/znet"
)

func main() {
    s := znet.NewServer("[zinx v0.9]")

    s.SetOnConnStart(func(conn ziface.IConnection) {
        log.Println("do something after connection establish")
        _ = conn.SendMsg(200, []byte("do after connection begin"))
    })
    s.SetOnConnStop(func(conn ziface.IConnection) {
        log.Println("do something before connection destroy")
        _ = conn.SendMsg(400, []byte("do before connection lost"))
    })
    
    s.AddRouter(1, &router.PingRouter{})
    s.AddRouter(2, &router.HelloRouter{})
    
    s.Serve()
}
