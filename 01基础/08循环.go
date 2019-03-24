// for: 表达式1 表达式2 表达式3 {
// 	循环体
// }
// 表达式1: 定义一个循环的变量,记录循环的次数
// 表达式2: 循环条件,循环多少次
// 表达式3: 循环条件的代码,是循环条件不再成立
package main

import "fmt"

func main81() {
	for i := 1; i <= 10; i++ {
		fmt.Println("hello world")

	}
}
func main82() {
	// i := 1
	// for {
	// 	if i > 10 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }
	// fmt.Println(i)

	for {
		fmt.Println("hello")
		break
	}
}

/* 水仙花 一个三位数100-999,每个位数的立方和等于这个数本身 153*/
func main83() {

	for i := 100; i < 1000; i++ {
		// 百位
		a := i / 100
		b := i % 100 / 10
		c := i % 10
		if a*a*a+b*b*b+c*c*c == i {
			fmt.Println(i)
		}
	}

}

/* 含有7或被7整除 */
func main84() {
	for i := 1; i < 101; i++ {
		if i%7 == 0 || i%10 == 7 || i/10 == 7 {
			fmt.Println(i)
		}
	}
}

// 百钱白鸡,公鸡5钱,母鸡3钱,3只小鸡1钱
func main85() {
	count := 0
	// 方法一
	/* for cock := 0; cock <= 20; cock++ {
		for hen := 0; hen <= 33; hen++ {
			for chicken := 0; chicken <= 100; chicken++ {
				count++
				if cock+hen+chicken == 100 && cock*5+hen*3+chicken/3 == 100 {
					fmt.Printf("%d公鸡,%d母鸡, %d小鸡\n", cock, hen, chicken)
				}
			}
		}
	} */
	// 优化
	for cock := 0; cock <= 20; cock++ {
		for hen := 0; hen <= 33; hen++ {
			count++
			chicken := 100 - cock - hen
			if cock*5+hen*3+chicken/3 == 100 && chicken%3 == 0 {
				fmt.Printf("%d公鸡,%d母鸡, %d小鸡\n", cock, hen, chicken)
			}
		}
	}
	fmt.Println("执行次数:", count)
}
