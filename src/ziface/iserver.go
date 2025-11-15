package ziface

type IServer interface {
	Start() // 启动服务器的方法
	Stop()  // 停止服务器的方法
	Serve() // 开启业务服务的方法
}
