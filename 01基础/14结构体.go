package main

import "fmt"

// 定义结构体
type Student struct {
	id   int
	name string
	sex  string // 字符类型
	age  int
	addr string
}

// 定义结构体
func main141() {
	// var s Student
	// s.id = 101
	// s.name = "张飞"
	// s.sex = "男"
	// s.age = 18
	// s.addr = "北京燕郊"
	// fmt.Println(s)

	// var s Student = Student{102, "关羽", "男", 19, "山西运城"}

	s := Student{102, "关羽", "男", 19, "山西运城"}
	fmt.Println(s)
}

// 在数组中
func main142() {
	var arr [1]Student
	// arr[0].id = 14
	fmt.Println(len(arr))
	for i := 0; i < len(arr); i++ {
		fmt.Scan(&arr[i].id, &arr[i].name, &arr[i].sex, &arr[i].age, &arr[i].addr)
	}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func main143() {
	// 放在切片中
	arr := []Student{{102, "关羽", "男", 19, "山西运城"}}
	arr = append(arr, Student{101, "曹操", "男", 19, "山西运城"}, Student{})
	fmt.Println(arr)
}

// 作为函数参数
type person struct {
	id    int
	name  string
	score int
	sex   string
}

// 结构体是值传递
func main144(s person) {
	fmt.Println(s.name)
	fmt.Println(s)
	s.name = "李逵"
}
func main145() {
	stu := person{101, "松江", 9, "男"}
	main144(stu)
	fmt.Println(stu)
}

// 切片是引用传递
func main146(stus []person) {
	// stus = append(stus, person{104, "卢俊义", 9, "男"})
	stus[0].name = "晁盖"
}
func main147() {
	stus := []person{person{101, "松江", 9, "男"}, person{102, "吴勇", 9, "男"}}
	// 结构体切片作为地址传递,
	// 结构体数组作为值传递
	stus = append(stus, person{103, "李逵", 9, "男"})
	main146(stus)
	fmt.Println(stus)
}

/*
type skills struct {
	名称
	耗蓝
	cd 冷却时间
	范围
	伤害
}
定义结构体切片 保持技能信息
type role struct {
	名称
	等级 lv
	经验 exp
	钻石
	金币
	生命值 hp
	攻击力
	暴击
	防御
	蓝量mp
}
type 信用卡 struct {
	卡号
	持卡人姓名
	额度
	有效期
	密码
	银行信息
}
type 消费记录 struct {
	卡号
	消费时间
	消费id
	流水号
	消费金额
	备注
}
*/
// 结构体练习,存储5位学生,三门成绩 求出每名学生的总成绩和平均成绩
type student struct {
	id    int
	name  string
	score [3]int
}

func main148() {
	students := []student{student{1, "张三", [3]int{80, 70, 60}}, student{2, "李四", [3]int{82, 50, 30}}, student{3, "王五", [3]int{90, 66, 56}}}
	for _, item := range students {
		var sum int
		for _, val := range item.score {
			sum += val
		}
		fmt.Println(sum)
		fmt.Println(sum / 3)
	}
}
