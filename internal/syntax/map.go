// map 结构
package syntax

import (
	"fmt"
	"sync"
	"time"
)

func MyMap() {
	// mpf1()
	// testMapIsUnsafe()

}

func mpf1() {
	// mp1 := map[string]string{"name": "zs", "age": "26", "sex": "man"}
	mp1 := make(map[string]string)
	mp1["name"] = "zs"
	mp1["age"] = "26"
	mp1["sex"] = "man"
	// 判断某个键值是否存在
	t, ok := mp1["name"]
	fmt.Printf("t: %v, ok: %v\n", t, ok) // t: zs, ok: true

	for k := range mp1 {
		if k == "sex" {
			delete(mp1, k)
		}
	}
	// delete(mp1, "age")
	fmt.Println(mp1)

}

// chapter3/sources/map_concurrent_read_and_write.go
// func doIteration(m map[int]int) {
// 	for k, v := range m {
// 		_ = fmt.Sprintf("[%d, %d] ", k, v)
// 	}
// }
// func doWrite(m map[int]int) {
// 	for k, v := range m {
// 		m[k] = v + 1
// 	}
// }

// func testMapIsUnsafe() {
// 	m := map[int]int{
// 		1: 11,
// 		2: 12,
// 		3: 13,
// 	}

// 	go func() {
// 		for i := 0; i < 1000; i++ {
// 			doIteration(m)
// 		}
// 	}()

// 	go func() {
// 		for i := 0; i < 1000; i++ {
// 			doWrite(m)
// 		}
// 	}()

// 	time.Sleep(5 * time.Second)

// }

// /*
// fatal error: concurrent map iteration and map write

// goroutine 18 [running]:
// */
func doIteration(m *sync.Map) {
	m.Range(func(key, value interface{}) bool {
		_ = fmt.Sprintf("[%v, %v] ", key, value)
		return true
	})
}

func doWrite(m *sync.Map) {
	m.Range(func(key, value interface{}) bool {
		m.Store(key, value.(int)+1)
		return true
	})
}

func testMapIsUnsafe() {
	m := &sync.Map{}
	m.Store(1, 11)
	m.Store(2, 12)
	m.Store(3, 13)

	go func() {
		for i := 0; i < 1000; i++ {
			doIteration(m)
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			doWrite(m)
		}
	}()

	fmt.Printf("safe map=%v\n", m)
	time.Sleep(5 * time.Second)

}
