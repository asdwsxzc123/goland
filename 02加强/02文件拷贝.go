package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/* 拷贝文件 */
func main21() {
	// 读取文件
	f_r, err_r := os.Open("./01内存地址.md")
	// 创建文件
	f_w, err_w := os.Create("./test.md")
	if err_r != nil || err_w != nil {
		fmt.Println("打开文件失败")
		return
	}
	defer f_r.Close()
	defer f_w.Close()
	// 从读文件获取数据
	b := make([]byte, 4*1024)
	// 循环读取数据
	for {
		n, err := f_r.Read(b)
		if err != nil && err == io.EOF {
			fmt.Println("读取完成")
			return
		}
		f_w.Write(b[:n])
	}

}

/* 目录项 */
// 判断目录中是否有.png文件
func main22() {
	// OpenFile(path,读写模式,打开权限) os.ModeDir
	f, err := os.OpenFile("./test", os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("文件读取失败")

		return
	}
	defer f.Close()
	info, err := f.Readdir(-1)
	if err != nil {
		fmt.Println(err)
	}
	for _, fileInfo := range info {
		if fileInfo.IsDir() {
			fmt.Println(fileInfo.Name(), "是目录")
		} else {
			// fmt.Println(fileInfo.Name(), "是文件")
			fileName := fileInfo.Name()
			if strings.HasSuffix(fileName, ".png") {
				fmt.Println(fileName, "是图片")
			}
		}
	}
}

func copy2Dir(fileName string) {
	source_path := "./test/"
	target_path := "./img1/"
	f_r, err := os.Open(source_path + fileName)
	f_w, err_w := os.Create(target_path + fileName)
	if err != nil {
		fmt.Println("文件获取失败")
	}
	if err_w != nil {
		fmt.Println("文件创建失败")
	}
	defer f_r.Close()
	defer f_w.Close()
	buf := make([]byte, 1024*4)
	for {
		n, err := f_r.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println(err)
			fmt.Println("文件写入完成")
			break
		}
		f_w.Write(buf[:n])

	}

}

// 从目录中拷贝.png文件到指定目录
func main23() {
	// OpenFile(path,读写模式,打开权限) os.ModeDir
	source_path := "./test/"
	suffix := ".png"
	f, err := os.OpenFile(source_path, os.O_RDONLY, os.ModeDir)
	if err != nil {
		fmt.Println("文件读取失败")
		return
	}
	defer f.Close()
	files, err := f.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), suffix) {
			fmt.Println(file.Name())
			copy2Dir(file.Name())
		}
	}
}

func readTxt(path string, count *int) {
	f_r, err_r := os.Open(path)
	if err_r != nil {
		fmt.Println(err_r)
		return
	}
	defer f_r.Close()
	// 按行读取,建立缓存区
	r := bufio.NewReader(f_r)
	for {
		buf, err := r.ReadBytes('\n')
		rowLine := strings.Fields(string(buf))
		for _, val := range rowLine {
			if val == "love" {
				*count++
			}
		}
		if err != nil || err == io.EOF {
			fmt.Println("读取完成")
			return
		}

	}
}

// 统计目录内单词出现的次数
func main24() {
	source_path := "./test/"
	f_r, err_r := os.OpenFile(source_path, os.O_RDONLY, os.ModeDir)
	if err_r != nil {
		fmt.Println("读取失败")
	}
	defer f_r.Close()
	// 打开多个txt,循环读取一行,将一行的字符串拆分成切片,fields,变量字符串,统计"love"出现的次数
	files, err := f_r.Readdir(-1)
	if err != nil {
		fmt.Println(err)
	}
	count := 0
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".txt") {
			fmt.Println(file.Name())
			readTxt(source_path+file.Name(), &count)
		}
	}
	fmt.Println(count)

}
