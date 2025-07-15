package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func isValid(word string) bool {
	// 判断条件一
	if len(word) < 3 {
		return false
	}
	vowels := map[int32]*struct{}{
		'a': nil, 'e': nil, 'i': nil, 'o': nil, 'u': nil,
		'A': nil, 'E': nil, 'I': nil, 'O': nil, 'U': nil,
	}
	hasConsonant, hasVowel := false, false
	for _, c := range word {
		// 判断条件二
		if !(c >= '0' && c <= '9') && !(c >= 'a' && c <= 'z') && !(c >= 'A' && c <= 'Z') {
			return false
		}
		if _, exist := vowels[c]; exist {
			// 判断条件三
			hasVowel = true
		} else if !(c >= '0' && c <= '9') {
			// 判断条件四
			hasConsonant = true
		}
	}
	return hasVowel && hasConsonant
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidWord(t *testing.T) {
	isValid("UuE6")
}
