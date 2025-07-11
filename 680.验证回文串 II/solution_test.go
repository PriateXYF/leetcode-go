package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func isPalindromeStr(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

// 验证允许删除一个字符的字符串是否是回文串
func validPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return isPalindromeStr(s[i:j]) || isPalindromeStr(s[i+1:j+1])
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidPalindromeIi(t *testing.T) {

}
