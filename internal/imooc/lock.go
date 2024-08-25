package imooc

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// WaitGroup 等待 goroutine 集合完成。主 goroutine 调用 Add 设置等待的 goroutine 数量。
// 然后每个 goroutine 运行并在完成时调用 Done。同时，Wait 可以用来阻塞，直到所有 goroutine 都完成。
var lwg sync.WaitGroup

// 锁不能复制使用，那样就失去锁的效果了
// var lock sync.Mutex
var total int32

func add() {
	defer lwg.Done()
	for i := 0; i < 100000; i++ {
		// lock.Lock()
		// total += 1
		// lock.Unlock()
		atomic.AddInt32(&total, 1)

	}
}
func sub() {
	defer lwg.Done()
	for i := 0; i < 100000; i++ {
		// lock.Lock()
		// total -= 1
		// lock.Unlock()
		atomic.AddInt32(&total, -1)
	}
}
func lock1() {
	fmt.Println("Enter main groutine...")
	lwg.Add(2)
	go add()
	go sub()
	lwg.Wait()

	fmt.Printf("total: %v\n", total)
	fmt.Printf("atomic.LoadInt32(&total): %v\n", atomic.LoadInt32(&total))
	fmt.Println("Leaving main groutine...")

	// mpp := map[string]string{
	// 	"sync.Mutex":      "锁可以复制使用",
	// 	"sync.WaitGroup":  "WaitGroup 等待 goroutine 集合完成",
	// 	"atomic.AddInt32": "atomic.AddInt32 保证原子操作",
	// }
	// for _, v1 := range mpp {
	// 	for _, v2 := range mpp {
	// 		fmt.Printf("%v vs %v\n", v1, v2)
	// 	}
	// }
}

// 锁的本质：将并行的代码串行化，使用lock肯定会影响性能
// 即使是设计锁，也应当尽量保证并行
// 有两组协程，一组负责写，一组负责读，web系统中绝大多数场景都是读多写少
// 虽然有多个goroutine，但分析得出：读协程应该是并发，读和写应该串行的（互斥的），读和读也应该是串行

func rwLockFn() {
	var rwlock sync.RWMutex
	var num int
	var wgg sync.WaitGroup

	wgg.Add(6)
	// 写锁
	go func() {
		time.Sleep(time.Second * 3)

		defer wgg.Done()
		rwlock.Lock() // 加写锁，写锁会阻止别人，的读锁获取和写锁获取
		num = 100
		rwlock.Unlock()
		fmt.Println("write num:", num)
		time.Sleep(time.Second * 5)
	}()

	// 读锁
	for i := 0; i < 5; i++ {
		go func() {
			defer wgg.Done()
			for {
				rwlock.RLock() // 加读锁，读锁不会阻止别人的读
				time.Sleep(time.Millisecond * 500)
				fmt.Println("get read, num:", num)
				rwlock.RUnlock()
			}
		}()
	}
	wgg.Wait()

	/*
		get read, num: 0
		get read, num: 0
		write num: 100
		get read, num: 100
	*/

}

func MyLock() {
	// lock1()
	rwLockFn()
}
