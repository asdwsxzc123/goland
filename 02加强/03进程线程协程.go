package main

import (
	"fmt"
	"runtime"
	"time"
)

// 1s = 1000ms
// 1ms = 1000us
// 1us =1000ns
// 1. 并行: 借助多核cpu来实现
// 2. 并发:
// 	1. 宏观: 用户体验上,程序在并行执行
// 	2. 微观: 多个任务计划,顺序执行,在飞快的切换,轮换使用cpu时间轮片
// 3. 进程并发: 程序创建进程,fork出来
// 		1. 程序: 编译好的二进制文件,占用磁盘空间
// 		2. 进程: 运行起来的程序,占用系统资源(内存)
// 		3. 进程状态: 初始态,就绪态,运行态,挂起态,终止态
// 		4. linux所有的进程都是init来fock出来的,会出现"孤儿进程"和"僵尸进程"
// 4. 线程并发:
// 	1. LWP: light weight process 轻量级的精彩
// 	2. 进程: 独立地址空间,用于pcb
// 	3. go程: 有独立的pcb,但没有独立的地址空间(共享)
// 	4. 区别: 在于是否共享地址空间,独居(进程),合租(线程)
// 		线程:最小的执行单位
// 		进程: 最小分配单元空间,可以看成是一个线程的进程
// 5. 线程同步
// 	一个线程发出某一个功能调用时,在没有得到结果之前,该调用不返回,同时其他线程为保证数据的一致性,不能调用该功能
// 	为了避免由于时间导致的数据混乱
// 6. 协程并发 轻量级的线程
// 	多个协程分享该线程分配到计算机的资源,提高程序执行的效率
// 	稳定性强,节省资源,效率高

// go语言支持协程,并发goroutine channel
func sing() {
	for i := 0; i < 5; i++ {
		fmt.Println("---正在唱")
		time.Sleep(100 * time.Millisecond)

	}
}
func dance() {
	for i := 0; i < 5; i++ {
		fmt.Println("---正在跳舞")
		time.Sleep(100 * time.Millisecond)
	}
}
func main31() {
	// 子go程才能运行
	// 共同争夺cpu时间
	go sing()
	go dance()
	// 主go程必须执行
	// 主go程结束,子go程也结束
	for {
	}
}

// runtime.gosched() 出让当前go程所占用的cpu时间片
func main32() {
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("this is goroutine test")
			// time.Sleep(100 * time.Millisecond)
		}
	}()
	for {
		runtime.Gosched() // 出让当前cpu时间,将go程让子go程先执行
		fmt.Println("this is main test")
		// time.Sleep(100 * time.Millisecond)
	}
}
func test() {
	defer fmt.Println("ccc")
	// 退出当前go程
	runtime.Goexit()
	fmt.Println("dddd")

}

// runtime.Goexit
func main33() {
	go func() {
		fmt.Println("aaa")
		test()
		fmt.Println("bbbb")
	}()
	for {
	}
}

// 设置cpu数量
func main34() {
	runtime.GOMAXPROCS(1)
}
