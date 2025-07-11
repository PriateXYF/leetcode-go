package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
// // 判断字母是否为元音字母
func isVowel(c int32) bool {
	vowelTable := map[int32]*struct{}{
		'a': nil, 'e': nil, 'i': nil, 'o': nil, 'u': nil, 'A': nil, 'E': nil, 'I': nil, 'O': nil, 'U': nil,
	}
	_, exists := vowelTable[c]
	return exists
}
func toGoatLatin(sentence string) string {
	isFirst, isVowelWord, firstChar, index := true, true, ' ', 1
	var res []int32
	sen := []int32(sentence)
	sen = append(sen, ' ')
	for i, c := range sen {
		if c != ' ' && isFirst {
			isFirst = false
			isVowelWord = isVowel(c)
			firstChar = c
			if isVowelWord {
				res = append(res, c)
			}
		} else if c != ' ' {
			res = append(res, c)
		} else {
			index++
			isFirst = true
			if !isVowelWord {
				res = append(res, firstChar)
			}
			res = append(res, 'm')
			for range index {
				res = append(res, 'a')
			}
			if i != len(sen)-1 {
				res = append(res, ' ')
			}
		}
	}
	return string(res)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestGoatLatin(t *testing.T) {

}
