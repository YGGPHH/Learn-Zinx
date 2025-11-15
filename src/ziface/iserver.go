package ziface

type IServer interface {
	Start()                   // 启动服务器的方法
	Stop()                    // 停止服务器的方法
	Serve()                   // 开启业务服务的方法
	AddRouter(router IRouter) // V0.3: 为当前服务注册一个路由业务方法
}
