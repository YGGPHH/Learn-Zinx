package main

import (
	"fmt"
	"zinx/src/ziface"
	"zinx/src/znet"
)

type PingRouter struct {
	znet.BaseRouter // 嵌入 BaseRouter
}

// 实现 IRouter 的三个方法
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write(
		[]byte("Before ping ...\n"),
	)
	if err != nil {
		fmt.Println("Call back Ping Ping Ping error.")
	}
}

func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	_, err := request.GetConnection().GetTCPConnection().Write(
		[]byte("Ping... Ping... Ping...\n"),
	)
	if err != nil {
		fmt.Println("Call back Ping Ping Ping error.")
	}
}

func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write(
		[]byte("After Ping ...\n"),
	)
	if err != nil {
		fmt.Println("Call back Ping Ping Ping error.")
	}
}

func main() {
	s := znet.NewServer("[Zinx V0.3]")

	s.AddRouter(&PingRouter{})

	s.Serve()
}
