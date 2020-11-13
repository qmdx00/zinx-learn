package main

import (
    "fmt"
    "net"
    "time"
)

func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:2333")
    if err != nil {
        fmt.Println("client start error:", err)
        return
    }

    for {
        _, err := conn.Write([]byte("hello server zinx v0.1"))
        if err != nil {
            fmt.Println("write to server error:", err)
            return
        }

        buf := make([]byte, 512)
        cnt, err := conn.Read(buf)
        if err != nil {
            fmt.Println("read from server error:", err)
            return
        }
        fmt.Printf("[server callback]: %s\n", string(buf[:cnt]))

        time.Sleep(time.Second)
    }

}
