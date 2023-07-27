// map 结构
package structer

import (
	"fmt"

	"github.com/spf13/viper"
)

func Mapp() {
	mpf1()
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

	fmt.Printf("viper.GetInt(\"app.port\"): %v\n", viper.GetInt("app.port")) // 9999

}
