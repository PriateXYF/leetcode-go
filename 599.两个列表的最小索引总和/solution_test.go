package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func findRestaurant(list1 []string, list2 []string) (res []string) {
	// 保证 list1 > list2
	if len(list1) < len(list2) {
		list1, list2 = list2, list1
	}
	// 对 list2 建立哈希表
	table := map[string]int{}
	for i, rest := range list2 {
		table[rest] = i
	}
	// 查找 list1
	minIdxSum := len(list1) + len(list2)
	for i1, rest := range list1 {
		if i1 > minIdxSum {
			break
		}
		i2, exists := table[rest]
		if !exists {
			continue
		}

		if i1+i2 < minIdxSum {
			res = []string{}
		}
		if i1+i2 <= minIdxSum {
			res = append(res, rest)
		}
		minIdxSum = min(i1+i2, minIdxSum)
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMinimumIndexSumOfTwoLists(t *testing.T) {

}
