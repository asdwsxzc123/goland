package main

import "fmt"

// func add(a, b int) int {
// 	return a + b
// }

// 定义函数类型
// 2. 为已存在的数据类型起别名
type Int int

// func (方法接受者)方法名(参数列表)返回值类型
func (a Int) add(b Int) Int {
	return a + b
}
func main181() {
	// result := add(10, 20)
	var a Int = 10
	value := a.add(20)
	fmt.Println(value)
}

/* 方法的定义和使用 */
type student5 struct {
	name string
	age  int
	sex  string
}

func (s student5) PrintInfo() {
	fmt.Println(s.name)
	fmt.Println(s.age)
	fmt.Println(s.sex)
}

// 通过return修改结构体的数据
func (s student5) EditInfo(name string, age int, sex string) student5 {
	s.name = name
	s.age = age
	s.sex = sex
	s.PrintInfo()
	return s
}

// 通过指针修改结构体的数据
func (s *student5) EditInfo1(name string, age int, sex string) {
	s.name = name
	s.age = age
	s.sex = sex
}
func main182() {
	// var s student5 = student5{"张三", 18, "男"}
	// s.PrintInfo()
	var s1 student5 = student5{"张飞", 28, "男"}
	s1.PrintInfo()
	// s1 = s1.EditInfo("刘备", 31, "男")
	s1.EditInfo1("刘备", 31, "男")
	s1.PrintInfo()
}

type student7 struct {
	name       string
	age        int
	sex        string
	ch_score   int
	math_score int
	en_score   int
}

// func main183() {
// 	stu := student7{"张三", 18, "男", 88, 77, 66}
// 	stu.PrintScore()
// 	stu.SayHello()
// 	stu.EditInfo("李四", 22, "男")
// 	stu.SayHello()
// }

/* 方法的继承 */
type person8 struct {
	id   int
	name string
	age  int
	sex  string
}
type student8 struct {
	// p     person8 // 实名
	person8
	score int
}

func (p *person8) SayHello() {
	fmt.Println("大家好,我是", p.name, p.sex, p.age)
}
func main183() {
	var stu student8
	stu.age = 19
	stu.name = "李四"
	stu.id = 203
	stu.sex = "男"
	stu.score = 88
	// 子结构继承父结构体,允许使用父结构体成员,允许使用父结构体方法
	// 父类不能继承子类
	stu.SayHello()
	// fmt.Println(stu)
}

/* 练习 */
// 1. 记者,我是记者,我的爱好是拍照,年龄是34,我是男狗仔
// 2. 程序猿: 我叫孙权,年龄23,我是男生,工作年限是3年
type person9 struct {
	name string
	age  int
	sex  string
}
type Dep struct {
	person9
	work_time int
}
type Rep struct {
	person9
	hobby string
}

func (p *person9) sayHello() {
	fmt.Printf("大家好,我是%s,我是%s生,我今年%d岁", p.name, p.sex, p.age)
}
func (r *Rep) SayHello() {
	r.person9.sayHello()
	fmt.Println("我的爱好是", r.hobby)
}

// 方法重新,会先调用子类方法,在一个对象中不能使用重复的方法名
func (d *Dep) SayHello() {
	d.person9.sayHello()
	fmt.Println("我的工作年限是", d.work_time)
}
func main184() {
	var r Rep = Rep{person9{"记者", 18, "男"}, "拍照"}
	var d Dep = Dep{person9{"周伟", 28, "男"}, 30}
	r.SayHello()
	d.SayHello()
	d.person9.sayHello()
}
