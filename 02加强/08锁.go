package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 互斥锁
// 新建的互斥锁状态为0
var mutex sync.Mutex

func printer1(str string) {
	mutex.Lock() // 访问共享数据之前,加锁
	for _, ch := range str {
		fmt.Printf("%c", ch)
		time.Sleep(time.Second * 3)
	}
	mutex.Unlock() // 共享数据访问结束,解锁
}

func person3() {
	printer1("HELLO")
}
func person4() {
	printer1("wolrd")
}
func main81() {
	person3()
	person4()
}

/* 读写锁 */
var rwMutex sync.RWMutex

/* chennel和读写锁的混用,导致了死锁 */
func readGo(in <-chan int, idx int) {
	for {
		rwMutex.RLock()
		num := <-in
		fmt.Printf("--第%d个读go程,读出%d\n", idx, num)
		rwMutex.RUnlock()
	}
}
func writeGo(out chan<- int, idx int) {
	for {
		// 生成随机数
		num := rand.Intn(1000)
		rwMutex.Lock()
		out <- num
		fmt.Printf("第%d个写go程,写入%d\n", idx, num)
		time.Sleep(time.Microsecond) // 放大实验现象
		rwMutex.Unlock()
	}
}
func main82() {
	// 随机数种子
	rand.Seed(time.Now().UnixNano())
	// quit := make(chan bool) //用于关闭主go程的channel
	ch := make(chan int) // 用于数据的传递的channel
	for i := 0; i < 5; i++ {
		go readGo(ch, i+1)
	}
	for i := 0; i < 5; i++ {
		go writeGo(ch, i+1)
	}
	// <-quit
	for {

	}
}

var value int

// 读写锁正确写法
func readGo1(idx int) {
	for {
		rwMutex.RLock()
		num := value
		fmt.Printf("--第%d个读go程,读出%d\n", idx, num)
		rwMutex.RUnlock()
	}
}
func writeGo1(idx int) {
	for {
		// 生成随机数
		num := rand.Intn(1000)
		rwMutex.Lock()
		value = num
		fmt.Printf("第%d个写go程,写入%d\n", idx, num)
		rwMutex.Unlock()
		time.Sleep(time.Microsecond) // 放大实验现象
	}
}
func main83() {
	// 随机数种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		go readGo1(i + 1)
	}
	for i := 0; i < 5; i++ {
		go writeGo1(i + 1)
	}
	for {

	}
}

// channel完成互斥锁
func readGo2(in <-chan int, idx int) {
	for {
		num := <-in
		fmt.Printf("--第%d个读go程,读出%d\n", idx, num)
	}
}
func writeGo2(out chan<- int, idx int) {
	for {
		// 生成随机数
		num := rand.Intn(1000)
		out <- num
		fmt.Printf("第%d个写go程,写入%d\n", idx, num)
		time.Sleep(time.Microsecond) // 放大实验现象
	}
}

// channel读取的数据是不规则的,而读写锁的数据是规律的
func main84() {
	// 随机种子
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go readGo2(ch, i+1)
	}
	for i := 0; i < 5; i++ {
		go writeGo2(ch, i+1)
	}
	for {

	}
}

// 条件变量不是锁,经常和锁一起使用,流程:
/* 1. 创建条件变量,var cond sync.Cond
2. 指定条件变量用的锁: cond.L = new(sync.Mutex)
3. 给公共区加锁(互斥量) cond.L.Lock()
4. 判断是否到达阻塞条件(缓冲区满/空) -- for 循环判断 for len(ch)==cap(ch) { cond.Wait() } // 阻塞等待 作用: 1.阻塞 2. 解锁 3.加锁
5. 访问公共区--读,写数据,打印
6. 解锁条件变量用的锁(粒度越小越好)  cond.Unlock()
7. 唤醒阻塞在条件变量上的对端 cond.Signal( )
*/
var cond sync.Cond // 定义全局条件变量
func producer2(out chan<- int, idx int) {
	for {
		// 先加锁
		cond.L.Lock()
		// 判断缓冲区是否满
		for len(out) == 5 {
			cond.Wait()
		}
		num := rand.Intn(800)
		out <- num
		fmt.Printf("生产者%dth,生产:%d\n", idx, num)
		// 访问公共区结束,并且打印结束,解锁
		cond.L.Unlock()
		// 唤醒阻塞在条件变量上的消费者
		cond.Signal()
		time.Sleep(time.Second)
	}
}

func consumer2(in <-chan int, idx int) {
	for {
		// 先加锁
		cond.L.Lock()
		// 判断缓冲区是否为空
		for len(in) == 0 {
			cond.Wait()
		}
		num := <-in
		fmt.Printf("----消费者%dth,消费:%d\n", idx, num)
		// 解锁
		cond.L.Unlock()
		// 唤醒
		cond.Signal()
		time.Sleep(time.Second)
	}

}
func main85() {
	producer := make(chan int, 5)
	rand.Seed(time.Now().UnixNano())
	quit := make(chan bool)
	// 指定条件变量 使用的锁
	cond.L = new(sync.Mutex) // 互斥锁
	for i := 0; i < 5; i++ {
		go producer2(producer, i+1)
	}
	for i := 0; i < 5; i++ {
		go consumer2(producer, i+1)
	}
	<-quit
}
