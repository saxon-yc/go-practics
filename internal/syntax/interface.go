// 接口 interface
package syntax

import (
	"fmt"
	"math"
)

func MyInterface() {
	useStructInterface()
	useBaseInterface()

}

type Operater interface {
	reade(brand string) bool
	write(size int) bool
}

type Computer struct {
	brand string
	size  int
}

func (com Computer) read(brand string) bool {
	return brand == com.brand
}
func (com Computer) write(size int) bool {
	return size == com.size
}

// 接口和结构体组合的形态
func useStructInterface() {
	com := Computer{
		brand: "Apple",
		size:  14,
	}
	fmt.Printf("com.read(\"Apple\"): %v\n", com.read("Apple"))
	fmt.Printf("com.read(): %v\n", com.write(14))
}

type Abser interface {
	Abs() float64
}

// 自定义类型
type MyFloat float64

// 普通接口
func useBaseInterface() {
	var a Abser
	sq := -math.Sqrt2
	f := MyFloat(sq)
	fmt.Printf("sq:%v,f:%v\n", sq, f)
	a = f
	a.Abs()
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		// float64() 强制类型转换
		return float64(-f)
	}
	return float64(f)
}
