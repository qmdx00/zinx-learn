package main

import (
    "zinx-learn/zinx/znet"
)

func main() {
    s := znet.NewServer("[zinx v0.2]")
    s.Serve()
}
