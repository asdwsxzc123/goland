package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 数组,指针,切片,结构体等复合类型

/* 1. 数组 */
func main111() {
	// 定义数组 var 数组名 [元素个数] 类型
	// var a [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// b := a[2]
	// fmt.Println(b)
	// fmt.Println(len(a))

	// 使用自动类型推导,声明数组
	// a := [...]int{1, 2, 3}

	// range 遍历集合信息
	// 返回下标和值
	// for i, v := range a {
	// 	fmt.Println(i, v)
	// }

	arr := [5]int{1, 2, 3, 4, 5}
	// arr[5] = 100 // 数组下标越界
	// arr[-1] = 20 // 数组下标越界
	// 数组在定义后,元素个数已经固定,不能添加
	// 数组是一个常量,不允许赋值,数组名代表整个数组
	// arr = 10
	fmt.Println(arr)
}

// 找最大值
func main112() {
	arr := [...]int{9, 1, 5, 6, 7, 3, 10, 4, 8}
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	fmt.Println(max)
}

// 交换数组
func main113() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	start := 0
	end := len(arr) - 1
	for {
		if start > end {
			break
		}
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
	fmt.Println(arr)
}

// 冒泡排序
func main114() {
	arr := [...]int{9, 1, 5, 6, 7, 3, 10, 4, 8}
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}

// 数组作为参数是值传递的
func main115(arr [9]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

// 数组作为返回值,
func BubbleSort(arr [9]int) [9]int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
func main116() {
	arr := [...]int{9, 1, 5, 6, 7, 3, 10, 4, 8}
	// main115(arr)
	// 只读变量,只能赋值自己的
	arr = BubbleSort(arr)
	fmt.Println(arr)
}

func main117() {
	// 1. 导入投文件 math/rand time
	// 2. 添加随机数
	// 3. 使用随机数
	rand.Seed(time.Now().UnixNano())
	fmt.Println(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		// 0-99以内的数据
		fmt.Println(rand.Intn(100))
	}
}

func main118() {
	// 1-33 红球6个,不能重复
	rand.Seed(time.Now().UnixNano())
	var redball [6]int
	for i := 0; i < len(redball); i++ {
		// 遍历之前存在的值和新随机数是否重复
		temp := rand.Intn(33) + 1
		for j := 0; j < i; j++ {
			if temp == redball[j] {
				temp = rand.Intn(33) + 1
				j = -1
				continue
			}
		}
		redball[i] = temp
	}
	fmt.Println(redball)
}

/* 二维数组 */
func main119() {
	// var 数组名 [行个数][列个数] 数组类型
	var arr [3][4]int
	arr[1][2] = 1
	fmt.Println(arr)

	var arr1 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr1)
}
