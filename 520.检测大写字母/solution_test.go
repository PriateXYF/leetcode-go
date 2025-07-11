package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func detectCapitalUse(word string) bool {
	allLower, allUpper, firstUpper := true, true, true
	for i, char := range word {
		if i == 0 && char >= 'A' && char <= 'Z' {
			firstUpper = true
		} else if i == 0 && char >= 'a' && char <= 'z' {
			firstUpper = false
		} else if char >= 'A' && char <= 'Z' {
			allLower = false
		} else if char >= 'a' && char <= 'z' {
			allUpper = false

		}
	}
	// 如果首字母大写,之后全是大写或小写
	if firstUpper && (allLower || allUpper) {
		return true
	}
	// 如果全是小写
	if !firstUpper && allLower {
		return true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestDetectCapital(t *testing.T) {

}
