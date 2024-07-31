package syntax

import "fmt"

func MySlice() {
	// testFuncParams()
	// testCap()
	// testCapIncrement()
	testNilAble()
}

// 测试数组和切片的数据类型：
// 因为 go 的设计是：数组为基本数据类型，切片为引用数据类型；且函数传参是按值传递（地址副本）
// 由此函数打印结果可得出：函数传递数组值副本时，会分配新的内存空间；而切片只会拷贝引用地址副本，即指向同一内存空间（这和JS的设计有所不同）
/* func testFuncParams() {
	arr := [5]int{1, 2, 3, 4, 5}
	sli := []int{10, 20, 30, 40, 50}
	sli2 := []int{}
	copy(sli2, sli)

	changeParams := func(innerSli []int, innerArr [5]int) {
		innerArr1 := innerArr
		innerSli1 := innerSli
		innerArr[len(innerArr)-1] = 555
		innerSli[0] = 111

		fmt.Printf("After changed {innerArr:%v, innerSli:%v} \n", innerArr, innerSli) // After changed {innerArr:[1 2 3 4 555], innerSli:[111 20 30 40 50]}
		// fmt.Printf("(innerArr == arr): %v\n", (innerArr == arr))
		// fmt.Printf("(&innerSli == &sli): %v\n", (&innerSli == &sli))

		// 由此可见：数组为基本数据类型，切片为引用数据类型
		fmt.Printf("innerArr1: %v, innerSli1: %v\n", innerArr1, innerSli1) // innerArr1: [1 2 3 4 5], innerSli1: [111 20 30 40 50]

	}
	changeParams(sli, arr)
	fmt.Printf("Source data {arr:%v, sli:%v} \n", arr, sli) // Source data {arr:[1 2 3 4 5],  sli:[111 20 30 40 50]}
	fmt.Printf("sli2: %v\n", sli2)
} */

// append 本身并不会扩容 slice，而是当切片容量不足时才会发生扩容
// slice 发生扩容后，会分配新的内存空间来存放slice
// func testCap() {
// 	csli := []int{1, 2, 3}

// 	changeCap := func(new_csli []int) {
// 		// create, 不能使用下标去添加元素
// 		new_csli = append(new_csli, 4, 5, 6)
// 		// read
// 		read := new_csli[2]
// 		// update 修改第1个元素
// 		new_csli[0] = 0
// 		// delete 删除第4个元素
// 		new_csli = append(new_csli[:3], new_csli[4:]...)
// 		fmt.Printf("new_csli: %v, read: %v \n", new_csli, read) // new_csli: [0 2 3 5 6]
// 	}
// 	fmt.Printf("csli: %v\n", csli) // csli: [1 2 3]
// 	changeCap(csli)
// }

func testCapIncrement() {
	csli := make([]int, 0)
	for i := 0; i < 1000; i++ {
		csli = append(csli, i)
		fmt.Printf("The cap is:%v \n", cap(csli))
	}
	/*
		The cap is:1		(2^0)
		The cap is:2		(2^1)
		The cap is:4		(2^2)
		The cap is:8		(2^3)
		The cap is:16		(2^4)
		The cap is:32		(2^5)
		The cap is:64		(2^6)
		The cap is:128	(2^7)
		The cap is:256	(2^8)
		The cap is:512	(2^9)
		在这之前扩容时，容量都是成倍增长

		The cap is:848    约 40%
		The cap is:1280   约 33%
		扩容容量放缓，为了防止造成开辟过多的容量造成资源浪费。
		比如当明确知道有845个学生，但若扩容还是成倍数增长，从512到 1024，那么就平白浪费 1024-854个容量
	*/
}

func testNilAble() {
	// var sli []int
	// sli[0] = 0
	// fmt.Printf("sli: %v\n", sli)
	/*
		没有分配内存，length=0，所以无法访问索引为0的地址，而导致panic（相当于数组越界）
		panic: runtime error: index out of range [0] with length 0
		goroutine 1 [running]:
	*/

	// append分配内存
	// sli = append(sli, 0)
	// fmt.Printf("sli: %v\n", sli) // sli: [0]

	// make 内置函数分配并初始化 slice、map 或 chan（仅）类型的对象。与 new 一样，第一个参数是类型，而不是值。与 new 不同，make 的返回类型与其参数的类型相同，而不是指向它的指针
	sli := make([]int, 1) // 长度和容量设置为 1
	sli[0] = 0
	fmt.Printf("sli: %v\n", sli) // sli: [0]

	var mp map[string]int
	// 读
	// fmt.Printf("mp[\"a\"]: %v\n", mp["a"]) // mp["a"]: 0
	// 写 （nil dereference in map update）
	// mp["a"] = 10
	/*
		panic: assignment to entry in nil map
		goroutine 1 [running]:
	*/

	// 已初始化
	fmt.Printf("mp == nil: %v\n", mp == nil) // mp == nil: true
	m2 := map[string]int{}
	// m2 := make(map[string]int)
	// m2["a"] = 10
	fmt.Printf("m2 == nil %v\n", m2 == nil) // m2 == nil: false
}
