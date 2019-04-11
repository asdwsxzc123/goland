package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

// 工作流程:
// 	1. 明确url
// 	2. 发送请求,
// 	3. 保存数据,过滤
// 	4.
func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()
	// 循环读取数据
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取完成")
			return
		}
		if err2 != nil && err2 != io.EOF {
			fmt.Println("read err", err2)
			err = err2
			return
		}
		result += string(buf[:n])
	}
	return result, err
}
func SpiderPage(i int, page chan int) {
	take := 50
	url := "http://tieba.baidu.com/f?kw=%E6%9D%83%E5%8A%9B%E7%9A%84%E6%B8%B8%E6%88%8F&ie=utf-8&pn=" + strconv.Itoa(take*i-1)
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("httpGet err", err)
	}
	// 将数据保存成文件
	path := "./test/"
	name := "第" + strconv.Itoa(i) + "页.html"
	f, err := os.Create(path + name)
	if err != nil {
		fmt.Println("create err", err)
		return
	}
	f.WriteString(result)

	f.Close()
	page <- i
}
func working(start int, end int) {
	fmt.Printf("正在爬取第%d-第%d页", start, end)
	page := make(chan int)
	for i := start; i <= end; i++ {
		go SpiderPage(i, page)
	}
	for {
		fmt.Printf("第 %d 个页面爬取完成\n", <-page)
	}
}
func main151() {
	var start, end int
	fmt.Println("请输入起始页")
	fmt.Scan(&start)
	fmt.Println("请输入结束页")
	fmt.Scan(&end)
	working(start, end)
}

/* 正则 */
func main() {
	str := "abc a7c mfc cat4 8ca azc 3.14 5.16 cba"
	// 解析,编译正则表达式
	// ret := regexp.MustCompile(`a.c`)
	// 匹配中间带数字
	// ret := regexp.MustCompile(`a[0-9]c`)
	// 批量小数
	// ret := regexp.MustCompile(`\d+\.\d+`)
	// 匹配标签
	// ?s:单行模式,点默认不能匹配\n,写了后可以匹配\n
	ret := regexp.MustCompile(`<div>(?s:(.*?))</div>`)

	// 提取需要信息
	alls := ret.FindAllStringSubmatch(str, -1)
	fmt.Println(alls)
}
