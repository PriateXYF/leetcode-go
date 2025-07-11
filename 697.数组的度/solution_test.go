package leetcode

import (
	"math"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func findShortestSubArray(nums []int) int {
	// 找到度
	degrees := map[int][]int{}
	for i, num := range nums {
		if info, exists := degrees[num]; exists {
			degrees[num] = []int{info[0] + 1, info[1], i}
		} else {
			degrees[num] = []int{1, i, -1}
		}
	}
	maxDegree, length := 0, math.MaxInt
	for _, info := range degrees {
		if info[0] > maxDegree {
			maxDegree = info[0]
			length = info[2] - info[1] + 1
		} else if info[0] == maxDegree {
			length = min(length, info[2]-info[1]+1)
		}
	}
	if length == math.MaxInt || length <= 0 {
		return 1
	}
	return length
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDegreeOfAnArray(t *testing.T) {

}
