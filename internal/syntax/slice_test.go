// 数组
package syntax

import (
	"testing"
)

// 不使用cap
const sliceSize = 10000

func BenchmarkSliceInitWithoutCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sl := make([]int, 0)
		for i := 0; i < sliceSize; i++ {
			sl = append(sl, i)
		}
	}
}

// 使用cap
func BenchmarkSliceInitWithCap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sl := make([]int, 0, sliceSize)
		for i := 0; i < sliceSize; i++ {
			sl = append(sl, i)
		}
	}
}
