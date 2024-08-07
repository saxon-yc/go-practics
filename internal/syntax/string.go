package syntax

import (
	"fmt"
	"testing"
	"unsafe"
)

// chapter3/sources/string_immutable2.go
func MyString() {
	// tChangeString()
	// compareString()
	// string_slice_to_string()
	// string_mallocs_in_convert()
	string_for_range_covert_optimize()
}
func compareString() {
	// 底层指向的值是相同的
	s1 := "世界和平"
	s2 := "世界" + "和平"
	fmt.Println(s1 == s2)   // true
	fmt.Println(&s1 == &s2) // false

}

func tChangeString() {
	// 原始string
	var s string = "hello"
	fmt.Println("original string:", s)

	// 试图通过unsafe指针改变原始string
	modifyString(&s)
	fmt.Println(s)
}
func modifyString(s *string) {
	// 取出第一个8字节的值
	p := (*uintptr)(unsafe.Pointer(s))

	// 获取底层数组的地址
	var array *[5]byte = (*[5]byte)(unsafe.Pointer(*p))

	var len *int = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) + unsafe.Sizeof((*uintptr)(nil))))

	for i := 0; i < (*len); i++ {
		fmt.Printf("%p => %c\n", &((*array)[i]), (*array)[i])
		p1 := &((*array)[i])
		v := (*p1)
		(*p1) = v + 1 //try to change the character
	}
}

/*
	$go run string_immutable2.go
	original string: hello
	0x10d1b9d => h
	unexpected fault address 0x10d1b9d
	fatal error: fault
	[signal SIGBUS: bus error code=0x2 addr=0x10d1b9d pc=0x109b079]
*/

// chapter3/sources/string_slice_to_string.go
/* func string_slice_to_string() {
	rs := []rune{
		0x4E2D,
		0x56FD,
		0x6B22,
		0x8FCE,
		0x60A8,
	}

	s := string(rs)
	fmt.Println(s)

	sl := []byte{
		0xE4, 0xB8, 0xAD,
		0xE5, 0x9B, 0xBD,
		0xE6, 0xAC, 0xA2,
		0xE8, 0xBF, 0x8E,
		0xE6, 0x82, 0xA8,
	}

	s = string(sl)
	fmt.Println(s)
}
*/

// chapter3/sources/string_mallocs_in_convert.go
/* func byteSliceToString() {
	sl := []byte{
		0xE4, 0xB8, 0xAD,
		0xE5, 0x9B, 0xBD,
		0xE6, 0xAC, 0xA2,
		0xE8, 0xBF, 0x8E,
		0xE6, 0x82, 0xA8,
		0xEF, 0xBC, 0x8C,
		0xE5, 0x8C, 0x97,
		0xE4, 0xBA, 0xAC,
		0xE6, 0xAC, 0xA2,
		0xE8, 0xBF, 0x8E,
		0xE6, 0x82, 0xA8,
	}

	_ = string(sl)
}

func stringToByteSlice() {
	s := "中国欢迎您，北京欢迎您"
	_ = []byte(s)
}

func string_mallocs_in_convert() {
	fmt.Println(testing.AllocsPerRun(1, byteSliceToString))
	fmt.Println(testing.AllocsPerRun(1, stringToByteSlice))
}
*/

// chapter3/sources/string_for_range_covert_optimize.go

func convert() {
	s := "中国欢迎您，北京欢迎您"
	sl := []byte(s)
	for _, v := range sl {
		_ = v
	}
}
func convertWithOptimize() {
	s := "中国欢迎您，北京欢迎您"
	for _, v := range []byte(s) {
		_ = v
	}
	fmt.Printf("string(v): %v\n", string([]byte(s)))
}

func string_for_range_covert_optimize() {
	fmt.Println(testing.AllocsPerRun(1, convert))
	fmt.Println(testing.AllocsPerRun(1, convertWithOptimize))
}
