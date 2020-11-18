package znet

import "zinx-learn/zinx/ziface"

type Request struct {
    conn ziface.IConnection
    msg  ziface.IMessage
}

func (r *Request) GetConn() ziface.IConnection {
    return r.conn
}

func (r *Request) GetData() []byte {
    return r.msg.GetMsgData()
}

func (r *Request) GetMsgId() uint32 {
    return r.msg.GetMsgId()
}
