package znet

import (
	"fmt"
	"net"
	"time"
	"zinx/src/ziface"
)

// IServer 的接口实现, Server 需要实现 iServer 接口定义的三个方法
type Server struct {
	Name      string // 服务器名称
	IPVersion string // tcp4 or other
	IP        string // 服务器绑定的 IP 地址
	Port      int    // 服务绑定的端口
}

func (s *Server) Start() {
	fmt.Printf("[START] Server Listenning at IP: %s, Port %d, is starting\n", s.IP, s.Port)

	// 开启一个 goroutine 来进行 Listener 业务
	go func() {
		// 1. 获取一个 TCP 的 Addr
		addr, err := net.ResolveTCPAddr(
			s.IPVersion,
			fmt.Sprintf("%s:%d", s.IP, s.Port),
		)
		if err != nil {
			fmt.Println("Resolve TCP addr failed, err: ", err)
			return
		}

		// 2. 监听服务器地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("Listen", s.IPVersion, "err", err)
			return
		}

		fmt.Println("Start Zinx Server ", s.Name, " succ, now listenning...")

		// 3. 启动 Server 网络连接业务
		for {
			// 3.1. 阻塞等待客户端建立连接
			conn, err := listenner.AcceptTCP()
			if err != nil {
				// 如果失败了, 则继续监听新的连接
				fmt.Println("Accept err ", err)
				continue
			}

			// 3.2. TODO: Server.Start() 设置服务器最大连接数, 超过最大连接则关闭这个新的连接

			// 3.3. TODO: Server.Start() 处理该新连接请求的业务方法, 此时应该有 handler, 与 conn 绑定

			// 此时暂时只做 512 Bytes 的字节回显
			go func() {
				// 从客户端获取数据
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err: ", err)
						continue
					}

					// 回显
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("Write back buf err: ", err)
						continue
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx Server, name ", s.Name)

	// TODO: Server.Stop() 应该清理其它连接与信息
}

func (s *Server) Serve() {
	s.Start()

	// TODO: Server.Serve() 如果需要在开启服务时完成其它事情, 可以在此处添加

	// 否则主 Goroutine 退出
	for {
		time.Sleep(10 * time.Second)
	}
}

// 创建服务器的句柄
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7777,
	}

	return s
}
