package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func twoEditWords(queries []string, dictionary []string) (res []string) {
	for _, query := range queries {
		for _, dic := range dictionary {
			errTimes := 0
			for i := range len(query) {
				if query[i] != dic[i] {
					errTimes++
				}
				if errTimes > 2 {
					break
				}
			}
			if errTimes <= 2 {
				res = append(res, query)
				break
			}
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestWordsWithinTwoEditsOfDictionary(t *testing.T) {

}
