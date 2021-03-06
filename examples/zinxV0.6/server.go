package main

import (
    "zinx-learn/examples/zinxV0.6/router"
    "zinx-learn/zinx/znet"
)

func main() {
    s := znet.NewServer("[zinx v0.6]")
    s.AddRouter(1, &router.PingRouter{})
    s.AddRouter(2, &router.HelloRouter{})
    s.Serve()
}
