package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/src/ziface"
)

// 定义具体的连接类型, 它将要实现 IConnection 接口
type Connection struct {
	Conn         *net.TCPConn   // 当前连接的 Socket 套接字
	ConnID       uint32         // 当前连接的 ID, ID 全局唯一
	isClosed     bool           // 判断当前连接是否关闭
	Router       ziface.IRouter // 该连接业务方法对应的 Router
	ExitBuffChan chan bool      // 告知该连接已经退出的 Channel
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		Router:       router,
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
		// 创建拆包解包对象
		dp := NewDataPack()

		// 读取客户端的 Msg Head
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("Read msg head error, ", err)
			c.ExitBuffChan <- true
			continue
		}

		// 拆包, 得到 msgId 和 DataLen, 放到 msg 当中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("Unpack error, ", err)
			c.ExitBuffChan <- true
			continue
		}

		// 根据 DataLen 来读取 data, 放到 msg.Data 当中. 在 dp.Unpack 当中只获取数据部分的长度,
		// 不负责获取数据. 获取数据的部分在 Connection 当中, 基于长度构建字节序列读取数据
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("Read message data error, ", err)
				c.ExitBuffChan <- true
				continue
			}
		}
		msg.SetData(data)

		// 构造当前 Request 的数据, 包含当前的连接 Conn 和数据 Data
		req := Request{
			conn: c,
			msg:  msg,
		}

		go func(request ziface.IRequest) {
			// 执行注册的路由方法
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
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

func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("Connection closed when sending message...")
	}
	// 将 data 封包, 并且发送. 这就要求 Zinx 的客户端基于与 Zinx 服务器本身同样的方式进行拆包解包
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgId) // 基于 msgId 可以定位到出问题的 msg
		return errors.New("Pack error message.")
	}

	// 写回给客户端
	if _, err := c.Conn.Write(msg); err != nil {
		fmt.Println("Write msg id ", msgId, " error.")
		c.ExitBuffChan <- true
		return errors.New("Conn write error.")
	}

	return nil
}
