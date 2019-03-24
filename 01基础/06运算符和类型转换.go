package main

import "fmt"

// 除法运算符
func main60() {
	a := 10
	b := 3
	// 整型的结果为整型
	// 除数不能为0
	c := a / b
	fmt.Println(c)
}

// 取余运算符
func main61() {
	a := 10
	b := 5
	// 取余只能用于整型数据
	// 不能除以0
	c := a % b
	fmt.Println(c)

}

// 自增,自减运算符,只能写在后面
func main62() {
	a := 10
	a++
	// 不能自增自减运算在表达式中
	// b = a--
	fmt.Println(a)
	a--
	fmt.Println(a)

}

/* 类型转换 */
func main63() {
	a := 10
	b := 3.14
	// 类型转化格式 数据类型(变量)
	// c := float64(a) * b
	// 浮点型转化整型,不会进行四舍五入
	c := a * int(b)
	fmt.Println(c)

}

func main64() {
	fmt.Printf("%d周%d天\n", 46/7, 46%7)
	time := 107653
	fmt.Printf("%d秒\n", time%60)
	fmt.Printf("%d分\n", time/60%60)
	fmt.Printf("%d小时\n", time/60/60%24)
	fmt.Printf("%d天\n", time/60/60/24%365)

}

/* 赋值运算符 */
// += -= /= %=
func main65() {
	a := 10
	// a += 5
	a += 5 * a

	fmt.Println(a)

}

/* 关系运算符 */
// != == >= <= < >
func main66() {
	a := 10
	b := 20
	fmt.Println(a > b)

}

/* 逻辑运算符 */
// ! || &&
func main67() {
	// ! 只能对bool使用
	a := true
	fmt.Println(!a)
	// && 与,两个表达式都为真
	// || 或,一个为真就为真
}

/* 其他运算符 */
// & 取地址
// * 去运算符
func main68() {
	a := 10
	p := &a
	fmt.Println(*p)

}

/* 运算符优先级 */
// 一元运算符最高级,二元从左往右执行
// 括号() 结构体成员. 数组下标[]
// 单目运算符
// 逻辑非! 去地址& 取值* ++ --
// 双目运算符
// 乘除 * / %
// 加减 + -
// 关系运算符  == != > >= <= <
// 逻辑 || &&
// 赋值 = += -= *= /= %=
