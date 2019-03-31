package main

import (
	"fmt"
	"time"
)

/* channel 管道*/
// 用来解决协程的同步问题和协程之间数据共享(数据传递)
// goroutine运行在相同的地址空间,访问共享内存必须做好同步,gorutine通过通信来共享内存,而不是共享内存来通信
// FIFO 栈, FIFO管道
// capacity 0 无缓冲区
// make(chan Type, capacity)
// make(chan int)
// 每当有一个进程启动时,系统会自动打开三个文件,标准输入,标准输出,标准错误---对应三个文件:stdin,stdout,stderr

// channel有两个端
// 1. 写端(传入端) chan <-
// 2. 读端(传出端) <-chan
// 要求读端和写端必须同时满足条件,读端有数据可读,写端有数据可写,才能在chan上进行数据流动,否则,则阻塞

// 全局定义channel,用来完成数据同步
var channel = make(chan int)

// 易主
func printer(s string) {
	for _, ch := range s {
		fmt.Printf(string(ch)) // stdout
		time.Sleep(300 * time.Millisecond)
	}
}

// 定义两个人使用打印机
func person1() {
	printer("hello")
	// 阻塞
	channel <- 8
}
func person2() {
	// 获取阻塞后的接受
	<-channel
	printer("world")

}
func main41() {
	go person1()
	go person2()
	for {
	}
}

// 无缓冲channel
// 读端和写端必须同时在线,channel的同步
func main42() {
	ch := make(chan string)
	// 死锁
	// ch <- "子进程打印完毕"
	// fmt.Println(len(ch), cap(ch)) // 0, 0
	go func() {
		for i := 0; i < 2; i++ {
			fmt.Println(i)
		}
		ch <- "子进程打印完毕"
	}()
	str := <-ch
	fmt.Println(str)
}

func main43() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			ch <- i
		}
	}()
	// 主要chennel没有被接受,就会阻塞进程
	time.Sleep(time.Second * 2)
	for i := 0; i < 5; i++ {
		num := <-ch
		fmt.Println("住", num)
	}
}

/* 有缓存channel */

func main44() {
	// 容量和长度非0,应用于两个go程中,一个读,一个写
	// 缓冲区可以进行数据存储,存储至容量上限,阻塞.具备异步能力
	// 长度为0,容量为3
	ch := make(chan int, 3)
	fmt.Println(len(ch), cap(ch))
	go func() {
		for i := 0; i < 15; i++ {
			ch <- i
			fmt.Println("go", i)
			// 长度增加
			// fmt.Println(len(ch), cap(ch))
		}
	}()
	for i := 0; i < 15; i++ {
		num := <-ch
		fmt.Println(num)
	}
}

// 关闭channel,使用close(ch),
// 1.确定不再相对端发送,关闭;ok = true
// 2. 已经关闭的channel不能再发送数据,会报错
// 3. 写端已经关闭channel,可以从中读取数据,num为0;如果是有缓存的channel,如果缓冲区有数据,先读数据,没有后可以读到0
func main45() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 8; i++ {
			ch <- 1
		}
		close(ch)
	}()
	for {
		if num, ok := <-ch; ok == true {
			fmt.Println(num)
		} else {
			// channel关闭,ok == false
			num, ok := <-ch
			fmt.Println(ok, num)

			break
		}
	}
}

// 双向channel,可以隐式转化为任意一种单向channel
// 单向channel不能转化为双向channel
func main46() {
	// channel需要在go程中使用
	ch := make(chan int)
	var sendCh chan<- int = ch
	sendCh <- 700
	var readCh <-chan int = ch
	num := <-readCh
	fmt.Println(num)

	// 反向赋值,会报错
	// var ch2 chan int = sendCh

}
func send(out chan<- int) {
	out <- 80
	close(out)
}
func recv(in <-chan int) {
	num, ok := <-in
	fmt.Println(num, ok)
}
func main47() {
	ch := make(chan int)
	go func() {
		send(ch) // 双向channel转化为单向
	}()
	recv(ch)
}
