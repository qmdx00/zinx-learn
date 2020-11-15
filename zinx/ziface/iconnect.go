package ziface

import "net"

type IConnection interface {
    // 启动链接
    Start()
    // 停止链接
    Stop()
    // 获取TCP链接
    GetTCPConn() *net.TCPConn
    // 获取链接ID
    GetConnID() uint32
    // 客户端TCP状态信息 IP和端口
    RemoteAddr() net.Addr
    // 发送数据给客户端
    SendMsg(uint32, []byte) error
}


type HandleFunc func(*net.TCPConn, []byte, int) error
