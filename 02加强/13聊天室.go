package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// 1. 处理用户连接go程 handleConnect
// 2. 用户消息广播 go程,manager
// 3. 主go程
// 4. go程间 应用数据及通信
// map: 存储所有登录聊天室的用户消息,key: 用户的IP+port,value:client结构体
// client结构体: 包含成员,用户名name,网络地址addr(ip + port),发送消息的通道c(channel)
// 通道message: 协调并发go程间信息的传递
// 创建用户结构体类型

/* 广播用户上线流程 */
// 1. 主go程,创建监听套接字,要记住defer
// 2. for 循环监听客户端连接请求,accept
// 3. 有一个客户端连接,创建新go程处理客户端数据 HandlerConnect(conn) defer
// 4. 定义全局结构体类型 C,Name,RemoteAddr
// 5. 创建全局map,channel
// 6. 实现handlerConnel,获取客户端ip+port -- remoteAddr(),初始化新用户结构体信息,name == Addr
// 7. 创建manger管理go程, -- accept之前
// 8. 实现manager,初始化在线用户map,循环读取全局channel,如果没有数据阻塞,有数据,遍历在线用户map,将数据写到用户的C中
// 9. 添加新用户到在线用户中,key == ip+port value=新用户结构体
// 10. 创建writeMsgToClient go程,专门给当前用户写数据,-- 来源于用户自带的C
// 11. 实现writeMsgToClient(client,conn) 将数据发送到客户端
// 12. handlerConnet中,结束为止,发送用户上线信息

/* 用户发送消息 */
// 1. 封装函数MakeMsg来处理广播,用户消息
// 2. HandlerConnect中,创建匿名go程,读取用户消息,写入到全局channel
// 3. for循环conn.Read n== 0 err != nil
// 4. 写给全局message 原来广播用户上线模块完成

/* 查询在线用户 */
// 1. 将读取到的用户消息msg结尾的"\n"去掉
// 2. 是否是"who"命令,如果是返回用户列表
// 3. 遍历onlineMap将用户,组织显示信息,写到套接字中

/* 用户改名 */
// 1. 是否是"rename|newName"命令,如果是改名
// 2. 提取"|"后面的字符串,存入到client的name成员中
// 3. 更新在线用户列表,onlineMap,key -- ip + port
// 4. 提示用户更新成功

/* 用户退出 */
// 1. 在用户成功登陆之后创建一个监听用户退出的channel-- isQuit
// 2. 当conn.read == 0 ,isQuit <- true
// 3. 在handleConnect 结尾 for中,添加select 监听 <-isQuit
// 4. 条件满足,将用户从在线列表移除,组织用户下线消息,写入message(广播通道)

/* 超时处理 */
// 1. 创建监听用户超时的channel,将用户重在线列表移除,组织用户下线消息,写入message(广播)
// 2. 创建监听,用户活跃的channel -- hasData
// 3. 只用户只需: 聊天,改名,who任意一个操作,hasData <-true
// 4. 在select中,添加监听<-channel,条件满足,不做任何事情,目的是重置计算器
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
func MakeMsg(client Client, msg string) (buf string) {
	buf = client.Name + ": " + msg
	return buf
}
func HandlerConnect1(conn net.Conn) {
	defer conn.Close()
	// 创建一个channel判断用户是否活跃
	hasData := make(chan bool)
	// 获取用户的网络地址
	netAddr := conn.RemoteAddr().String()
	// 创建用户用户名IP+port
	client := Client{make(chan string), netAddr, netAddr}
	// 将新连接用户添加到在线用户map中
	onlineMap[client.Addr] = client
	// 创建专门用来给当前用户发送消息的go程
	go WriteMsgToClient(client, conn)
	// 发送 用户上线消息到 全局channel中
	message <- MakeMsg(client, "login")

	// 创建一个channel1,判断用户退出状态
	isQuit := make(chan bool)
	// 创建一个go程,专门处理用户发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Printf("检测到客户端:%s退出\n", client.Name)
				return
			}
			if err != nil {
				fmt.Println("read err", err)
				return
			}
			// 将读到的用户信息,写入到messag
			msg := string(buf[:n-1])
			// 提取在线用户列表
			if msg == "who" && len(msg) == 3 {
				conn.Write([]byte("online user list:\n"))
				// 遍历当前map,获取在线用户
				for _, user := range onlineMap {
					userInfo := user.Addr + ":" + user.Name + "\n"
					conn.Write([]byte(userInfo))
				}
				// 判断用户发送了改名命令
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				newName := strings.Split(msg, "|")[1]
				client.Name = newName       // 修改结构体成员name
				onlineMap[netAddr] = client // 更新onlineMap
				conn.Write([]byte("rename successful: " + newName))
			} else {
				message <- MakeMsg(client, msg)
			}
			hasData <- true
		}
	}()
	// 保证
	for {
		// 监听channel上的数据流动
		select {
		case <-isQuit:
			close(client.C)
			delete(onlineMap, client.Addr)       // 将用户从onlie移除
			message <- MakeMsg(client, "logout") // 写入用户退出提示
			return
		case <-hasData:
			// 什么都不做,为了重置下面的定时
		case <-time.After(time.Second * 10):
			delete(onlineMap, client.Addr)       // 将用户从onlie移除
			message <- MakeMsg(client, "logout") // 写入用户退出提示
			return
		}
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
func main131() {
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
