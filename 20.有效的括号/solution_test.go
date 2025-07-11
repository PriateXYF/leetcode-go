package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	// 括号映射表
	table := map[rune]rune{
		'[': ']',
		'(': ')',
		'{': '}',
	}
	chs := make([]rune, len(s))
	// 栈尾
	tail := -1
	// 期待吞噬的字符
	needed := ' '
	for _, c := range s {
		// 如果接受到的是左括号字符,直接入栈
		if c == '[' || c == '(' || c == '{' {
			// 更新栈尾
			tail++
			chs[tail] = c
			// 更新期待字符
			needed = table[c]
		} else { // 如果接受的是右括号字符
			// 判断是否是期待字符
			if c != needed {
				return false
			} else { // 如果接受到了期待字符
				tail--
				// 更新期待字符
				if tail == -1 {
					needed = ' '
				} else {
					needed = table[chs[tail]]
				}
			}
		}
	}
	if tail == -1 {
		return true
	} else {
		return false
	}
}

//leetcode submit region end(Prohibit modification and deletion)

func TestValidParentheses(t *testing.T) {

}
