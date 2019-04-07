package main

import (
	"fmt"
	"net"
)

// 协议: 一组规则,要求使用协议的双方必须遵守的规则
/* 典型协议:
1. 传输层: TCP/UDP
2. 应用层: HTTP,FTP
3. 网络层: IP协议,ICMP控制报文,IGMP
4. 网络结构层: ARP请求包,获取mac地址,进行局域网广播, RARP */
// OSI/RM: 应用层,表示层,会话层,传输层,网络层,数据链路层,物理层
// TCP/IP: 应用层,传输层,网络层,链路层

// 双向单双工 只能说或者读 对讲机
// 双向全双工 可以同时读说 手机
// 单工通信 遥控器

// 网络预约程序设计模式
// 1. C/S模型,提高数据传输效率,采用的协议相对灵活
// 2. B/S模型,开发量较低,不收平台限制,协议选择不灵活

// socket
// tcp协议: 面向连接的,可靠的数据包传输
// mss: 最大尺寸(滑动窗口,动态变化的)
// 3次握手,建立连接,1个字节8个位 Accept,Dial完成
// 1. 客户端发送 SYN标志位2000(0),给服务端,并发送mss规定数据大小
// 2. 服务端接收,做ACK 2001应答,并发送SYN7000(0)标志位给客户端(7000(0)发送标志位,数据大小为0)
// 3. 客户端最后发了一个ACK7001应答给服务端(表示7001之前的数据都收到了)

// 发送数据
// 发送数据包,并发送ACK应答,告诉双方收到了数据

// 4次挥手,关闭连接 (因为半关闭,所以4次挥手)
// 1.客户端主动关闭连接,发送FIN标志位(包好和数据位)
// 2. 服务端收到了FIN标志位,给客户端发送ACK应答(半关闭状态),客户端不能再发数据了,但是可以收
// 3. 服务器发送FIN标志位,告诉客户端我要关闭连接
// 4. 客户端发送ACK应答

// HTTP(应用层)
// 超文本传输协议,基于tcp协议(传输层),ip(网络层)
// https ssl,tls协议

// URL 统一资源定位符,浏览器地址栏内容
// 协议,服务器名称(IP地址),路径和文件名

// http请求
// 请求行: 请求方法 (空格) 请求文件URL (空格) 协议版本(\r\n)
// 请求头: 语法格式: key: value
// 空行: \r\n  --- 代表http请求头结束
// 请求主体: 请求方法对应的数据内容.get方法没有内容s
func main93() {
	// 指定服务器,通信协议,ip地址:端口,创建一个用于监听的socket
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	defer listener.Close()
	fmt.Println("服务器等待客户端建立连接...")
	// 阻塞监听客户端连接,成功建立连接,返回用于通信的socket
	connect, err := listener.Accept()
	if err != nil {
		fmt.Println("连接失败", err)
		return
	}
	fmt.Println("成功建立连接")
	// 读取客户端发送的数据
	buf := make([]byte, 4096)
	n, err := connect.Read(buf)
	if err != nil {
		fmt.Println("read err")
		return
	}
	// 处理数据,打印
	fmt.Println("服务器端读取到数据", string(buf[:n]))
	connect.Close()
}

// 客户端
func main94() {
	// 指定服务器的ip端口
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
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
