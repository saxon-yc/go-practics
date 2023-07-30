package syntax

import "fmt"

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

// 当接受者是值类型需要拷贝副本，是指针类型时不需要拷贝副本。
func MyMethod() {
	ReceiveType1()
	ReceiveType2()
}
