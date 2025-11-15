package ziface

import "net"

// 定义处理连接的接口
type IConnection interface {
	Start()                                  // 启动连接, 让当前连接开始工作
	Stop()                                   // 停止连接, 结束当前连接状态
	GetConnID() uint32                       // 获取当前连接的 ID
	GetTCPConnection() *net.TCPConn          // 获取当前 Connection 底层的 TCP 连接
	RemoteAddr() net.Addr                    // 获取远程客户端的地址信息
	SendMsg(msgId uint32, data []byte) error // 直接将 Message 数据发送给远程的 TCP 客户端
}

// 定义一个统一处理连接业务的接口
type HandFunc func(*net.TCPConn, []byte, int) error
