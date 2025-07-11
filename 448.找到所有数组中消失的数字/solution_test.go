package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func findDisappearedNumbers(nums []int) (res []int) {
	length := len(nums)
	for _, num := range nums {
		nums[(num-1)%length] += length
	}
	for i, num := range nums {
		if num <= length {
			res = append(res, i+1)
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindAllNumbersDisappearedInAnArray(t *testing.T) {

}
