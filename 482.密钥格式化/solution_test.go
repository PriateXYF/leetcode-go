package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func licenseKeyFormatting(s string, k int) string {
	var chars []rune
	var res []rune
	for _, char := range s {
		if char == '-' {
			continue
		}
		if char >= 'a' {
			char = char - 'a' + 'A'
		}
		chars = append(chars, char)
	}
	if len(chars) <= k {
		return string(chars)
	}
	skips := len(chars) % k
	for i := range skips {
		res = append(res, chars[i])
	}
	for i := skips; i < len(chars); i++ {
		if i != 0 && (i-skips)%k == 0 {
			res = append(res, '-')
		}
		res = append(res, chars[i])

	}
	return string(res)
}

//leetcode submit region end(Prohibit modification and deletion)

func TestLicenseKeyFormatting(t *testing.T) {

}
