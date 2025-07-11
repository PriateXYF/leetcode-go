package leetcode

import (
	"regexp"
	"strings"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	pattern := regexp.MustCompile(`[^a-zA-Z0-9]`)
	str := pattern.ReplaceAllString(s, "")
	str = strings.ToLower(str)
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidPalindrome(t *testing.T) {

}
