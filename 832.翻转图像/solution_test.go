package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func flipAndInvertImage(image [][]int) [][]int {
	for _, pixel := range image {
		for head, tail := 0, len(pixel)-1; head <= tail; head, tail = head+1, tail-1 {
			pixel[head], pixel[tail] = pixel[tail]^1, pixel[head]^1
		}
	}
	return image
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFlippingAnImage(t *testing.T) {
	flipAndInvertImage([][]int{[]int{1, 1, 0}, []int{1, 0, 1}, []int{0, 0, 0}})
}
