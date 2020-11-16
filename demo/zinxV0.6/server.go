package main

import (
    "qmdx00.cn/demo/zinxV0.6/router"
    "qmdx00.cn/zinx/znet"
)

func main() {
    s := znet.NewServer("[zinx v0.5]")
    s.AddRouter(1, &router.PingRouter{})
    s.AddRouter(2, &router.HelloRouter{})
    s.Serve()
}
