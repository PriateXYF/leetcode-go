package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	allowNum := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 && i+1 >= len(flowerbed) && i-1 < 0 {
			flowerbed[i] = 1
			allowNum++
		}
		if i > 0 && i < len(flowerbed)-1 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 && flowerbed[i] == 0 {
			flowerbed[i] = 1
			allowNum++
		}
		if i == 0 && i < len(flowerbed)-1 && flowerbed[i+1] == 0 && flowerbed[i] == 0 {
			flowerbed[i] = 1
			allowNum++
		}
		if i == len(flowerbed)-1 && i > 0 && flowerbed[i-1] == 0 && flowerbed[i] == 0 {
			flowerbed[i] = 1
			allowNum++
		}
		if allowNum >= n {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCanPlaceFlowers(t *testing.T) {

}
