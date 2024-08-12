package syntax

import (
	"fmt"
	"io"
	"reflect"
	"time"
)

// 当接受者是值类型需要拷贝副本，是指针类型时不需要拷贝副本。
func MyMethod() {
	// ReceiveType1()
	// ReceiveType2()
	// testReceiverType()
	testQuestion()
	testMethodUni()
	testMethodUni2()
}

type Animal struct {
	name string
}

func (cat Animal) eat(food string) {
	fmt.Printf("%v is eating %v\n", cat.name, food)
}

func ReceiveType1() {
	cat := Animal{
		name: "Miaomiao",
	}
	cat.eat("fesh")          // Miaomiao is eating fesh
	cat.play("playing ball") // Miaomiao is playing ball
}

func (dog *Animal) play(playing string) {
	fmt.Printf("%v is %v\n", dog.name, playing) // Wangwang is running
	fmt.Printf("dog: %v\n", dog)                // dog: &{Wangwang}
}
func ReceiveType2() {
	dog := Animal{
		name: "Wangwang",
	}
	dog.play("running")
}

// chapter4/sources/method_nature_1.go
/*
type T struct {
	a int
}

func (t T) M1() {
	t.a = 10
}

func (t *T) M2() {
	t.a = 11
}

func testReceiverType() {
	var t T // t.a = 0
	println(t.a)
	t.M1()
	println(t.a)
	t.M2()
	println(t.a)
} */

// chapter4/sources/method_nature_3.go

type field struct {
	name string
}

func (p field) print() {
	fmt.Println(p.name)
}

func testQuestion() {
	data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		go fmt.Printf("v1: %v\n", v)
		go v.print()
	}
	fmt.Printf("main groutine \n")

	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		go fmt.Printf("v2: %v\n", &v)
		go v.print()
	}

	time.Sleep(3 * time.Second)

	/* data1 := []*field{{"one"}, {"two"}, {"three"}}
	for _, v := range data1 {
		go (*field).print(v)
	}

	data2 := []field{{"four"}, {"five"}, {"six"}}
	for _, v := range data2 {
		go (*field).print(&v)
	}

	time.Sleep(3 * time.Second) */
}

// chapter4/sources/method_set_utils.go

func DumpMethodSet(i interface{}) {
	v := reflect.TypeOf(i)
	elemTyp := v.Elem()

	n := elemTyp.NumMethod()
	if n == 0 {
		fmt.Printf("%s's method set is empty!\n", elemTyp)
		return
	}

	fmt.Printf("%s's method set:\n", elemTyp)
	for j := 0; j < n; j++ {
		fmt.Println("-", elemTyp.Method(j).Name)
	}
	fmt.Printf("\n")
}

// chapter4/sources/method_set_2.go

// interface check
var _ Interface = (*T)(nil)

type Interface interface {
	M1()
	M2()
	M3()
}

type T struct{}
type T1 = T

func (t T) M1()  {}
func (t *T) M2() {}
func (t T) M3()  {}

func testMethodUni() {
	var t T
	var pt T
	DumpMethodSet(&t)
	DumpMethodSet(&pt)

	var t1 T1
	var pt1 T1
	DumpMethodSet(&t1)
	DumpMethodSet(&pt1)

	DumpMethodSet((*Interface)(nil))
}

// chapter4/sources/method_set_3.go

func testMethodUni2() {
	DumpMethodSet((*io.Writer)(nil))
	DumpMethodSet((*io.Reader)(nil))
	DumpMethodSet((*io.Closer)(nil))
	DumpMethodSet((*io.ReadWriter)(nil))
	DumpMethodSet((*io.ReadWriteCloser)(nil))
}
