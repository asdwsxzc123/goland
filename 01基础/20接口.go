package main

// 继承:继承自拖拉机,实现了扫地接口
// 封装: 无序知道如何工作,开动即可
// 多态: 平时扫地,天热当风扇
// 重用: 没有二外动力,重复利用了发动机能量
// 多线程: 多个扫把同时工作
// 低耦合:扫把改为拖把而无需改动

import "fmt"

type Opt struct {
	num1 int
	num2 int
}

type addOpt struct {
	Opt
}
type SubOpt struct {
	Opt
}
type Operate interface {
	Operate() int
}

func (add *addOpt) Operate() int {
	return add.num1 + add.num2
}
func (sub *SubOpt) Operate() int {
	return sub.num1 - sub.num2
}
func main201() {
	// var add addOpt
	// add.num1 = 10
	// add.num2 = 20
	// sum := add.Operate()
	// fmt.Println(sum)

	// var sub SubOpt
	// sub.num1 = 20
	// sub.num2 = 10
	// value2 := sub.Operate()
	// fmt.Println(value2)

	var opt Operate
	opt = &addOpt{Opt{10, 20}}
	value := opt.Operate()
	opt = &SubOpt{Opt{10, 20}}
	value = opt.Operate()
	fmt.Println(value)
}

/* 接口 */
// 需要一个能理清所有业务的架构师来定义一些主要的接口
type person10 struct {
	name string
	sex  string
	age  int
}
type student9 struct {
	person10
	score int
}
type teacher9 struct {
	person10
	subject string
}

func (s *student9) SayHello() {
	fmt.Printf("我是学生,我叫%s,年龄%d,性别%s,成绩为%d\n", s.name, s.age, s.sex, s.score)
}
func (s *teacher9) SayHello() {
	fmt.Printf("我是老师,我叫%s,年龄%d,性别%s,交的学科%s\n", s.name, s.age, s.sex, s.subject)
}

// 接口调用
type Humaner interface {
	// 方法 函数声明 没有具体实现 根据对象的不同,实现方式也不同
	SayHello()
}

func main202() {
	t := teacher9{person10{"刘老师", "男", 28}, "中文"}
	s := student9{person10{"王同学", "女", 18}, 88}
	// t.SayHello()
	// s.SayHello()
	var h Humaner
	h = &t
	h.SayHello()
	h = &s
	h.SayHello()
}

/* 多态 */
type person11 struct {
	name string
	sex  string
	age  int
}
type student11 struct {
	person11
	score int
}
type teacher11 struct {
	person11
	subject string
}

type Personer interface {
	SayHello()
}

// 多态是将接口类型作为函数参数,实现了接口的统一处理
func sayHello(p Personer) {
	p.SayHello()
}

func (s *student11) SayHello() {
	fmt.Printf("我是学生,我叫%s,年龄%d,性别%s,成绩为%d\n", s.name, s.age, s.sex, s.score)
}
func (s *teacher11) SayHello() {
	fmt.Printf("我是老师,我叫%s,年龄%d,性别%s,交的学科%s\n", s.name, s.age, s.sex, s.subject)
}
func main203() {
	var p Personer
	p = &student11{person11{"王同学", "女", 18}, 88}
	sayHello(p)
	p = &teacher11{person11{"刘老师", "女", 18}, "英语"}
	sayHello(p)
}

// 多态练习
// 将移动硬盘和u盘或MP3插入到电脑上进行读写数据(分析类,接口,方法)
// 创建对象
type USBDev struct {
	id     int
	name   string
	rspeed int
	wspeed int
}
type Mobile struct {
	USBDev
	call string
}
type UDisk struct {
	USBDev
	storage string
}

// 接口
type usber interface {
	Read()
	Write()
}

// 方法
func (f *Mobile) Read() {
	fmt.Printf("%s正在读取数据,速度为%d,还能%s\n", f.name, f.rspeed, f.call)
}
func (f *Mobile) Write() {
	fmt.Printf("%s正在写入数据,速度为%d,还能%s\n", f.name, f.wspeed, f.call)
}
func (f *UDisk) Read() {
	fmt.Printf("%s正在读取数据,速度为%d,还能%s\n", f.name, f.rspeed, f.storage)
}
func (f *UDisk) Write() {
	fmt.Printf("%s正在写入数据,速度为%d,还能%s\n", f.name, f.wspeed, f.storage)
}

// 多态
func usb(u usber) {
	u.Read()
	u.Write()
}
func main204() {
	// 声明接口
	var u usber
	u = &Mobile{USBDev{1, "手机", 5, 10}, "打电话"}
	usb(u)
}

/* 接口的继承,接口的命名以er作为结尾 */
type myPersoner interface { // 子集
	sayHello()
}
type Childrener interface { // 超集
	myPersoner
}

/* 空接口 */
func main205() {
	// 空接口,可以接受任意类型数据,接口定义的空间是固定
	// var i interface{}
	// i = 10
	// fmt.Println(i)
	// i = "hello"
	// fmt.Println(i)

	// 空接口切片
	var i1 []interface{}
	i1 = append(i1, 1, 2, "hello")
	for _, val := range i1 {
		fmt.Println(val)
	}
}

/* 数据断言 */

func main206() {
	arr := make([]interface{}, 4)
	arr[0] = 123
	arr[1] = 3.14
	arr[2] = "hello"
	arr[3] = []int{1, 2, 3}
	for _, v := range arr {
		// 对数据进行类型断言
		if data, ok := v.(int); ok {
			fmt.Println("整型", data)
		} else if data, ok := v.(float64); ok {
			fmt.Println("浮点型", data)
		} else if data, ok := v.(string); ok {
			fmt.Println("字符串", data)
		} else if data, ok := v.([]int); ok {
			fmt.Println("切片", data)
		}
	}
}
