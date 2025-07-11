package leetcode

import (
	"testing"
)

//leetcode submit region begin(Prohibit modification and deletion)
import "regexp"
import "strings"

func mostCommonWord(paragraph string, banned []string) (res string) {
	// 禁用词集合
	banSet := map[string]bool{}
	for _, ban := range banned {
		banSet[ban] = true
	}
	re := regexp.MustCompile(`[a-zA-Z]+`)
	words := re.FindAllString(paragraph, -1)
	wordSet := map[string]int{}
	maxTimes := 0
	for _, word := range words {
		lowerWord := strings.ToLower(word)
		// 跳过禁用词
		if _, exist := banSet[lowerWord]; exist {
			continue
		}
		if _, exists := wordSet[lowerWord]; exists {
			wordSet[lowerWord]++
		} else {
			wordSet[lowerWord] = 1
		}
		if wordSet[lowerWord] > maxTimes {
			res = lowerWord
			maxTimes = wordSet[lowerWord]
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMostCommonWord(t *testing.T) {

}
