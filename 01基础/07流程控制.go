package main

import "fmt"

/* 选择结构 */
// 顺序结构,选择结构,循环结构
func main70() {
	var score int
	fmt.Scan(&score)
	if score > 700 {
		fmt.Println("我要上清华")
	} else if score > 600 {
		fmt.Println("我要上一本")
	} else {
		fmt.Println("我要上蓝翔")
	}

}

func main71() {
	var w int
	fmt.Scan(&w)
	switch w {
	case 1:
		// fmt.Println("星期一")
		fallthrough // 让switch执行下一个分支的代码
	case 2:
		fmt.Println("星期2")
	case 3:
		fmt.Println("星期3")
	case 4:
		fmt.Println("星期4")
	default:
		fmt.Println("输入错误")
	}

}
