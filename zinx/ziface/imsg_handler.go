package ziface

type IMsgHandler interface {
    // 执行对应的方法
    DoMsgHandle(request IRequest)
    AddRouter(msgId uint32, router IRouter)
}
