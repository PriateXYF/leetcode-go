package leetcode

import (
	"sort"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	gp, sp := 0, 0
	for gp < len(g) && sp < len(s) {
		if s[sp] >= g[gp] {
			gp++
		}
		sp++
	}
	return gp
}

//leetcode submit region end(Prohibit modification and deletion)

func TestAssignCookies(t *testing.T) {

}
