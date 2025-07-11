package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) (res int) {
	table := make(map[uint8]int)
	for head, tail := -1, 0; tail < len(s); tail++ {
		if before, exist := table[s[tail]]; exist && before > head {
			head = before
		}
		res = max(res, tail-head)
		table[s[tail]] = tail
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestSubstringWithoutRepeatingCharacters(t *testing.T) {
	lengthOfLongestSubstring("abcdef")
}
