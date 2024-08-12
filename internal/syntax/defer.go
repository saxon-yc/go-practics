package syntax

import "fmt"

func MyDefer() {
	b()

	x, y := foo(1, 2)
	fmt.Println("x=", x, "y=", y)

	testDeferOrder()
}
func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func foo(a, b int) (x, y int) {

	defer func(s string) {
		fmt.Printf("s: %v\n", s)
		x = x * 5
		y = y * 10
	}("a")

	x = a + 5
	y = b + 6
	return
}

// chapter4/sources/deferred_func_7.go
func foo1() {
	for i := 0; i <= 3; i++ {
		defer fmt.Println(i)
	}
}

func foo2() {
	for i := 0; i <= 3; i++ {
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}
}

func foo3() {
	for i := 0; i <= 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}

func testDeferOrder() {
	fmt.Println("foo1 result:")
	foo1()
	fmt.Println("\nfoo2 result:")
	foo2()
	fmt.Println("\nfoo3 result:")
	foo3()
	testDeferParams()
}

// chapter4/sources/deferred_func_8.go

func fnoo1() {
	// 重新分配新切片
	/* sl := []int{1, 2, 3}
	defer func(a []int) {
		fmt.Println(a) // [1 2 3]
	}(sl)

	sl = []int{3, 2, 1} //
	_ = sl */

	// 修改底层数组元素
	sl := []int{1, 2, 3}
	defer func(a []int) {
		fmt.Println(a) // [30 2 3]
	}(sl)
	sl[0] = 30
	_ = sl
}

func fnoo2() {
	sl := []int{1, 2, 3}
	defer func(p *[]int) {
		fmt.Println(*p)
		fmt.Printf("&p==&sl:%v\n", p == &sl) // true
	}(&sl)

	sl = []int{3, 2, 1}
	_ = sl
}

func testDeferParams() {
	fnoo1()
	fnoo2()
}
