package znet

import (
    "errors"
    "log"
    "sync"
    "zinx-learn/zinx/ziface"
)

type ConnManager struct {
    Connections map[uint32]ziface.IConnection
    ConnLock    sync.RWMutex
}

func NewConnManager() *ConnManager {
    return &ConnManager{
        Connections: make(map[uint32]ziface.IConnection),
    }
}

func (c *ConnManager) Add(conn ziface.IConnection) {
    c.ConnLock.Lock()
    defer c.ConnLock.Unlock()

    c.Connections[conn.GetConnID()] = conn
    log.Printf("Connection [ConnID = %d] Add to ConnManager Successful: connections number = %d\n",
        conn.GetConnID(), c.Size())
}

func (c *ConnManager) Remove(conn ziface.IConnection) {
    c.ConnLock.Lock()
    defer c.ConnLock.Unlock()

    delete(c.Connections, conn.GetConnID())
    log.Printf("Connection [ConnID = %d] Remove from ConnManager Successful: connections number = %d\n",
        conn.GetConnID(), c.Size())
}

func (c *ConnManager) Get(connId uint32) (ziface.IConnection, error) {
    c.ConnLock.RLock()
    defer c.ConnLock.RUnlock()

    if conn, ok := c.Connections[connId]; ok {
        return conn, nil
    } else {
        return nil, errors.New("Connection not found ")
    }
}

func (c *ConnManager) Size() int {
    return len(c.Connections)
}

func (c *ConnManager) Clear() {
    c.ConnLock.Lock()
    defer c.ConnLock.Unlock()

    for connId, conn := range c.Connections {
        conn.Stop()
        delete(c.Connections, connId)
    }
    log.Printf("Clear All Connections Successful: Connections number = %d\n", c.Size())
}
