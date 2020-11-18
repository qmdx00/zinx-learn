package znet

import "zinx-learn/zinx/ziface"

// 实现router时的基础router，根据需求重写相应的方法
type BaseRouter struct {
}

func (b *BaseRouter) PreHandler(request ziface.IRequest) {}

func (b *BaseRouter) Handler(request ziface.IRequest) {}

func (b *BaseRouter) PostHandler(request ziface.IRequest) {}
