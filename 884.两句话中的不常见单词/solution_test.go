package leetcode

import (
	"strings"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func uncommonFromSentences(s1 string, s2 string) (res []string) {
	l1, l2 := strings.Split(s1, " "), strings.Split(s2, " ")
	table := make(map[string]uint8)
	for _, w1 := range l1 {
		table[w1]++
	}
	for _, w2 := range l2 {
		table[w2]++
	}
	for word, times := range table {
		if times == 1 {
			res = append(res, word)
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestUncommonWordsFromTwoSentences(t *testing.T) {

}
