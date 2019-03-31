package main

import (
	"fmt"
	"time"
)

// 写入数据 => 缓冲区(公共区) => 读取数据
// 1. 解耦(降低生产者和消费者之间的耦合度)
// 2. 并发(生产者和消费者数量不对等时,依然能保持正常通讯)
// 3. 缓存(生产者和消费者 数据处理速度不一致时,暂存数据)

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
		// 生产的过程是有消费,所有可能会超过5个
		fmt.Println("生产者:", i)
	}
	close(out)
}
func consumer(in <-chan int) {
	for num := range in {
		// 打印是有io操作的,有延时
		fmt.Println("消费者拿到数据:", num)
		time.Sleep(time.Second)
	}
}
func main51() {
	ch := make(chan int, 5)
	go producer(ch) // 子go程 生产者
	consumer(ch)    // 主go程 消费者
}
