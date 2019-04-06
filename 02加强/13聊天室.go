package main

import (
	"fmt"
	"net"
)

// 1. 处理用户连接go程 handleConnect
// 2. 用户消息广播 go程,manager
// 3. 主go程
// 4. go程间 应用数据及通信
// map: 存储所有登录聊天室的用户消息,key: 用户的IP+port,value:client结构体
// client结构体: 包含成员,用户名name,网络地址addr(ip + port),发送消息的通道c(channel)
// 通道message: 协调并发go程间信息的传递
// 创建用户结构体类型
type Client struct {
	C    chan string
	Name string
	Addr string
}

// 创建全局map,存储在线用户
var onlineMap map[string]Client

// 创建全局channel传递用户消息
var message = make(chan string)

func WriteMsgToClient(client Client, conn net.Conn) {
	for msg := range client.C {
		conn.Write([]byte(msg + "\n"))
	}
}
func HandlerConnect1(conn net.Conn) {
	defer conn.Close()
	// 获取用户的网络地址
	netAddr := conn.RemoteAddr().String()
	// 创建用户用户名IP+port
	client := Client{make(chan string), netAddr, netAddr}
	// 将新连接用户添加到在线用户map中
	onlineMap[client.Addr] = client
	// 创建专门用来给当前用户发送消息的go程
	go WriteMsgToClient(client, conn)
	// 发送 用户上线消息到 全局channel中
	message <- "[" + netAddr + "]" + client.Name + "上线了"

	// 保证
	for {

	}
}
func Manager() {
	// 初始化map
	onlineMap = make(map[string]Client)
	for {
		// 监听全局channel中是否有数据,有数据存储,没数据阻塞
		msg := <-message
		// 训话发送消息给所有在线用户
		for _, client := range onlineMap {
			client.C <- msg
		}
	}
}
func main() {
	// 创建监听套接字
	listener, err := net.Listen("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	defer listener.Close()
	// 创建一个goManger,管理全局channel和map
	go Manager()
	// 有用户连接过来
	for {
		conn, err := listener.Accept()
		fmt.Println("用户连接成功!")
		if err != nil {
			fmt.Println("accept err", err)
			return
		}
		// 启动go程处理客户端数据秦秋
		go HandlerConnect1(conn)
	}
}
