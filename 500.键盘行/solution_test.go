package leetcode

import (
	"strings"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func isKeyboardLine(word string) (res bool) {
	word = strings.ToLower(word)
	lines := []string{"qwertyuiop", "asdfghjkl", "zxcvbnm"}
	for _, line := range lines {
		res = true
		for _, char := range word {
			if !strings.ContainsRune(line, char) {
				res = false
				break
			}
		}
		if res {
			return res
		}
	}
	return false
}

// 判断单词所有字母是否在美式键盘同一行
func findWords(words []string) (res []string) {
	for _, word := range words {
		if isKeyboardLine(word) {
			res = append(res, word)
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestKeyboardRow(t *testing.T) {

}
