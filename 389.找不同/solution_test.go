package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func findTheDifference(s string, t string) byte {
	table := make([]rune, 26)
	for _, char := range s {
		table[char-'a']++
	}
	for _, char := range t {
		table[char-'a']--
		if table[char-'a'] < 0 {
			return byte(char)
		}
	}
	return byte('a')
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFindTheDifference(t *testing.T) {

}
