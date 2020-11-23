package znet

import (
    "errors"
    "io"
    "log"
    "net"
    "zinx-learn/zinx/utils"
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
    // 读写goroutine之间的通信
    msgChan chan []byte
}

func NewConnection(conn *net.TCPConn, connID uint32, handler ziface.IMsgHandler) *Connection {
    return &Connection{
        Conn:       conn,
        ConnID:     connID,
        isClosed:   false,
        ExitChan:   make(chan bool, 1),
        MsgHandler: handler,
        msgChan:    make(chan []byte),
    }
}

func (c *Connection) StartReader() {
    log.Println("[Reader Goroutine is running ...]")
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
        if utils.GlobalObject.WorkerPoolSize > 0 {
            c.MsgHandler.SendMsgToTaskQueue(req)
        } else {
            // 将 Request 发送到 Router Handler 中处理业务
            go c.MsgHandler.DoMsgHandle(req)
        }
    }
}

// 写消息到客户端
func (c *Connection) StartWriter() {
    log.Println("[Writer Goroutine is running ...]")
    defer log.Printf("%s connection writer exit\n", c.RemoteAddr().String())
    for {
        select {
        case data := <-c.msgChan:
            if _, err := c.Conn.Write(data); err != nil {
                log.Printf("Send data error: %v\n", err)
                return
            }
        case <-c.ExitChan:
            return
        }
    }
}

func (c *Connection) Start() {
    log.Printf("Conn Start() -- ConnID = %d\n", c.ConnID)
    go c.StartReader()
    go c.StartWriter()
}

func (c *Connection) Stop() {
    log.Printf("Conn Stop() -- ConnID = %d\n", c.ConnID)

    if c.isClosed {
        return
    }
    c.isClosed = true

    _ = c.Conn.Close()
    c.ExitChan <- true
    close(c.ExitChan)
    close(c.msgChan)
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

    c.msgChan <- send
    return nil
}
