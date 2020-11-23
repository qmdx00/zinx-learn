package znet

import (
    "log"
    "strconv"
    "zinx-learn/zinx/utils"
    "zinx-learn/zinx/ziface"
)

type MsgHandler struct {
    Apis           map[uint32]ziface.IRouter
    TaskQueue      []chan ziface.IRequest
    WorkerPoolSize uint32
}

func NewMsgHandler() *MsgHandler {
    return &MsgHandler{
        Apis:           make(map[uint32]ziface.IRouter),
        WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
        TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
    }
}

func (handler *MsgHandler) DoMsgHandle(request ziface.IRequest) {
    hd, ok := handler.Apis[request.GetMsgId()]
    if !ok {
        log.Printf("api msgId = %d Not Found, Need to be registed\n", request.GetMsgId())
    }
    hd.PreHandler(request)
    hd.Handler(request)
    hd.PostHandler(request)
}

func (handler *MsgHandler) AddRouter(msgId uint32, router ziface.IRouter) {
    if _, ok := handler.Apis[msgId]; ok {
        panic("repeated Api, msgId = " + strconv.Itoa(int(msgId)))
    }
    handler.Apis[msgId] = router
    log.Printf("Add Api (msgId = %d) succeed\n", msgId)
}

func (handler *MsgHandler) StartWorkerPool() {
    for i := 0; i < int(handler.WorkerPoolSize); i++ {
        handler.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTask)
        go handler.startWorker(i, handler.TaskQueue[i])
    }
}

func (handler *MsgHandler) startWorker(workId int, taskQueue chan ziface.IRequest) {
    log.Printf("Worker ID = %d is started ...\n", workId)
    for {
        select {
        case request := <-taskQueue:
            handler.DoMsgHandle(request)
        }
    }
}

func (handler *MsgHandler) SendMsgToTaskQueue(request ziface.IRequest) {
    workId := request.GetConn().GetConnID() % handler.WorkerPoolSize
    log.Printf("Add ConnID = %d, Request MsgId = %d to WorkerId = %d\n",
        request.GetConn().GetConnID(),
        request.GetMsgId(),
        workId)
    handler.TaskQueue[workId] <- request
}
