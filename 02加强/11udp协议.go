package main

import (
	"fmt"
	"net"
	"time"
)

// UDP:无连接的,不可靠的报文传递
// 1. 创建用于通信的socket
// 2. 阻塞读socket
// 3. 处理读到的数据
// 4. 写数据给客户端
// nc -u 127.0.0.1 8003 测试
func main111() {
	// 组织一个udp地址结构
	srvAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("ResolveUDPAddr err", err)
		return
	}
	fmt.Println("udp监听中..")
	udpConn, err := net.ListenUDP("udp", srvAddr)
	if err != nil {
		fmt.Println("ListenUDP err", err)
		return
	}
	defer udpConn.Close()
	fmt.Println("创建socket成功!")
	// 读取客户端发送的数据
	buf := make([]byte, 4096)
	n, cltAddr, err := udpConn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("readFromUDP err", err)
		return
	}
	// 模拟处理数据
	fmt.Printf("服务器读到:%v的数据,%s\n", cltAddr, string(buf[:n]))
	// 提前系统当前时间
	daytime := time.Now().String()

	// 回写数据给客户端
	_, err1 := udpConn.WriteToUDP([]byte(daytime), cltAddr)
	if err1 != nil {
		fmt.Println("writeToUDP err1", err1)
		return
	}
}

// 客户端
func main112() {
	// 指定服务器的ip端口
	conn, err := net.Dial("udp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("连接失败")
		return
	}
	defer conn.Close()
	// 主动向服务器写数据
	conn.Write([]byte("Are you Ready?"))
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("读取失败")
		return
	}
	fmt.Println(string(buf[:n]))
}

// tcp, udp的区别
// tcp:
// 	面向连接
// 	需求系统资源较多
// 	tcp程序结构较复杂
// 	使用流式
// 	保证数据准确性
// 	保证数据顺序
// 	通讯速度较慢
//  对不稳定网络层,做完全弥补操作
// 使用场景: 对数据传输安全性,稳定性要求较高的场合,网络文件传输,下载,上传
// 优点: 稳定,安全,有序
// 缺点: 效率低,开销大,开发复杂度高
// udp:
// 	面向无连接
// 	要求资源较少
// 	udp程序结构较简单
// 	使用数据包式
// 	不保证数据准确性
// 	不保证数据顺序
// 	通讯速度较快
// 	对不稳定的网络层,不作为.
// 使用场景: 对数据传输实时性要求较高的场合,视频直播,在线电话会议,游戏
// 优点: 效率高,开销小,开发复杂度低
// 缺点: 稳定性差,安全性低,无序
