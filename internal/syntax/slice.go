package syntax

import "fmt"

func MySlice() {
	// testFuncParams()
	testCap()
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

// append 会扩容 slice
// slice 发生扩容后，会分配新的内存空间来存放slice
func testCap() {
	csli := []int{1, 2, 3}

	changeCap := func(new_csli []int) {
		// create, 不能使用下标去添加元素
		new_csli = append(new_csli, 4, 5, 6)
		// read
		read := new_csli[2]
		// update 修改第1个元素
		new_csli[0] = 0
		// delete 删除第4个元素
		new_csli = append(new_csli[:3], new_csli[4:]...)
		fmt.Printf("new_csli: %v, read: %v \n", new_csli, read) // new_csli: [0 2 3 5 6]
	}
	fmt.Printf("csli: %v\n", csli) // csli: [1 2 3]
	changeCap(csli)
}
