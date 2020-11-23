package ziface

type IServer interface {
    Start()
    Stop()
    Serve()
    // 注册路由，供客户端链接处理业务
    AddRouter(msgId uint32, router IRouter)
    GetConnManager() IConnManager

    SetOnConnStart(func(conn IConnection))
    SetOnConnStop(func(conn IConnection))
    CallOnConnStart(conn IConnection)
    CallOnConnStop(conn IConnection)
}
