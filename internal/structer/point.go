package structer

import "fmt"

type Stu struct {
	name string
	age  int
}

func Pointer() {
	arr := [5]int{0, 1, 2, 3, 4}
	fmt.Printf("Source1 arr: %v\n", arr) // Source1 arr: [0 1 2 3 4]
	p1(arr)                              // // After p1 arr: [0 666 2 3 4]
	p2(&arr)                             // // After p2 arr: [0 1 2 3 40]
	fmt.Printf("Source2 arr: %v\n", arr) // Source2 arr: [0 1 2 3 40] 受指针类型传递的影响
	p3(&Stu{"zhangsan", 26})             // {zhangsan 26}
}

func p1(arr [5]int) {
	arr[1] = 666
	fmt.Printf("After p1 arr: %v\n", arr)
}
func p2(arr *[5]int) {
	arr[4] = 40
	fmt.Printf("After p2 arr: %v\n", *arr)
}

func p3(s *Stu) {
	fmt.Printf("%v\n", *s)
}
