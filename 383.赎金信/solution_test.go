package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func canConstruct(ransomNote string, magazine string) bool {
	table := make(map[rune]int)
	for _, char := range magazine {
		table[char]++
	}
	for _, char := range ransomNote {
		if _, exists := table[char]; exists {
			table[char]--
			if table[char] < 0 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRansomNote(t *testing.T) {

}
