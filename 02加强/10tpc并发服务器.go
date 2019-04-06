package main

import (
	"fmt"
	"net"
	"strings"
)

func HandlerConnect(conn net.Conn) {
	defer conn.Close()
	// 获取连接的客户端addr
	addr := conn.LocalAddr()
	fmt.Println(addr, "客户端连接成功:")

	for {
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if "exit\n" == string(buf[:n]) {
			fmt.Println("服务端接受客户端退出请求,关闭服务器")
			return
		}
		if n == 0 {
			fmt.Println("服务器检测到客户端已关闭,端口连接")
			return
		}
		if err != nil {
			fmt.Println("read err", err)
			return
		}
		fmt.Println("读取数据", string(buf[:n]))
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}
func main101() {
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	fmt.Println("等待连接中...")
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("connect err", err)
			return
		}
		go HandlerConnect(conn)
	}
}
