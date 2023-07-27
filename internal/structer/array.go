// 数组

package structer

import (
	"fmt"
	"sort"
)

func init() {
	fmt.Printf("init2...\n")
}

func init() {
	fmt.Printf("init1...\n")
}

func Array() {
	ramPoint()

	slice1 := []int{0, 1, 12, 55, 255, 0, 255, 5, 89, 78}
	fmt.Printf("源切片：%v\n", slice1)
	unique(slice1)
	sorted(slice1, "asc")

	fmt.Printf("isIncludes(slice1, 255): %v\n", isIncludes(slice1, 255)) // isIncludes(slice1, 255): true
}

// 内存空间分配与指针指向
func ramPoint() {
	// 可以看出数组赋值会开辟新的内存空间，而不是传递指针，arr2 和 arr3指向不同的内存空间, 这和JS相反。
	arr2 := [3]int{0, 1, 2}
	arr3 := arr2
	arr3[0] = 10
	fmt.Printf("arr2:%v, arr3:%v\n", arr2, arr3) // arr2:[0 1 2], arr3:[10 1 2]
	// 从打印结果可以看出：切片赋值时，赋值的是指针，sli2 和 sli3指向同一内存空间
	sli2 := []int{0, 1, 2}
	sli3 := sli2
	sli3[0] = 10
	fmt.Printf("sli2:%v, sli3:%v\n", sli2, sli3) // sli2:[10 1 2], sli3:[10 1 2]
}

// 切片排序
func sorted(sli []int, sortStatus string) []int {
	sort.Slice(sli, func(i, j int) bool {
		if sortStatus == "asc" { // 生序
			return sli[i] < sli[j]
		} else { // 降序
			return sli[i] > sli[j]
		}
	})
	fmt.Printf("排序后的切片：%v\n", sli)
	return sli
}

// 去重
func unique(sli []int) []int {
	var result []int
	// 使用map来判断某个key对应的value是否存在，
	// 且空的struct不占内存空间，空间换时间
	// temp := map[int]struct{}{}

	temp := make(map[int]struct{})
	for _, v := range sli {
		if _, ok := temp[v]; !ok {
			temp[v] = struct{}{}
			result = append(result, v)
		}
	}
	fmt.Printf("去重后的切片：%v\n", result)

	return result
}

// 判断数组中是否存在某个元素
// 因为 go 中没有 includes、find、indexOf 等方法，
// 但可以通过sort包的排序和二分法，去判断（这样比loop更加高效）
func isIncludes(arr []int, target int) bool {
	sort.Ints(arr)
	index := sort.SearchInts(arr, target)
	if index < len(arr) && arr[index] == target {
		return true
	}
	return false
}
