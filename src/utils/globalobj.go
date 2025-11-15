package utils

import (
	"encoding/json"
	"os"
	"zinx/src/ziface"
)

type GlobalObj struct {
	TcpServer ziface.IServer // 当前 Zinx 的全局 Server 对象
	Host      string         // 当前服务器主机 IP
	TcpPort   int            // 当前服务器监听的端口号
	Name      string         // 当前服务器名称
	Version   string         // 当前 Zinx 版本号

	MaxPacketSize uint32 // 读取数据包的最大字节数
	MaxConn       int    // 当前服务器允许的最大连接个数
}

var GlobalObject *GlobalObj // 定义一个全局对象

// 读取用户配置文件
func (g *GlobalObj) Reload() {
	data, err := os.ReadFile("./config/zinx.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

// init 方法, 默认加载
func init() {
	GlobalObject = &GlobalObj{
		Name:          "Zinx Server",
		Version:       "V0.4",
		TcpPort:       7777,
		Host:          "0.0.0.0",
		MaxConn:       12000,
		MaxPacketSize: 4096,
	}

	// 从配置文件中加载用户配置的参数
	GlobalObject.Reload()
}
