package leetcode

import (
	"strings"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
	words := strings.Split(s, " ")
	// 反转单词数组
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	res := []rune(strings.Join(words, " "))
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseWordsInAStringIii(t *testing.T) {

}
