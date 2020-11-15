package ziface

// Router Handle Hooks
type IRouter interface {
    PreHandler(request IRequest)
    Handler(request IRequest)
    PostHandler(request IRequest)
}
