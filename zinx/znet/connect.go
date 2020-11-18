package znet

import (
    "errors"
    "io"
    "log"
    "net"
    "zinx-learn/zinx/ziface"
)

// IConnection implement
type Connection struct {
    // 当前链接的套接字
    Conn *net.TCPConn
    // 链接的 ID
    ConnID uint32
    // 告知当前链接是否退出
    ExitChan chan bool
    // 当前链接 MsgHandler
    MsgHandler ziface.IMsgHandler

    // 当前的链接是否关闭
    isClosed bool
}

func NewConnection(conn *net.TCPConn, connID uint32, handler ziface.IMsgHandler) *Connection {
    return &Connection{
        Conn:     conn,
        ConnID:   connID,
        isClosed: false,
        ExitChan: make(chan bool, 1),
        MsgHandler:  handler,
    }
}

func (c *Connection) StartReader() {
    log.Println("Reader Goroutine is running ...")
    defer log.Printf("[ConnID = %d] Reader is exit, Remote addr is %s\n", c.ConnID, c.RemoteAddr().String())
    defer c.Stop()
    // Message 拆包
    for {
        dp := NewDataPack()
        // 读取客户端的 message Head 8 bytes
        head := make([]byte, dp.GetHeadLen())
        if _, err := io.ReadFull(c.GetTCPConn(), head); err != nil {
            log.Printf("Read Message Head error: %v\n", err)
            break
        }
        // Message 拆包获取 id 和 len
        msg, err := dp.UnPack(head)
        if err != nil {
            log.Printf("UnPack Message Head error: %v\n", err)
            break
        }
        // 根据 len 读取 message body
        if msg.GetMsgLen() > 0 {
            body := make([]byte, msg.GetMsgLen())
            if _, err := io.ReadFull(c.GetTCPConn(), body); err != nil {
                log.Printf("Read Message Body error: %v\n", err)
                break
            }
            msg.SetMsgData(body)
        }
        // 封装 Message 到 Request 中
        req := &Request{
            conn: c,
            msg:  msg,
        }
        // 将 Request 发送到 Router Handler 中处理业务
        go c.MsgHandler.DoMsgHandle(req)
    }
}

func (c *Connection) Start() {
    log.Printf("Conn Start() -- ConnID = %d\n", c.ConnID)
    go c.StartReader()
}

func (c *Connection) Stop() {
    log.Printf("Conn Stop() -- ConnID = %d\n", c.ConnID)

    if c.isClosed {
        return
    }
    c.isClosed = true

    _ = c.Conn.Close()
    close(c.ExitChan)
}

func (c *Connection) GetTCPConn() *net.TCPConn {
    return c.Conn
}

func (c *Connection) GetConnID() uint32 {
    return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
    return c.Conn.RemoteAddr()
}

func (c *Connection) SendMsg(id uint32, data []byte) error {
    if c.isClosed {
        return errors.New("Connection Closed When Send Msg ")
    }
    // 将 data 封包
    dp := NewDataPack()

    send, err := dp.Pack(NewMessage(id, data))
    if err != nil {
        log.Printf("Pack Message error: %v\n", err)
    }
    if _, err = c.Conn.Write(send); err != nil {
        log.Printf("Write Message(id = %d) error: %v\n", id, err)
        return err
    }
    return nil
}
