package main

import (
    "zinx-learn/examples/zinxV0.7/router"
    "zinx-learn/zinx/znet"
)

func main() {
    s := znet.NewServer("[zinx v0.7]")
    s.AddRouter(1, &router.PingRouter{})
    s.AddRouter(2, &router.HelloRouter{})
    s.Serve()
}
