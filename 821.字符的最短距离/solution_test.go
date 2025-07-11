package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func shortestToChar(s string, c byte) []int {
	// 先遍历一遍查找对应位置
	var matchIndex []int
	for i, char := range s {
		if byte(char) == c {
			matchIndex = append(matchIndex, i)
		}
	}
	// 再次遍历原始字符串计算距离
	var res []int
	now := 0
	for i, _ := range s {
		if i <= matchIndex[0] {
			res = append(res, matchIndex[0]-i)
		} else if i >= matchIndex[len(matchIndex)-1] {
			res = append(res, i-matchIndex[len(matchIndex)-1])
		} else {
			if i > matchIndex[now+1] {
				now++
			}
			dist := min(i-matchIndex[now], matchIndex[now+1]-i)
			res = append(res, dist)
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestShortestDistanceToACharacter(t *testing.T) {
	//shortestToChar("loveleetcode", 'e')
}
