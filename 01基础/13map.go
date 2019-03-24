package main

import "fmt"

func main131() {
	// 字典结构 map[keyType]valueType,没有具体的顺序
	// var m map[int]string
	// map是无序存储的,自动扩充
	// m := make(map[int]string, 1)
	// m[1] = "张三"
	// m[2] = "张三1"
	// m[3] = "张三2"
	// fmt.Println(m)
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	// make定义的可以覆盖
	// m := make(map[string]int, 1)
	// m["张三"] = 1
	// m["张三1"] = 2
	// m["张三1"] = 4

	// map在定义是key是唯一的,不允许重复
	// m := map[string]int{"张三": 3, "张三1": 1, "张三2": 2}
	// fmt.Println(m)
	// for k, v := range m {
	// 	fmt.Println(k, v)
	// }

	m := map[int]string{1: "张三", 2: "张三1", 3: "张三2"}
	v, ok := m[1]
	// 可以验证是否需要赋值
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("不存在")
	}
}

// 删除map元素
func main132() {
	m := map[int]string{1: "刘备", 2: "关羽", 3: "李逵"}
	// fmt.Printf("%T\n", m)
	// 只能用来删除map
	delete(m, 1)
	fmt.Println(m)
}

// map在函数中是引用传递
func main133(m map[int]string) {
	m[1] = "杨戬"
	m[4] = "李靖"
}
func main134() {
	m := map[int]string{1: "刘备", 2: "关羽", 3: "李逵"}
	main133(m)
	fmt.Println(m)
}
