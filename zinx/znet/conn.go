package znet

import (
    "log"
    "net"
    "qmdx00.cn/zinx/ziface"
)

// IConnection implement
type Connection struct {
    // 当前链接的套接字
    Conn *net.TCPConn
    // 链接的ID
    ConnID uint32
    // 告知当前链接是否退出
    ExitChan chan bool
    // 当前链接Router
    Router ziface.IRouter

    // 当前的链接是否关闭
    isClosed bool
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
    return &Connection{
        Conn:     conn,
        ConnID:   connID,
        isClosed: false,
        ExitChan: make(chan bool, 1),
        Router:   router,
    }
}

func (c *Connection) StartReader() {
    log.Println("Reader Goroutine is running ...")
    defer log.Printf("[ConnID = %d] Reader is exit, Remote addr is %s", c.ConnID, c.RemoteAddr().String())
    defer c.Stop()

    for {
        // 读取数据到buff中
        buf := make([]byte, 512)
        cnt, err := c.Conn.Read(buf)
        if err != nil {
            log.Printf("receive buffer error: %v", err)
            continue
        }

        req := &Request{
            Conn: c,
            Data: buf[:cnt],
        }

        go func(request *Request) {
            c.Router.PreHandler(req)
            c.Router.Handler(req)
            c.Router.PostHandler(req)
        }(req)
    }

}

func (c *Connection) Start() {
    log.Printf("Conn Start() -- ConnID = %d", c.ConnID)
    go c.StartReader()
}

func (c *Connection) Stop() {
    log.Printf("Conn Stop() -- ConnID = %d", c.ConnID)

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

func (c *Connection) Send(data []byte) error {
    panic("implement me")
}
