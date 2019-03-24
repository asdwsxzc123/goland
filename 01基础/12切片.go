package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 建议使用切片来代替数据

// 切片不定长,在堆区;数组是定长的,在栈区
func main121() {
	// 切片 var 切片名 []数据类型
	// var s []int
	// 自动推到类型,make([]数据类型,5)
	s := make([]int, 3)
	s[0] = 123
	s[1] = 2
	s[2] = 3
	// s[6] = 3 // 下标越界
	// 通过append添加切片信息
	s = append(s, 456)

	// 	遍历操作
	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(s[i])
	// }
	for _, v := range s {
		fmt.Println(v)
	}
	// 容量大于等于长度
	// 如果整体数据没有超过1024字节,每次拓展为上一次的倍数,超过1024每次拓展上一次的1/4
	// cap后面的数据没有进行初始化,不能作为数据遍历的条件
	fmt.Println(cap(s))
	fmt.Println(s)
}

// 切片操作
func main122() {
	s := []int{1, 2, 3, 4, 5}
	// 不包括起始位置,包括结束位置
	// slice := s[2:] // 3,4,5
	// slice := s[:4] // 1,2,3,4
	// slice := s[2:5] // 3,4,5
	// 切片名[起始:结束:容量]
	// slice := s[2:4:5] // 3,4
	slice := s[:] // 1,2,3,4,5
	fmt.Println(slice)

}

func main123() {
	var s []int
	s = append(s, 1, 2, 3)
	s2 := make([]int, 5)
	// 拷贝后两个是独立空间
	copy(s2, s)
	fmt.Println(s2)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}

func bubbleSort1(s []int) []int {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}
func main124() {
	s := []int{4, 21, 5, 2, 3, 6}
	// 切片作为函数参数
	arr := bubbleSort1(s)
	fmt.Println(arr)
}

// 猜数字
func main125() {
	// 创建随机数种子
	rand.Seed(time.Now().UnixNano())
	// 生成随机数
	randNum := rand.Intn(900) + 100
	random := make([]int, 3)
	random[0] = randNum / 100
	random[1] = randNum / 10 % 10
	random[2] = randNum % 10

	// random := make([]int, 3)
	// random[0] = rand.Intn(9) + 1
	// random[1] = rand.Intn(10)
	// random[2] = rand.Intn(10)
	var num int
	var flag int = 0
	usernum := make([]int, 3)
	for {

		for {
			fmt.Println("请输入一个三位数")
			fmt.Scan(&num)
			if num >= 100 && num <= 999 {
				break
			}
			fmt.Println("输入错误,请重新输入")
		}
		usernum[0] = num / 100
		usernum[1] = num / 10 % 10
		usernum[2] = num % 10
		for i := 0; i < 3; i++ {
			if usernum[i] > random[i] {
				fmt.Printf("您输入的第%d位数太大了\n", i+1)
			} else if usernum[i] < random[i] {
				fmt.Printf("您输入的第%d位数太小了\n", i+1)
			} else {
				fmt.Printf("您输入的第%d位数相同\n", i+1)
				flag++
			}

		}
		if flag == 3 {
			break
		} else {
			flag = 0
		}
	}
}
