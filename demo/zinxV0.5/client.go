package main

import (
    "io"
    "log"
    "net"
    "zinx-learn/zinx/znet"
    "time"
)

func main() {
    conn, err := net.Dial("tcp", "127.0.0.1:8888")
    if err != nil {
        log.Println("client start error:", err)
        return
    }

    for {
        // 发送数据到服务端
        dp := znet.NewDataPack()
        bin, err := dp.Pack(znet.NewMessage(0, []byte("zinx v0.5 client test message")))
        if err != nil {
            log.Printf("Pack message error: %v\n", err)
            return
        }
        if _, err = conn.Write(bin); err != nil {
            log.Printf("Write to server error: %v\n", err)
        }

        // 接收服务端的消息
        binHead := make([]byte, dp.GetHeadLen())
        if _, err := io.ReadFull(conn, binHead); err != nil {
            log.Printf("Read message head error: %v\n", err)
            return
        }
        msg, err := dp.UnPack(binHead)
        if err != nil {
            log.Printf("UnPack message head error: %v\n", err)
            return
        }
        if msg.GetMsgLen() > 0 {
            body := make([]byte, msg.GetMsgLen())
            if _, err := io.ReadFull(conn, body); err != nil {
                log.Printf("Read Message Body error: %v\n", err)
                return
            }
            msg.SetMsgData(body)
        }
        log.Printf("----> Received Message: { id = %d, len = %d, data = %s }\n", msg.GetMsgId(), msg.GetMsgLen(), string(msg.GetMsgData()))

        time.Sleep(time.Millisecond)
    }
}
