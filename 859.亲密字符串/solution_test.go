package leetcode

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) || len(s) < 2 {
		return false
	}
	diffCount, mark, hasRepeat := 0, -1, false
	table := map[uint8]*struct{}{}
	for i := 0; i < len(s); i++ {
		if !hasRepeat {
			if _, exist := table[s[i]]; exist {
				hasRepeat = true
			}
			table[s[i]] = nil
		}
		if s[i] != goal[i] {
			diffCount++
			if diffCount == 2 && goal[i] == s[mark] && goal[mark] == s[i] {
				continue
			} else if diffCount >= 2 {
				return false
			}
			mark = i
		}
	}
	return diffCount == 2 || (diffCount == 0 && hasRepeat)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBuddyStrings(t *testing.T) {
	fmt.Println(buddyStrings("ab", "ca"))
}
