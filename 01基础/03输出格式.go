package main

import "fmt"

func main9() {
	// 可以打印各种格式,打印时自带换行
	// fmt.Println(22)
	// 可以打印各种格式,打印时不带换行
	// fmt.Print(11)

	// 整型
	// 占位符,打印时不带换行
	fmt.Printf("==%3d==\n", 10)  // int
	fmt.Printf("==%-3d==\n", 10) // int
	// 用0补位保留5位数
	fmt.Printf("==%05d==\n", 10) // int

	// 浮点型
	// 默认保留6位小数
	fmt.Printf("%f\n", 10.1) // float64
	// 保留3位小数,会四舍五入
	fmt.Printf("%.3f\n", 10.1) // float64

	// 布尔
	var a bool
	fmt.Printf("%t\n", a) // 字符串
	var b string
	b = "10.1"
	fmt.Printf("%s\n", b) // 字符串

	var c byte
	c = 'a'
	// 字符型变量对应的ascii码值
	fmt.Printf("%c\n", c)
}
