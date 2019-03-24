package main

import (
	"fmt"
)

// 指针类型定义
func main151() {
	var a int = 10
	// 定义整型指针变量, 指向a的地址
	// var 指针 *数据结构 一级指针
	var p *int
	// 将a的地址赋值个指针变量p
	p = &a
	a = 122
	// *p拿到变量对应的值
	fmt.Println(*p) // 122
	fmt.Println(p)  // 0xc000014080
	fmt.Println(&a) // 0xc000014080

}

func main152() {
	// a := 10
	// p := &a
	// fmt.Println(p)

	// 这个是指向nil(0)空指针
	// 0-255为系统占用,不允许用户进行读写操作
	var p *int
	fmt.Println(p)
}

func main153() {
	var p *int
	// 使用new创建数据类型,不需要管理空间的释放,在堆区创建空间
	p = new(int)
	*p = 10
	fmt.Println(*p)
}

// 指针作为参数,是地址传递参数
func main154(a *int, b *int) {
	*a, *b = *b, *a
}
func main155() {
	a := 10
	b := 20
	main154(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}

// 数组指针定义
func main156() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%p\n", &arr)
	fmt.Printf("%p\n", &arr[0])
	// 定义指针数组
	var p *[5]int
	p = &arr
	// 运算符优先级问题,需要先加括号
	(*p)[0] = 2
	fmt.Println(*p)
	// 数组指针不能直接打印
	fmt.Printf("%p\n", arr)
	fmt.Printf("%p\n", p)
}

// 指针可以求长度,也可以通过下标获取汉字
func main157() {
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	p := &arr
	fmt.Println(len(p))
	for i := 0; i < len(p); i++ {
		fmt.Println(p[i])
	}
}

/* 切片指针 */
func main158() {
	// 切面名本身就是地址
	var slice []int = []int{1, 2, 3, 4, 5}
	// 切片本身就是一个地址,因此切片地址的赋值是重新赋值
	// 这里的p是二级指针
	p := &slice
	fmt.Printf("%p\n", p)
	fmt.Printf("%p\n", slice)
	// 本身是两层地址,p的地址指向slice的地址,然后在去找到堆地址
	(*p)[2] = 200
	fmt.Println(slice)

}

func main159(s []int) {
	// 切片创建了新的指针
	s[0] = 22
	// return s
}
func main1510(s *[]int) {
	*s = append(*s, 4, 5, 6)
}
func main1511() {
	s := []int{1, 2, 3}
	// main159(s)
	fmt.Printf("%p\n", s)
	// 指针地址发生了改变
	main1510(&s)
	fmt.Printf("%p\n", s)
	fmt.Println(s)
}

/* 创建切片指针 */
func main1512() {
	var p *[]int
	fmt.Printf("%p\n", p)
	p = new([]int)
	fmt.Printf("%p\n", p)
	*p = append(*p, 1)
	for i := 0; i < len(*p); i++ {
		fmt.Println((*p)[i])
	}
	for _, val := range *p {
		fmt.Println(val)
	}
}

/* 指针数组 */
func main1513() {
	var arr [3]*int
	a := 10
	b := 20
	c := 30
	arr[0] = &a
	arr[1] = &b
	arr[2] = &c
	fmt.Println(*arr[1])
	fmt.Println(arr)
}

/* 指针切片 */
func main1514() {
	var slice []*int
	a := 10
	b := 20
	c := 30
	slice = append(slice, &a, &b, &c)
	fmt.Println(slice)
}

/* 结构体指针 */
type student1 struct {
	id   int
	name string
	age  int
	sex  string
}

func main1515() {
	// 定义结构体变量
	var stu student1 = student1{1, "王五", 14, "男"}
	fmt.Println(stu)
	// 定义结构体指针指向变量的地址
	p := &stu
	// 结构体地址和首元素地址一样
	fmt.Printf("%p\n", p)
	fmt.Printf("%p\n", &stu.id)

	// 指针可以直接操作结构体成员
	p.name = "李四"
	fmt.Println(stu)
}

/* 结构体切片 */
func main1516() {
	var stu []student1 = make([]student1, 3)
	// 结构体切片指针
	p := &stu
	fmt.Printf("%p\n", p)
	// *p = append(*p, student1{3, "飞机", 15, "男"})
	(*p)[0] = student1{3, "飞机", 15, "男"}
	fmt.Println(stu)
}

// 多级指针
func main1517() {
	a := 10
	p := &a
	// 二级指针存储一级指针的地址
	pp := &p
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", p)
	fmt.Printf("%T\n", pp)
}
