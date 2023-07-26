// 基本数据类型

package structer

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

func Base() {
	fmt.Printf("****基本数据类型****\n")

	fmt.Printf("---- 布尔类型 ----\n")
	boolType()

	fmt.Printf("---- 数字类型 ----\n")
	numType()

	fmt.Printf("---- 字符串类型 ----\n")
	strType()

	fmt.Printf("****基本数据类型****\n")
}

func boolType() {
	var boolean bool = true
	fmt.Printf("boolean's type is:%T, value: %v\n", boolean, boolean) // boolean's type is:bool, value: true
}

func numType() {
	// 整型, 整型计算速度最快
	var num int = 10
	fmt.Printf("num's type is:%T, value: %v\n", num, num) // num's type is:int, value: 10
	var numi int64
	fmt.Printf("整型的零值为：%v\n", numi) // 整型的零值为：0

	// 其他类型
	// byte 是 unit8 的别名，范围为：0～255
	// var num1 uint8 = 0
	// var num2 byte = 255
	var num2Arr []byte = []byte{0, 'A', 'Z', 'a', 'z', 255} // byte's array: [0 65 90 97 122 255]
	fmt.Printf("byte's array: %v\n", num2Arr)

	// var num3 uint16 = 16

	// var num4 uint32 = 32
	// var num44 rune = 32

	// var num4 uint64 = 64
	var num5 uint64 = uint64(math.Pow(2, 64))
	fmt.Printf("2 的64次方: %v\n", num5) // 2 的64次方: 18446744073709551615

	// 浮点型
	// var num6 float32 = 32.32
	var numf float64
	fmt.Printf("浮点型的零值为：%v\n", numf) // 浮点型的零值为：0
	var num7 float64 = math.Pow(40, 2)
	fmt.Printf("40 的平方：%v\n", num7) // 40 的平方：1600

	var judge bool = 0.1+0.2 == 0.3
	fmt.Printf("0.1+0.2=%v is: %v\n", 0.1+0.2, judge) // 0.1+0.2=0.3 is: true

}

func strType() {
	var str string = "string"
	fmt.Printf("str's type is:%T, value: %v\n", str, str) // str's type is:string, value: string

	// 字符串拼接
	str1 := "hello"
	str2 := "world"

	// "+" 拼接
	fmt.Printf("'+' 拼接: %v\n", str1+" "+str2) // '+' 拼接: hello world

	// strings.Join() 拼接
	fmt.Printf("strings.Join() 拼接: %v \n", strings.Join([]string{str1, str2}, " ")) // strings.Join() 拼接: hello world

	// fmt.Sprintf() 拼接
	result := fmt.Sprintf("%s %s", str1, str2) // fmt.Sprintf() 拼接: hello world
	fmt.Printf("fmt.Sprintf() 拼接: %v\n", result)

	// bytes.Buffer.WriteString() 拼接, 缓冲区写入速度最快
	var buffer bytes.Buffer
	buffer.WriteString(str1)
	buffer.WriteString(" ")
	buffer.WriteString(str2)
	fmt.Printf("bytes.Buffer.WriteString() 拼接: %v\n", buffer.String()) // bytes.Buffer.WriteString() 拼接: hello world

	// Refer: https://cloud.tencent.com/developer/article/1615783
	fmt.Printf(`
		反引号
	`)

}
