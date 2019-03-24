package main

import "fmt"

type person1 struct {
	name string
	age  int
	sex  string
}
type student2 struct {
	// 通过匿名字段实现继承操作
	person1
	id    int
	score int
}

func main1710() {
	var stu student2
	stu.id = 1
	stu.person1.name = "战三"
	// stu.name = "战三"
	stu.score = 99
	stu.sex = "男"
	stu.age = 18
	fmt.Println(stu)
}
func main1711() {
	var stu student2 = student2{person1{"张三丰", 18, "男"}, 11, 80}
	fmt.Println(stu)
}

// 同名字段
type person3 struct {
	name string
	age  int
	sex  string
}
type student3 struct {
	person3
	id    int
	name  string
	score int
}

func main1712() {
	var stu student3 = student3{person3{"张三丰", 18, "男"}, 11, "疯子", 80}
	stu.name = "疯子2"
	stu.person3.name = "李四"
	fmt.Println(stu)
}

/* 指针匿名字段 */
type person4 struct {
	name string
	age  int
	sex  string
}
type student4 struct {
	*person4
	id    int
	score int
}

func main1713() {
	var stu student4
	// stu.person4 = new(person4)
	// stu.name = "郭襄"
	// stu.id = 102
	// stu.score = 88
	// stu.person4.name = "郭小姐"
	var per = person4{"杨过", 35, "男"}
	stu.person4 = &per
	fmt.Println(stu.name)
	fmt.Println(stu)
}

/* 多重继承 */
type TestA struct {
	name string
	id   int
}
type TestB struct {
	TestA
	sex string
	age int
}
type TestC struct {
	TestB
	score int
}

func main1714() {
	var s TestC
	s.name = "李四"
	fmt.Println(s)
}
