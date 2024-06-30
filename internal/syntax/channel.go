// channel 通道
package syntax

import (
	"fmt"
	"log"
	"strings"
	"time"
	"unicode/utf8"
)

var chan1 = make(chan string)

func sendData1(msg string) {
	data := "World"
	fmt.Printf("Default data=%v \n", data)
	time.Sleep(time.Millisecond * 100)
	chan1 <- strings.Join([]string{msg, data}, " ") // 写数据
}
func testChanIO() {
	fmt.Println("----------- channel 数据 I/O -----------")
	defer close(chan1) // 函数退出出之前关闭通道
	go sendData1("Hello")
	fmt.Printf("wait... \n")

	received, ok := <-chan1 // 读数据
	if ok {
		fmt.Printf("received data=%v \n", received)
	}
	fmt.Println("end...")

	/*
		wait...
		Default data=World
		received data=Hello World
		end...
	*/
}

var chan2 = make(chan string, 6)
var chan3 = make(chan string, 4)

func forReadStr(str string) {
	defer wg.Done()

	length := len(str)
	log.Printf("len(str)字符长度=%v, utf8.RuneCountInString(str)字符串长度=%v", length, utf8.RuneCountInString(str))
	// for 按个写入数据
	for i := 0; i < length; {
		r, size := utf8.DecodeRuneInString(str[i:])
		chan2 <- string(r)
		i += size
	}
}
func forRangeReadStr(str string) {
	defer wg.Done()
	// for range 按个写入数据
	for _, v := range str { // v Tyep: Rune
		s := string(v)
		chan3 <- s
	}
}
func testSelect() {
	fmt.Println("----------- channel select -----------")
	defer close(chan2)
	defer close(chan3)

	wg.Add(2)
	go forReadStr("我爱世界和平")
	go forRangeReadStr("第一等人")
	wg.Wait()
	time.Sleep(time.Second)

	select {
	case r2, ok := <-chan2: // go 语言的特色语法，防止读不到报错
		if ok {
			fmt.Printf("chan2: %v\n", r2)
			// 读多写少会造成死锁
			// for v := range chan2 {
			// 	log.Printf("v is %v", v)
			// 	/*
			// 		2024/06/30 16:14:25 我
			// 		2024/06/30 16:14:25 爱
			// 		2024/06/30 16:14:25 世
			// 		2024/06/30 16:14:25 界
			// 		2024/06/30 16:14:25 和
			// 		2024/06/30 16:14:25 平
			// 	*/
			// }
		}
	case r3, ok := <-chan3:
		if ok {
			fmt.Printf("chan3: %v\n", r3)
		}
	}
	/*
		随机取一个执行，保证每一个都能被执行到
	*/

}
func MyChannel() {
	testChanIO()
	testSelect()
}
