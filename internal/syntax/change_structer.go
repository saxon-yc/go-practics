// 数据类型相互转换
package syntax

import "fmt"

func Changesyntax() {
	int2string()
}

func int2string() {

	// 使用 fmt.Sprintf()
	num := 20
	result := fmt.Sprintf("%v", num)
	fmt.Printf("result: %T, %v, %v\n", result, result, result == "20")

}
