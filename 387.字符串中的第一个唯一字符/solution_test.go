package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func firstUniqChar(s string) int {
	table := make(map[rune]int)
	for _, char := range s {
		table[char]++
	}
	for i, char := range s {
		if times, exists := table[char]; exists && times == 1 {
			return i
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFirstUniqueCharacterInAString(t *testing.T) {

}
