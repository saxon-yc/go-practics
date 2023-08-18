/*
	并发编程
	@refer: https://gfw.go101.org/article/control-flows-more.html
*/

package syntax

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func MyConcurrent() {
	// myGoroutine()
	// myCurrenencySync()
	// myDeadLock()

	// myDefer()
	// fmt.Printf("triple(5): %v\n", triple(5))
	// valMoment()
	// valGoMoment()

	pr1()
	// pr2()
	changeVal()
}

/*
	*** goroutine ***

	当一个程序的主协程退出后，此程序也就退出了，即使还有一些其它协程在运行。

	log标准库中的打印函数是经过了同步处理的（下一节将解释什么是并发同步），而fmt标准库中的打印函数却没有被同步。
	如果我们在上例中使用fmt标准库中的Println函数，则不同协程的打印可能会交织在一起。（虽然对此例来说，交织的概率很低。）
*/
func sayGretting1(gretting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(gretting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d) // 睡眠片刻：0～2.5s的随机数
	}
}
func myGoroutine() {
	rand.Seed(time.Now().UnixNano()) // // Go 1.20之前需要
	log.SetFlags(0)
	go sayGretting1("Hi~", 10)
	go sayGretting1("Hello!", 10)
	time.Sleep(2 * time.Second)
	log.Println("主协程")
	// return
}

// 以上的程序存在缺陷，当主协程退出时，那20条打印语句还未完成，需要通过以下的并发同步进行修改

/*
	*** 并发同步 ***

	不同的并发计算可能共享一些资源，其中共享内存资源最为常见。 在一个并发程序中，常常会发生下面的情形：
		在一个计算向一段内存写数据的时候，另一个计算从此内存段读数据，结果导致读出的数据的完整性得不到保证。
		在一个计算向一段内存写数据的时候，另一个计算也向此段内存写数据，结果导致被写入的数据的完整性得不到保证。

	这些情形被称为数据竞争（data race）。并发编程的一大任务就是要调度不同计算，控制它们对资源的访问时段，以使数据竞争的情况不会发生。 此任务常称为并发同步（或者数据同步）
*/

func sayGretting2(gretting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(gretting)
		// fmt.Println(gretting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d)
	}
	wg.Done() // 通知当前任务已经完成。
}
func myCurrenencySync() {
	rand.Seed(time.Now().UnixMicro())
	log.SetFlags(0) // SetFlags 设置标准记录器的输出标志。 标志位有日期、时间等。
	wg.Add(2)       // 注册两个新任务
	go sayGretting2("Hi~", 10)
	go sayGretting2("Hello!", 10)
	wg.Wait() // 阻塞在这里，直到所有任务都已完成

	// 此刻就能打印完整的20条语句了
}

// 程序只能从运行状态退出，而不能从阻塞状态退出

/* 程序死锁 */
func myDeadLock() {
	log.Println(runtime.NumCPU()) // 8
	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 2)
		wg.Wait() // 阻塞在此
	}()
	wg.Wait() // 阻塞在此

	// fatal error: all goroutines are asleep - deadlock!
}

/*
*** 延迟调用函数 ***
 */
func myDefer() {
	// defer fmt.Println("The third line")
	// defer fmt.Println("The second line")
	// fmt.Println("The first line")

	/*
		output:
			The first line
			The second line
			The third line
	*/
	df2()

}
func df2() {
	defer fmt.Println("9")
	fmt.Println("0")
	defer fmt.Println("8")
	fmt.Println("1")
	if false {
		defer fmt.Println("not reachable")
	}
	defer func() {
		defer fmt.Println("7")
		fmt.Println("3")
		defer func() {
			fmt.Println("5")
			fmt.Println("6")
		}()
		fmt.Println("4")
	}()
	fmt.Println("2")
	return
	defer fmt.Println("not reachable")

	/*
		output:
		0 1 2 3 4 5 6 7 8 9
	*/
}

// 一个延迟调用 可以修改包含此延迟调用的 最内层函数的返回值
// 换句话说：一个函数内部的 延迟调用 可以修改该函数的返回值
func triple(n int) (r int) {
	defer func() {
		r += n                             // 修改返回值
		fmt.Printf("(r,n): %v %v\n", r, n) // (r,n): 15 5
	}()
	r = n + n
	fmt.Printf("(r,n): %v %v\n", r, n) // (r,n): 10 5
	return r                           // 15
	// return n + n // <=> r = n + n; return
}

// 延迟调用估值时刻
// 一个匿名函数体内的表达式是在此函数被执行的时候才会被逐渐估值的，不管此函数是被普通调用还是延迟/协程调用。
// defer 与js中的setTimeout比较相似，但setTimeout不是一个栈的结构
func valMoment() {
	func() {
		for i := 0; i < 3; i++ {
			defer fmt.Println("a:", i) // a:2 a:1 a:0
		}
	}()
	fmt.Println()

	func() {
		for i := 0; i < 3; i++ {
			defer func() {
				fmt.Println("b:", i) // b:3 b:3 b:3
			}()
		}
	}()
	fmt.Println()

	// => 对第二个函数变形
	func() {
		for i := 0; i < 3; i++ {
			defer func(i int) {
				// 此i为形参i，非实参循环变量i。
				fmt.Println("b:", i) // b:2 b:1 b:0
			}(i)
		}
	}()
	fmt.Println()

	// 或者
	for i := 0; i < 3; i++ {
		i := i // 在下面的调用中，左i遮挡了右i。
		// <=> var i = i
		defer func() {
			// 此i为上面的左i，非循环变量i。
			fmt.Println("b:", i) // // b:2 b:1 b:0
		}()
	}
}

// 估值时刻同样适用于 goroutine
func valGoMoment() {
	var a = 123
	go func(x int) {
		time.Sleep(time.Second) // 睡眠1s再执行后续任务
		fmt.Println(x, a)       // 123 789 => x是形参，a为函数外层的变量
	}(a)

	a = 789

	time.Sleep(2 * time.Second)
}

/*
*** 程序panic(恐慌) 和 recover(恢复) ***
panic 传入的值就是 recover返回的值
*/
func pr1() {
	defer func() {
		fmt.Println("正常退出")
	}()
	fmt.Println("嗨！")
	defer func() {
		v := recover()
		fmt.Println("恐慌被恢复了：", v)
	}()

	panic("拜拜！") // 产生一个恐慌
	fmt.Println("执行不到这里")

	/*
		嗨！
		恐慌被恢复了： 拜拜！
		正常退出
	*/
}

// 下面的例子在一个新协程里面产生了一个恐慌，并且此协程在恐慌状况下退出，所以整个程序崩溃了。
func pr2() {
	fmt.Printf("Hi!")
	go func() {
		time.Sleep(time.Second)
		panic("123")
	}()

	for {
		time.Sleep(time.Second)
	}
}

func changeVal() {
	var str string = "outer"

	func(v string) {
		v = "inner"
		fmt.Printf("inner str: %v\n", v)
	}(str)

	fmt.Printf("outer str: %v\n", str)

}
