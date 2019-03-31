package main

import (
	"fmt"
	"runtime"
	"time"
)

// select可以监听channel的数据流动
// 每隔case语句必须是一个IO操作
// select{
// 	case <- chan1;成功读到数据
// 	case chan2 <- 1;成功写入数据
// 	break // 只能跳出select
// 	default;
// }
func main71() {
	ch := make(chan int)    // 用来进行数据通信的channel
	quit := make(chan bool) // 用来判断是否退出的channel
	go func() {             // 写数据
		for i := 0; i < 8; i++ {
			ch <- i
			time.Sleep(time.Second * 1)
			fmt.Println("写入", i)
		}
		close(ch)
		quit <- true // 通知主go程结束
		runtime.Goexit()
	}()
	for { // 主go程 读数据
		select {
		case num := <-ch:
			fmt.Println("读到:", num)
		case <-quit:
			return
		}
		fmt.Println("....")
	}

}

func fibonacci(ch <-chan int, quit <-chan bool) {
	for {
		select {
		case num := <-ch:
			fmt.Println(num)
		case <-quit:
			fmt.Println("退出操作")
			return
			// runtime.Goexit()
		}
	}
}
func main72() {
	ch := make(chan int)
	quit := make(chan bool)
	go fibonacci(ch, quit) // 打印斐波拉契
	x, y := 1, 1
	for i := 0; i < 20; i++ {
		x, y = y, x+y
		ch <- x
	}
	quit <- true
}

/* 超时处理 */
func main73() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for {
			select {

			case num := <-ch:
				fmt.Println(num)
			case <-time.After(3 * time.Second):
				fmt.Println("到了3秒")
				quit <- true
				goto lable
			}
		}
	lable:
		fmt.Println("break to lable")
	}()
	for i := 0; i < 2; i++ {
		ch <- i
		time.Sleep(time.Second * 2)
	}
	<-quit // 默认主go程,阻塞等等子go程通知,然后退出
	fmt.Println("finish!")
}

/* 死锁,是一种错误使用所导致的现象
1. 单go程自己死锁
2. go程间channel访问顺序导致死锁
3. 多go程,多channel交叉死锁
4. 尽量不要讲互斥锁,读写锁与channel混用,--隐形死锁
*/
// 死锁1
func main74() {
	ch := make(chan int)
	ch <- 72
	num := <-ch
	fmt.Println(num)
}

// 死锁2
func main75() {
	ch := make(chan int)
	// 读数据,死锁,还没写入数据
	num := <-ch
	fmt.Println(num)
	go func() {
		ch <- 2
	}()
}

// 死锁3
func main76() {
	ch1 := make(chan int)
	ch2 := make(chan int, 0)
	go func() {
		for {
			select {
			case num := <-ch1:
				ch2 <- num
			}
		}
	}()
	for {
		select {
		case num := <-ch2:
			ch1 <- num
		}
	}
}
