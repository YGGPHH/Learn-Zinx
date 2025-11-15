package ziface

/*
	IRequest Interface:
	将客户端请求的连接信息以及请求数据封装到了 Request 当中
*/

type IRequest interface {
	GetConnection() IConnection // 获取请求连接信息
	GetData() []byte            // 获取请求消息的数据
	GetMsgID() uint32           // 获取消息的 Id
}
