package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 查找字符串最长公共前缀
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	min_len := 200
	max_len := 0
	for _, str := range strs {
		min_len = min(len(str), min_len)
		max_len = max(len(str), max_len)
	}
	pre := 0
	flag := false
	for pre < min_len {
		c := strs[0][pre]
		for s := 1; s < len(strs); s++ {
			if strs[s][pre] != c {
				flag = true
				break
			}
		}
		if flag {
			return strs[0][0:pre]
		}
		pre++
	}
	return strs[0][0:pre]
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestCommonPrefix(t *testing.T) {

}
