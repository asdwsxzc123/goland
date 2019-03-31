package main

import (
	"fmt"
	"time"
)

// time.Timer可以告诉你timer要等待多长时间,定时到达后,系统会自动向定时器的成员C写入系统当前时间(对chan的写操作)
func main61() {
	// 1.Timer.C
	fmt.Println("当前时间", time.Now())
	myTimer := time.NewTimer(time.Second * 2)
	nowTime := <-myTimer.C
	fmt.Println("现在时间", nowTime)

	// 2.time.Sleep()

	// 3.time.After()
	nowTime1 := <-time.After(time.Second * 2)
	fmt.Println("after现在时间", nowTime1)
}

/* 重置,停止 */
func main62() {
	myTimer := time.NewTimer(time.Second)
	// 重置定时时长
	myTimer.Reset(1 * time.Second)
	go func() {
		<-myTimer.C
		fmt.Println("子go程,定时完毕")
	}()
	// 停止定时器
	// myTimer.Stop()
	for {
	}
}

// 周期定时time.newTicker
func main63() {
	fmt.Println("当前时间", time.Now())
	// 每隔一秒执行一次
	myTIme := time.NewTicker(time.Second * 3)
	go func() {
		for {
			nowTime := <-myTIme.C
			fmt.Println(nowTime)

		}
	}()
	for {
	}
}
