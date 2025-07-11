package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func isVowel(c uint8) bool {
	table := map[uint8]*struct{}{
		'a': nil,
		'e': nil,
		'i': nil,
		'o': nil,
		'u': nil,
		'A': nil,
		'E': nil,
		'I': nil,
		'O': nil,
		'U': nil,
	}
	_, exists := table[c]
	return exists
}

// 反转字符串中的元音字母
func reverseVowels(s string) string {
	str := []uint8(s)
	head, tail := 0, len(str)-1
	for head < tail {
		if isVowel(str[head]) && isVowel(str[tail]) {
			str[head], str[tail] = str[tail], str[head]
		} else if isVowel(str[head]) && !isVowel(str[tail]) {
			head--
		} else if !isVowel(str[head]) && isVowel(str[tail]) {
			tail++
		}
		head++
		tail--
	}
	return string(str)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseVowelsOfAString(t *testing.T) {

}
