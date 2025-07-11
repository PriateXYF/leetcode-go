package leetcode

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := range len(dp) {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	res := s[0:1]
	for length := 1; length < len(s); length++ {
		for head := 0; head < len(s)-length; head++ {
			tail := head + length
			if s[head] != s[tail] {
				continue
			} else if length <= 2 {
				dp[head][tail] = true
			} else {
				dp[head][tail] = dp[head+1][tail-1]
			}
			if dp[head][tail] && length+1 > len(res) {
				res = s[head : tail+1]
			}
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLongestPalindromicSubstring(t *testing.T) {
	fmt.Println(longestPalindrome("babab"))
	fmt.Println(longestPalindrome("cbbdccdabc"))
}
