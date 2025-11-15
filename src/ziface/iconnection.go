package ziface

import "net"

// 定义处理连接的接口
type IConnection interface {
	Start()            // 启动连接, 让当前连接开始工作
	Stop()             // 停止连接, 结束当前连接状态
	GetConnID() uint32 // 获取当前连接的 ID
}

// 定义一个统一处理连接业务的接口
type HandFunc func(*net.TCPConn, []byte, int) error
