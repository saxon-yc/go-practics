package syntax

import "fmt"

var chanInt = make(chan int)
var chanStr = make(chan string)

func Myselect() {
	go func() {
		chanInt <- 100
		chanStr <- "hello"
		close(chanInt)
		close(chanStr)
	}()
	for {

		select {
		case r := <-chanInt:
			fmt.Printf("chanInt: %v\n", r)
		case r := <-chanStr:
			fmt.Printf("chanStr: %v\n", r)
		default:
			fmt.Printf("default....\n")

		}
	}

}
