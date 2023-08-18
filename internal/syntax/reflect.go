package syntax

import (
	"fmt"
	"reflect"
)

func MyReflect() {

	mp := map[interface{}]interface{}{
		"name":  "zs",
		"age":   24,
		2:       "order",
		"a":     'a',
		"array": []string{"x", "y"},
	}
	m := reflect.ValueOf(mp)
	// v, _ := json.Marshal(mp)
	fmt.Println(m.FieldByName("name"))
}
