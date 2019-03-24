package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main2301() {
	// 文件的路径,绝对路径和相对路径
	// os.Create创建文件时,文件不存在会创建一个新的文件,如果存在,会覆盖原有内容
	fp, err := os.Create("./test.txt")
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	// 关闭文件
	defer fp.Close()

	fmt.Println("文件创建成功")
	// 写入文件
	// \n在linux可以换行,windows换行需要\r\n
	n, _ := fp.WriteString("hello\r\n")
	fmt.Println(n)
	// 1个汉字是3个字符
	n, _ = fp.WriteString("你好")
	fmt.Println(n)
}

// 切片写入
func main2302() {
	fp, err := os.Create("./test.txt")
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	defer fp.Close()
	// 1. 字符切片写入文件
	// b := []byte{'h', 'e', 'l', 'l', 'o'}
	// 2. 字符串转化成字符切片写入
	str := "静夜思"
	b := []byte(str)
	fp.Write(b)
}

func main2304() {
	// 创建文件
	// fp, err := os.Create("./test.txt")

	// Open只读方式打开文件
	// os.Open("./test.txt")

	// openFile不能创建新文件,只能打开文件
	// O_RDONLY, O_WDONLY,O_RDWR,O_APPEND(追加模式)
	// 第三个参数: 打开的权限
	fp, err := os.OpenFile("./test.txt", os.O_RDWR, 6)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	b := []byte("abcd")
	// 偏移量
	ret, _ := fp.Seek(0, io.SeekEnd)
	fmt.Println(ret)
	// 只会从偏移开始,不会覆盖超过输入的内容
	fp.WriteAt(b, ret)
}

/* 读取文件 */
func main2305() {
	fp, err := os.Open("./test.txt")
	/*
		文件打开失败
		1. 文件不存在
		2. 文件权限
		3. 文件打开上限
	*/
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	/* 块读写 */
	// b := make([]byte, 1024)
	// fp.Read(b)
	// // for i := 0; i < len(b); i++ {
	// // 	fmt.Printf("%c", b[i])
	// // }
	// fmt.Println(string(b))

	/* 按行读取 */
	// 创建缓存区
	// r := bufio.NewReader(fp)
	// // 读取一行内容
	// // 如何在文件截取中,没有标志位(分隔符)到文件末尾自动终止,EOF -1文件结束标志
	// b, _ := r.ReadBytes('\n')
	// fmt.Println(string(b))
	// b, _ = r.ReadBytes('\n')
	// fmt.Println(string(b))

	/* 通过判断来读取信息 */
	b := make([]byte, 20)
	for {
		// 通过文件返回值 个数和错误信息
		n, err := fp.Read(b)
		if err == io.EOF {
			break
		}
		fmt.Print(string(b[:n]))
	}
}

// 行读写
func main2306() {
	fp, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()
	// 读取行数
	r := bufio.NewReader(fp)
	for {
		b, err := r.ReadBytes('\n')
		fmt.Print(string(b))
		if err == io.EOF {
			break
		}
	}
}

/* 文件拷贝 */
func main2307() {
	fp1, err1 := os.Open("./test/2977436563-559628a35a110.png")
	fp2, err2 := os.Create("./test/1.png")
	if err1 != nil || err2 != nil {
		fmt.Println("拷贝失败")
		return
	}
	defer fp1.Close()
	defer fp2.Close()
	// 使用块读写
	b := make([]byte, 1024)
	for {
		n, err := fp1.Read(b)
		if err == io.EOF {
			break
		}
		fp2.Write(b[:n])
	}
}
