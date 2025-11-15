package znet

import (
	"fmt"
	"net"
	"zinx/src/ziface"
)

// 定义具体的连接类型, 它将要实现 IConnection 接口
type Connection struct {
	Conn         *net.TCPConn    // 当前连接的 Socket 套接字
	ConnID       uint32          // 当前连接的 ID, ID 全局唯一
	isClosed     bool            // 判断当前连接是否关闭
	handleAPI    ziface.HandFunc // 该连接的处理方法
	ExitBuffChan chan bool       // 告知该连接已经退出的 Channel
}

func NewConnection(conn *net.TCPConn, connID uint32, callback_api ziface.HandFunc) *Connection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		handleAPI:    callback_api,
		ExitBuffChan: make(chan bool, 1),
	}

	return c
}

// 处理 Conn 读数据的 goroutine
func (c *Connection) StartReader() {
	fmt.Println("Reader goroutine is running...")
	defer fmt.Println(c.Conn.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err ", err)
			c.ExitBuffChan <- true
			// 发生错误, 通过 Channel 关闭当前连接
			continue
		}
		// 否则, 调用当前连接的业务函数
		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println("ConnID ", c.ConnID, " handle is error")
			c.ExitBuffChan <- true
			return
		}
	}
}

// 启动连接
func (c *Connection) Start() {
	go c.StartReader()

	for {
		select {
		case <-c.ExitBuffChan:
			// 监控 ExitBuffChan 管道, 如果监控到退出信号, 则退出.
			return
		}
	}
}

// 停止连接, 结束当前连接状态
func (c *Connection) Stop() {
	if c.isClosed {
		return
	}
	c.isClosed = true

	// TODO: Connection Stop() 时, 如果用户注册了该链接的回调业务, 则应该执行

	// 关闭 socket 连接
	c.Conn.Close()

	// 通知从缓冲队列对数据的业务, 当前连接结束
	c.ExitBuffChan <- true

	close(c.ExitBuffChan) // 关闭连接的管道
}

// 从当前连接获取原始的 socket TCPConn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// 获取当前连接的 ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// 获取远程客户端的地址信息
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
