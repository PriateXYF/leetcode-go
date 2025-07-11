package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func isIsomorphic(s string, t string) bool {
	tableA, tableB := map[uint8]uint8{}, map[uint8]uint8{}
	for i := 0; i < len(s); i++ {
		sc, tc := s[i], t[i]
		// 如果已存在正向映射且映射错误
		if vA, existsA := tableA[sc]; existsA && (vA != tc) {
			return false
		} else {
			tableA[sc] = tc
		}
		// 如果已存在逆向映射且映射错误
		if vB, existsB := tableB[tc]; existsB && (vB != sc) {
			return false
		} else {
			tableB[tc] = sc
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestIsomorphicStrings(t *testing.T) {

}
