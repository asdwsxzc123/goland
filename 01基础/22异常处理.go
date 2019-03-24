package main

import (
	"errors"
	"fmt"
)

/* error */
// 编辑时异常
// 编译时异常
// 运行时异常
func test22(a, b int) (value int, err error) {
	// 0不能作为被除数
	if b == 0 {
		err = errors.New("0不能作为除数")
		return
	} else {
		value = a / b
		return
	}
}
func main2201() {
	value, err := test22(10, 10)
	if err == nil {
		fmt.Println(value)
	} else {
		fmt.Println(err)
	}
}

/* panic */
// panic返回让程序崩溃的错误
func test23(a, b int) (value int, err error) {
	// 调用后程序会终止执行
	panic("hello world")
}

func main2202() {
	value, _ := test23(10, 0)
	fmt.Println(value)
}

/* defer */
// defer调用的函数没有直接使用,而是先假装到栈区内存,在函数结束时调用,出栈调用
func main2203() {
	fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
	defer fmt.Println("4")
	// 1,3,4,2
}

/* recover接口拦截一次 */
func test24(i int) {
	var arr [10]int
	// 通过匿名函数和recover进行错误拦截
	// recover需要放到错误之前,捕获异常
	defer func() {
		// recover 可以从panic异常中重新获取控制权,程序还能继续执行
		// recover()
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	arr[i] = 100
	fmt.Println(arr)
}
func test25() {
	fmt.Println("hello world")
}
func main2204() {
	test24(11)
	test25()
}
