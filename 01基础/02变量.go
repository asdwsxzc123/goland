// 内存是一个连续的数据集合,每一个内存存储区都有唯一的地址标识,成为内存地址
// 内存地址编号是一个无符号十六进制整型数据表示的,可以为内存指定区域起别名称为变量名
// 计算机能够识别的数据格式: 二进制,八进制,十进制,十六进制
package main

import "fmt"

func main21() {
	// 存储在内存的栈区,在程序运行结束后会自动释放
	// var 变量名 数据类型
	// 布尔 整型 浮点型 字符型 字符串了下
	var a int
	a = 10
	a = a + 25
	fmt.Println(a)
}

func main22() {
	// 计算圆的面积和轴承
	// 面积 PI* r * r
	// 周长 2 * PI * r
	var PI float64 = 3.14159
	var r float64 = 2.5
	var s float64
	var l float64
	// 计算面积
	s = PI * r * r
	l = PI * 2 * r
	fmt.Println("面积,", s)
	fmt.Println("周长,", l)
}

/* 自动推到类型 */

func main23() {
	// 计算圆的面积和轴承
	// 面积 PI* r * r
	// 周长 2 * PI * r
	PI := 3.14159
	r := 2.5
	s := PI * r * r
	l := PI * 2 * r
	fmt.Println("面积,", s)
	fmt.Println("周长,", l)
}

func main24() {
	// 在go中不同的数据类型不能计算操作,可以通过类型转换解决
	w := 2.0 // float64
	// p := 5 // int
	var p float64 = 5
	fmt.Println(w * p)
}

/* 数据类型 */
func main25() {
	a := false // 布尔类型
	b := 10    // int
	c := 1.2   // float64
	d := 'a'   // byte
	e := "你好"  // 字符串
	fmt.Println(a, b, c, d, e)
}

/* 多重赋值 */
func main26() {
	a, b, c, d, e := 1, 2, 3, 4, 5.2
	fmt.Println(a, b, c, d, e)
}

/* 交换变量的值 */
func main27() {
	a, b := 1, 2
	a, b = b, a
	// 匿名变量,
	// _ := 1
	fmt.Println(a, b)
}

/*  变量规范 */
// 1. 名称必须以一个字母或下划线开头
// 2. 只能使用字母,数字,下划线
// 3. 区分大小写
// 4. 关键字: break,default,func,interface,select,case,defer,go,map,struct,chan,else,goto,package,switch,const,fallthrough, if ,range ,type,continue, for,import,return, var
// 5. 预定义的名称: int, true, false,iota,nil,uint16,complex128,bool,byte,rune,string,error,make, len,cap,new,append,copy,close,delete,panic,recover
// 6. 正确的命名: principal, cost_price,marks_3,city
// 7: 大驼峰,小驼峰,下划线连接

/* 常量 */
func main28() {
	// 常量的定义和使用,存储在内存的数据区,不能通过&来取地址,一般用大写定义
	const A int = 10
	// 常量的值不允许修改,一般用在不能修改的地方,生日id,pi,游戏中商品的固定价格
	// A = 20
	fmt.Println(A)

}
func main29() {
	// 字面常量
	c := 20
	fmt.Println(123)
	fmt.Println("hello")
	// 硬常量 32
	d := c + 32
	fmt.Println(d)
}

/* iota枚举 */
// 常量声明可以使用iota常量生成器初始化,iota将会被置为0
func main211() {
	// 1. iota常量自动生成器,每个一行,自动累加1
	// 2. iota给常量赋值使用
	// const (
	// 	a = iota // 0
	// 	b = iota // 1
	// 	c = iota // 2
	// )
	// fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
	// value := a
	// fmt.Println(value)
	// value = b
	// fmt.Println(value)

	// 每次换行才会加1
	const (
		a    = iota
		b, c = iota, iota
		d, e
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
