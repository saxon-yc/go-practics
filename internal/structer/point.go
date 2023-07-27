package structer

import "fmt"

func Pointer() {
	arr := [5]int{0, 1, 2, 3, 4}
	fmt.Printf("Source1 arr: %v\n", arr) // Source1 arr: [0 1 2 3 4]
	p1(&arr)
	p2(arr)
	fmt.Printf("Source2 arr: %v\n", arr) // Source2 arr: [0 1 2 3 40] 受指针类型传递的影响
}

func p1(arr *[5]int) {
	arr[4] = 40
	fmt.Printf("After arr: %v\n", *arr) // After arr: [0 1 2 3 40]
}

func p2(arr [5]int) {
	arr[1] = 666
	fmt.Printf("After p2 arr: %v\n", arr) // After p2 arr: [0 666 2 3 40]
}
