package ziface

type IRequest interface {
    // 获取链接
    GetConn() IConnection
    // 获取数据
    GetData() []byte
    // 获取消息ID
    GetMsgId() uint32
}
