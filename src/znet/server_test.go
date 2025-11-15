package znet

import (
	"fmt"
	"net"
	"testing"
	"time"
)

// 模拟客户端
func ClientTest() {

	fmt.Println("Client Test... Start...")

	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "localhost:7777")
	if err != nil {
		fmt.Println("Client Start err, exit!")
		return
	}

	for {
		_, err := conn.Write([]byte("Hello, YGGP")) // 基于 conn.Write 可以向连接写入字节
		if err != nil {
			fmt.Println("Write buf error, err: ", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf) // 基于 conn.Read 从连接读取字节
		if err != nil {
			fmt.Println("Read buf error...")
			return
		}

		fmt.Printf("Server Call Back: %s, cnt = %d\n", buf, cnt)
		time.Sleep(1 * time.Second)
	}
}

func TestServer(t *testing.T) {
	s := NewServer("[Zinx V0.1]")

	go ClientTest()

	s.Serve()
}
