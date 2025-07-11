package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) (res int) {
	table := make([]int, 26*2)
	for _, char := range s {
		idx := char - 'A'
		if char >= 'a' {
			idx = char - 'a' + 26
		}
		table[idx]++
		if table[idx]%2 == 0 {
			res += 2
		}
	}
	if len(s)-res > 0 {
		res++
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestPalindrome(t *testing.T) {

}
