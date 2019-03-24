package main

import "fmt"

func main42() {
	// var a int
	// & 运算符, 取地址运算符
	// fmt.Scan(&a)
	// fmt.Println(a)
	//  %p 占位符 表示输出一个数据对应的内存地址 &a
	// 0x表示十六进制数据
	// fmt.Printf("%p", &a)

	// 空格或会从作为接受结果
	// var str string
	// fmt.Scan(&str)
	// fmt.Println(str)

	// 接受两个数据
	// var s1, s2 string
	// fmt.Scan(&s1, &s2)
	// fmt.Println(s1)
	// fmt.Println(s2)

}
func main43() {
	var r float64
	PI := 3.14159
	// 通过键盘获取半径
	fmt.Scan(&r)
	fmt.Printf("面积: %.2f\n", PI*r*r)
	fmt.Printf("周长: %.2f\n", PI*2*r)

}
func main41() {
	var a int
	var b string
	// 格式化输入
	fmt.Scanf("%d%c", &a, &b)

}
