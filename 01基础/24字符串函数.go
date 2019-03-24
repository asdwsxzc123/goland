package main

import (
	"fmt"
	"strconv"
)

func main2401() {
	//  1.Contains查找
	/* str := "hello world"
	value := strings.Contains(str, "llo") // true
	if value {
		fmt.Println("出现")
	} else {
		fmt.Println("未出现")
	} */

	// 2.Join拼接
	/* s := []string{"13", "dsf"}
	value := strings.Join(s, " ") // 13 dsf
	*/

	// 3.Index查找位置
	// str := "你好,在吗"
	// // 1个汉字是3个字符
	// value := strings.Index(str, "好") // 3

	// 4.Repeat 重复
	// str := "你好,你好呀"
	// value := strings.Repeat(str, 2) // 你好,你好呀你好,你好呀

	// 5.Replace 替换
	// str := "hello world"
	// // -1表示全部替换
	// value := strings.Replace(str, "l", "2", -1) // he22o world

	// 6.Split 分割成切片
	// str := "www.it.com"
	// value := strings.Split(str, ".")

	// 7. Trim 删除两边的字符串
	// str := "===www.==it.com=="
	// value := strings.Trim(str, "=")

	// 8.Fields 去掉字符串中的空格,并返回切片
	// str := "    www. it.com  "
	// value := strings.Fields(str)

	// 9.Format,吧其他类型转化成字符串
	// 将bool类型转化成字符串
	// value := strconv.FormatBool(false)
	// 将整型转化成字符串
	// value := strconv.FormatInt(123, 10)

	// 10.Parse将字符串转化成其他类型
	// 字符串的true转化成true,其他多东西都会失败
	// str := "false"
	// value, err := strconv.ParseBool(str)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(value)
	// }

	// 将字符串转化成整型
	// str := "123"
	// value, _ := strconv.ParseInt(str, 10, 12)
	// fmt.Println(value)

	// 11.Append追加
	b := make([]byte, 0, 1024)
	value := strconv.AppendBool(b, false)
	value = strconv.AppendInt(value, 123, 10)
	value = strconv.AppendQuote(value, "hello")
	fmt.Println(string(value))

}
