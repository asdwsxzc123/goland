package main

import "fmt"

func add(a int, b int) {
	sum := a + b
	fmt.Println(sum)

}

// 不定参函数,...数据类型
func test(args ...int) {
	fmt.Println(args)
	// for i := 0; i < len(args); i++ {
	// 	fmt.Println(args[i])
	// }

	// for 和range可以遍历
	for i, data := range args {
		fmt.Println(i, data)
	}
	// for _, data := range args {
	// 	fmt.Println(data)
	// }
}

// 函数嵌套调用
func test1(a int, b int) {
	test2(a, b)
}
func test2(a int, b int) {
	fmt.Println(a, b)
}

func test3(a ...int) {
	test4(a[0:]...)
}
func test4(b ...int) {
	for i := 0; i < len(b); i++ {
		fmt.Println(i)
	}
}

/* 返回值 */
// 函数名 (形参列表) 返回类型列表{代码体}
func test5(a int, b int) int {
	return a + b
}

func test6(a int, b int) (sum int) {
	sum = a - b
	return sum
}

/* 多个返回值 */
func test7() (a int, b int, c int) {
	a, b, c = 1, 2, 3
	return a, b, c
}

/* 函数类型 */
func test8() {
	fmt.Println("hello")
}
func test9(a int, b int) {
	fmt.Println(a, b)
}

func test10(a int, b int) {

}

// type 可以定义函数类型
type FUNCTYPE func()

type FUNCTEST func(int, int)
type funcdemo func(int, int) int

func test11() {
	// 局部作用域 在函数内部定义的变量 作用域限定于函数内部,从变量定义到函数结束有效
	// 在同一作用域范围,变量名是唯一的
	a := 10
	{
		// var a int = 20
		a = 20
		fmt.Println(a)
	}
	fmt.Println(a)

}

// 所有的函数都是全局函数,因此函数都是唯一的
func main91() {
	// add(1, 2)
	// test(1, 2, 3)
	// test1(1, 2)
	// value := test5(1, 2)
	// value := test6(1, 2)
	// fmt.Println(value)
	// a, b, c := test7()
	// fmt.Println(a, b, c)

	// 定义函数类型变量
	// var f FUNCTYPE
	// f = test8
	// // 通过函数变量调用函数
	// f()

	// var f1 FUNCTEST
	// f1 = test9
	// f1(10, 20)

	// // 不通过type定义
	// var f2 func(int, int)
	// f2 = test10
	// f2(10, 20)

	// // 函数自动推倒
	// f4 := test10
	// f4(10, 10)

	// test11()

	/* 匿名函数 */
	// a := 10
	// b := 20
	// func(a int, b int) {
	// 	fmt.Println(a + b)
	// }(a, b)
	// f := func(a int, b int) {
	// 	fmt.Println(a + b)
	// }
	// f(a, b)

	fmt.Println(g_a)
}

// 全局变量,在函数外部定义,在函数中重新赋值,其他地方也会改变
// 全局变量作用域是项目中所有文件
// 全局变量在内存中数据区存储,和const定义的常量都是数据区
var g_a int = 10

/* 闭包 */
func main92() func() int {
	var a int
	return func() int {
		a++
		return a
	}
}
func main93() {
	f := main92()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

}

/* 递归 */
func main94(a int) {
	if a == 0 {
		return
	}
	a--
	main94(a)
	fmt.Println(a)
}
func main95() {
	main94(5)
}
