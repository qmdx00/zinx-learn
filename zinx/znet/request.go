package znet

import "qmdx00.cn/zinx/ziface"

type Request struct {
    Conn ziface.IConnection
    Data []byte
}

func (r *Request) GetConn() ziface.IConnection {
    return r.Conn
}

func (r *Request) GetData() []byte {
    return r.Data
}
