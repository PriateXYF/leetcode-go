package leetcode

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func findLast(s string, p int) int {
	if p >= len(s) || p < 0 {
		return -1
	}
	if s[p] != '#' {
		return p
	}
	for back := 2; back > 0; back-- {
		if p--; p >= 0 && s[p] == '#' {
			back += 2
		}
	}
	return p
}
func backspaceCompare(s string, t string) bool {
	p1, p2 := len(s)-1, len(t)-1
	for ; p1 >= 0 && p2 >= 0; p1, p2 = p1-1, p2-1 {
		p1, p2 = findLast(s, p1), findLast(t, p2)
		if p1 >= 0 && p2 >= 0 && s[p1] == t[p2] {
			continue
		} else if p1 < 0 && p2 < 0 {
			return true
		}
		return false
	}
	p1, p2 = findLast(s, p1), findLast(t, p2)
	if p1 >= 0 || p2 >= 0 {
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBackspaceStringCompare(t *testing.T) {
	//s := "abcdef#g###hij"
	//p := len(s) - 1
	//for p >= 0 {
	//	p = findLast(s, p)
	//	fmt.Printf("%d %c\n", p, s[p])
	//	p--
	//}
	//fmt.Println(findLast(s, len(s)-1))
	fmt.Println(backspaceCompare("aab#a###ccc", "aa#ccc"))
}
