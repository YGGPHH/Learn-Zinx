package znet

import (
	"zinx/src/ziface"
)

type Request struct {
	conn ziface.IConnection // 和客户端建立的连接
	msg  ziface.IMessage    // 客户端请求的数据
}

// 获取请求连接信息
func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

// 获取请求消息的数据
func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

// 获取请求消息的数据
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
