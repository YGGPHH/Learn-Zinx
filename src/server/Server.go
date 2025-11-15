package main

import (
	"fmt"
	"zinx/src/ziface"
	"zinx/src/znet"
)

type PingRouter struct {
	znet.BaseRouter // 嵌入 BaseRouter
}

func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	fmt.Println("Recv from client: msgId = ", request.GetMsgID(), ", data = ", string(request.GetData()))

	// 回写数据
	err := request.GetConnection().SendMsg(1, []byte("Ping... Ping... Ping..."))
	if err != nil {
		fmt.Println("Server response error, err: ", err)
	}
}

func main() {
	s := znet.NewServer()

	s.AddRouter(&PingRouter{})

	s.Serve()
}
