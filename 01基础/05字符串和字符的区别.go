package main

import "fmt"

func main51() {
	var a string = "hello\nworld"
	fmt.Println(a)
}
func main52() {
	var str string = "你好"
	// 计算字符串个数,一个汉字是3个字符
	num := len(str)
	fmt.Println(num)

}

/* 占位符的使用 */
func main53() {
	// var a int = 10
	a := 10
	fmt.Printf("%d\n", a)

	// var b float64 = 10
	b := 10.0
	// 默认保留6位小数
	fmt.Printf("%f\n", b)

	// var c bool = true
	c := true
	fmt.Printf("%t\n", c)

	var d byte = 'a'
	fmt.Printf("%c\n", d)

	e := "hello"
	fmt.Printf("%s\n", e)

	// 占位符
	fmt.Printf("%p\n", &a)

	// %T,打印变量对应的数据类型
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)

	// %% 打印一个百分号的信息
	fmt.Printf("35%%\n")
}

func main54() {
	// 计算机能识别的进制 2,8,10,16
	// 十进制数据
	a := 123

	// 八进制
	b := 0123

	// 十六进制
	c := 0x123

	// go语言不能直接表示二进制数据
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// %o表示8进制
	fmt.Printf("%o\n", a)
	// %x表示16进制, %X表示16进制,根据写的字母的大小写
	fmt.Printf("%x\n", a)
}
