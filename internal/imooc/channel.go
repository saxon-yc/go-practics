package imooc

import (
	"fmt"
	"sync"
	"time"
)

/*
不用通过共享内存来通信，而是通过通信来实现内存共享
java、python多线程编程时，两个线程之间的通信最常用的是通过一个全局变量
当然go也会使用一个消息队列的机制，消费者和生产者之间的关系
*/

// channel 的基础用法
func cha1() {
	// channel 是一个环形数组
	// msg := make(chan string, 1)

	// 缓冲与否场景：
	// 无缓冲channel： 适用于通知，B要第一时间知道A是否已经完成
	// 有缓冲channel： 适用于消费者与生产者
	/*
		go 中channel的应用场景：
			1. 消息传递、消息过滤
			2. 信号广播
			3. 并发控制
			4. 事件订阅和广播
			5. 任务分发
			6. 结果汇总
			7.同步和异步
			......
	*/

	// channel的初始化值如果为0(无缓存)，放值进去会被阻塞。（需要启动goroutine能处理）
	msg := make(chan string, 0)

	go func(msg chan string) { // happen-before 机制可以保证读写不出错
		data := <-msg
		fmt.Printf("data: %v\n", data)
	}(msg)
	msg <- "hello"
	time.Sleep(time.Second)
}

// 容易发生死锁的场景：waitgroup少了done、无缓存channel也容易出现

// chan+forrage
func chan2() {
	msg := make(chan int, 2)

	go func(msg chan int) {
		for v := range msg {
			fmt.Printf("data: %d\n", v)
		}
		fmt.Printf("all done \n")
	}(msg)
	msg <- 1
	msg <- 2

	close(msg) //关闭一个队列。 与其他编程语言有很大的区别

	// 已经关闭的channel，可以继续去值，但不能再放值了。
	// msg <- 3 // panic: send on closed channel。

	// d := <-msg
	// fmt.Printf("d: %v\n", d)
	/*
		d: 1
		data: 2
		all done
	*/

	time.Sleep(time.Second)
}

// 单向channel
var wg sync.WaitGroup

func producer(out chan<- int) {
	defer wg.Done()
	for i := 0; i <= 9; i++ {
		out <- i * i
	}
	close(out)
}
func consumer(in <-chan int) {
	defer wg.Done()
	for v := range in {
		fmt.Printf("read value: %d\r\n", v)
	}
}
func chan3() {
	// 默认情况下chan是双向的，但是我们经常讲chan作为一个阐述传递，希望对方是单向使用
	// 单向chan不能转换为双向chan，反之可以
	// var ch1 chan string    // 双向chan
	// var ch2 chan<- float32 // 单向chan，只允许写入float32的数据
	// var ch3 <-chan int     // 单向chan，只能读取int的数据
	c := make(chan int)
	wg.Add(2)
	go producer(c)
	go consumer(c)
	wg.Wait()

}
func MyChannel() {
	// cha1()
	// chan2()
	chan3()
}
