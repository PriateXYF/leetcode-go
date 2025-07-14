package leetcode

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func letterCombinations(digits string) []string {
	table := map[uint8][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	var dfs func(string) []string
	dfs = func(s string) []string {
		if len(s) == 0 {
			return []string{}
		} else if len(s) == 1 {
			return table[s[len(s)-1]]
		}
		before := dfs(s[1:len(s)])
		var res []string
		for _, letter := range table[s[0]] {
			for _, b := range before {
				res = append(res, letter+b)
			}
		}
		return res
	}
	return dfs(digits)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLetterCombinationsOfAPhoneNumber(t *testing.T) {
	res := letterCombinations("232")
	fmt.Println(res)

}
