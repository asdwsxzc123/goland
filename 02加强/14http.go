package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

// 回调函数
// 1. 本质:函数指针,通过地址,在某一特定条件调用,调用函数.
// 2. 在程序中,定义了一个函数,但不显示调用,当某一条件满足时,该函数由操作系统自动调用
func handler(w http.ResponseWriter, r *http.Request) {
	// w:写回给客户端的数据
	// r: 从客户端读到的数据
	w.Write([]byte("hello world"))
	fmt.Println(r.Header)
	fmt.Println(r.URL)
	fmt.Println(r.Method)
	fmt.Println(r.Host)
	fmt.Println(r.RemoteAddr)
	fmt.Println(r.Body)
}
func httpServer() {
	// 注册回调函数,该回调函数会在服务器被访问时,自动被调用
	http.HandleFunc("/it", handler)
	// 绑定服务器监听地址
	http.ListenAndServe("127.0.0.1:8008", nil)
}
func errFunc(err error, info string) {
	if err != nil {
		fmt.Println(info, err)
		// return // 返回函数调用
		// runtime.Goexit() // go程结束
		os.Exit(1)
	}

}
func httpClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:8008")
	defer conn.Close()
	errFunc(err, "tcp")
	httpRequest := "GET /it HTTP/1.1\r\nHost:127.0.0.1:8000\r\n\r\n"
	conn.Write([]byte(httpRequest))
	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	errFunc(err, "read")
	if n == 0 {
		return
	}
	fmt.Println(string(buf[:n]))
}
func main141() {
	args := os.Args
	command := args[1]
	if command == "server" {
		httpServer()
	} else if command == "client" {
		httpClient()
	}
}
func OpenSendFile(url string, w http.ResponseWriter) {
	pathFileName := "/Users/li/Desktop/git/goland/02加强/test" + url
	f, err := os.Open(pathFileName)
	if err != nil {
		fmt.Println("文件不存在")
		w.Write([]byte("No such file or directory!"))
		return
	}
	defer f.Close()
	buf := make([]byte, 4096)
	for {
		n, err := f.Read(buf)
		if n == 0 {
			return
		}
		errFunc(err, "read err")
		w.Write(buf[:n])
	}
}
func myHandle(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	fmt.Println("客户端请求: ", url)
	OpenSendFile(url, w)
}
func main142() {
	http.HandleFunc("/", myHandle)
	http.ListenAndServe("127.0.0.1:8008", nil)
}

// http客户端
func main143() {
	resp, err := http.Get("http://www.baidu.com")
	errFunc(err, "get err")
	defer resp.Body.Close()
	fmt.Println(resp.Header)
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Proto)
	buf := make([]byte, 4096)
	var result string
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("--- read finish")
			return
		}

		if err != nil {
			fmt.Println(err)
		}

		result += string(buf[:n])
	}
	fmt.Printf("|%v|\n", result)
}
