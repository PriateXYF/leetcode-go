package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	table := map[uint8]int{}
	for i := 0; i < len(s); i++ {
		sc, tc := s[i], t[i]
		if _, exists := table[sc]; !exists {
			table[sc] = 0
		}
		if _, exists := table[tc]; !exists {
			table[tc] = 0
		}
		table[sc]++
		table[tc]--
	}
	for _, v := range table {
		if v != 0 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidAnagram(t *testing.T) {

}
