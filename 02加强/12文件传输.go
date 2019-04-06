package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

// 1. 发送方想服务端发送文件名,服务端保证该文件名
// 2. 接收方向客户端返回一个消息ok,缺点文件名保存成功
// 3. 发送方收到消息后,开始想服务端发送文件数据
// 4. 接受方读取文件内容,写入到之前保存好的文件中
func writeFile1(conn net.Conn, filePath string) {
	// 4. 打开文件
	fmt.Println(filePath)
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("file open err", err)
		return
	}
	// 6. 关闭客户端
	defer f.Close()
	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("接受文件完成")
			} else {
				fmt.Println("read err", err)
			}
			return
		}
		// 5. 写入文件到客户端
		_, err1 := conn.Write(buf[:n])
		if err1 != nil {
			fmt.Println("write err1", err1)
			return
		}
	}
}
func main121() {
	// 接受文件名参数
	list := os.Args
	if len(list) != 2 {
		fmt.Println("请输入格式: go run xxx.go 文件名")
		return
	}
	filePath := list[1]
	// 1. 创建客户端
	conn, err := net.Dial("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("conn err", err)
		return
	}
	defer conn.Close()

	// 2. 发送文件名
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println("stat err", err)
		return
	}
	fileName := fileInfo.Name()
	fmt.Println(fileInfo.Size())
	conn.Write([]byte(fileName))

	// 3. 接受返回值ok
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("read err", err)
		return
	}
	if string(buf[:n]) == "ok" {
		fmt.Println("收到ok")
		writeFile1(conn, filePath)
	}

}

func resvFile(conn net.Conn, fileName string) {
	// 创建文件
	fmt.Println(fileName)
	f, err := os.Create(fileName)
	if err != nil {
		fmt.Println("f create err", err)
		return
	}
	defer f.Close()

	// 5. 接收文件
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if n == 0 {
				fmt.Println("完成文件传输")
			} else {
				fmt.Println("file read err", err)
			}
			return
		}
		// 6. 写入文件
		f.Write(buf[:n])
	}

}

// 接收端
func main122() {
	// 1. 建立连接listen
	listener, err := net.Listen("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	defer listener.Close()
	fmt.Println("建立连接中...")
	// 2. 建立套接字accept
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("accept err", err)
		return
	}
	fmt.Println("建立套接字")
	defer conn.Close()
	buf := make([]byte, 4096)

	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("read err", err)
		return
	}
	// 3. 接受文件名
	fileName := string(buf[:n])

	// 4. 返回ok
	n, err1 := conn.Write([]byte("ok"))
	fmt.Println("发送ok")
	if err1 != nil {
		fmt.Println("write err", err1)
		return
	}
	resvFile(conn, fileName)
}
