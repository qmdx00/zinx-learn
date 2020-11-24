package main

import (
    "log"
    "zinx-learn/examples/zinxV0.7/router"
    "zinx-learn/zinx/ziface"
    "zinx-learn/zinx/znet"
)

func main() {
    s := znet.NewServer("[zinx v1.0]")

    s.SetOnConnStart(func(conn ziface.IConnection) {
        log.Println("================ after conn start ===============>")
        conn.SetProperty("name", "zinx v1.0")
        conn.SetProperty("github", "https://github.com/qmdx00")
    })
    s.SetOnConnStop(func(conn ziface.IConnection) {
        log.Println("================ before conn stop ===============>")
        if name, err := conn.GetProperty("name"); err == nil {
            log.Printf("[Property] name = %v\n", name)
        }
        if url, err := conn.GetProperty("github"); err == nil {
            log.Printf("[Property] github = %v\n", url)
        }
    })

    s.AddRouter(1, &router.PingRouter{})
    s.AddRouter(2, &router.HelloRouter{})

    s.Serve()
}
