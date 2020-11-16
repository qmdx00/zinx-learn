package znet

import (
    "log"
    "qmdx00.cn/zinx/ziface"
    "strconv"
)

type MsgHandler struct {
    Apis map[uint32]ziface.IRouter
}

func NewMsgHandler() *MsgHandler {
    return &MsgHandler{
        Apis: make(map[uint32]ziface.IRouter),
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
