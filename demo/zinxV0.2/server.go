package main

import (
    "qmdx00.cn/zinx/znet"
)

func main() {
    s := znet.NewServer("[zinx v0.2]")
    s.Serve()
}
