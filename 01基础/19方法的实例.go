package main

import "fmt"

// 18方法.go调用
// 方法1 打招呼
func (s *student7) SayHello() {
	fmt.Printf("大家好,我叫%s,今年%d岁,%s\n", s.name, s.age, s.sex)
}

// 方法2 打印成绩
func (s *student7) PrintScore() {
	sum := s.ch_score + s.math_score + s.en_score
	fmt.Printf("总成绩为%d,语文成绩为%d,数学成绩%d,英语成绩%d\n", sum, s.ch_score, s.math_score, s.en_score)
}

// 方法3 修改数据
func (s *student7) EditInfo(name string, age int, sex string) {
	s.name = name
	s.age = age
	s.sex = sex
}

func SayHello() {
	fmt.Println("hello")
}
