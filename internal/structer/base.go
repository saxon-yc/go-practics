// 基本数据类型

package structer

import "fmt"

func Base() {
	fmt.Printf("****基本数据类型****")
	// 字符串
	var str string = "string"
	fmt.Printf("str's type is:%T, value: %v\n", str, str)

	// 布尔
	var boolean bool = true
	fmt.Printf("boolean's type is:%T, value: %v\n", boolean, boolean)

	// int
	var num int = 10
	fmt.Printf("num's type is:%T, value: %v\n", num, num)

	fmt.Printf("****基本数据类型****")

}
