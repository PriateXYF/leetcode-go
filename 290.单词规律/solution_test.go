package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func wordPattern(pattern string, s string) bool {
	pPoint, sStart := 0, 0
	// 正向字典
	dict := map[uint8]string{}
	// 反向字典
	dictB := map[string]uint8{}
	for sEnd, char := range s {
		if char == ' ' {
			pPoint++
			sStart = sEnd + 1
			continue
		}
		// 如果 pattern 位数缺失
		if pPoint >= len(pattern) {
			return false
		}
		// 正向字典如果存在目标值
		if target, exists := dict[pattern[pPoint]]; exists {
			if sEnd == len(s)-1 || s[sEnd+1] == ' ' {
				if s[sStart:sEnd+1] != target {
					return false
				}
			} else {
				continue
			}
		} else {
			if sEnd == len(s)-1 || s[sEnd+1] == ' ' {
				dict[pattern[pPoint]] = s[sStart : sEnd+1]
			}
		}

		if sEnd == len(s)-1 || s[sEnd+1] == ' ' {
			str := s[sStart : sEnd+1]
			// 反向字典如果存在目标值
			if targetB, exists := dictB[str]; exists {
				if targetB != pattern[pPoint] {
					return false
				}
			} else {
				dictB[str] = pattern[pPoint]
			}
		}
	}
	// 单词缺失
	if pPoint != len(pattern)-1 {
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestWordPattern(t *testing.T) {

}
