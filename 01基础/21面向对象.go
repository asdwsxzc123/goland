package main

import "fmt"

/*
// 设计模式,面向对象基于M(模型)V(视图)C(控制器)有26种设计模式
// 工厂模式,定义结构体
type OptFractory struct {
}

// 基于继承 方法 接口 多态和设计模式
func (of *OptFractory) OptCalc(num1, num2 int, op string) (value int) {
	// 通过运算符进行分类计算
	var opter Opter
	switch op {
	case "+":
		var add addOpt = addOpt{Opt{num1, num2}}
		// value = add.Operate()
		opter = &add
	case "-":
		var sub SubOpt = SubOpt{Opt{num1, num2}}
		opter = &sub
	}
	// value = opter.Operate()
	value = Fratory(opter)
	return value
}

type Opter interface {
	Operate() int
}

func Fratory(o Opter) (value int) {
	value = o.Operate()
	return value
}
func main2101() {
	var optf OptFractory
	value := optf.OptCalc(10, 20, "+")
	fmt.Println(value)
} */

// 完整的计算器实现
// 定义结构体
type FuLei struct {
	num1 int
	num2 int
}
type Jiafa struct {
	FuLei
}
type Jianfa struct {
	FuLei
}

// 定义接口
type Jiekou interface {
	Jisuan() int
}

// 定义方法
func (f *Jiafa) Jisuan() int {
	return f.num1 + f.num2
}
func (f *Jianfa) Jisuan() int {
	return f.num1 - f.num2
}

// 创建工厂函数分发数据
type Gongchang struct{}

func (f *Gongchang) init(num1, num2 int, op string) (value int) {
	var oper Jiekou
	switch op {
	case "+":
		oper = &Jiafa{FuLei{num1, num2}}
		value = oper.Jisuan()
	case "-":
		oper = &Jianfa{FuLei{num1, num2}}
		value = oper.Jisuan()
	}
	return value
}

func main2102() {
	var gc Gongchang
	value := gc.init(20, 10, "+")
	fmt.Println(value)
}
