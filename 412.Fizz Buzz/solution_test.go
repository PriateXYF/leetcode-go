package leetcode

import (
	"strconv"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func fizzBuzz(n int) (res []string) {
	for i := range n {
		i++
		if i%5 == 0 && i%3 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%5 == 0 {
			res = append(res, "Buzz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else {
			res = append(res, strconv.Itoa(i))
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestFizzBuzz(t *testing.T) {

}
