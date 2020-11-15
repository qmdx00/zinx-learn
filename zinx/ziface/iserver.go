package ziface

type IServer interface {
    Start()
    Stop()
    Serve()
    // 注册路由，供客户端链接处理业务
    AddRouter(IRouter)
}
