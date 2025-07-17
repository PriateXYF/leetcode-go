package leetcode

import (
	"fmt"
	"strconv"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
var cache = map[int]string{}

func countAndSay(n int) (res string) {
	if val, exist := cache[n]; exist {
		return val
	}
	if n <= 1 {
		return "1"
	}
	last := countAndSay(n - 1)
	times := 1
	for i := 0; i < len(last)-1; i++ {
		if last[i] == last[i+1] {
			times++
		} else {
			res += strconv.Itoa(times) + string(last[i])
			times = 1
		}
	}
	res += strconv.Itoa(times) + string(last[len(last)-1])
	cache[n] = res
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCountAndSay(t *testing.T) {
	for i := range 10 {
		res := countAndSay(i)
		fmt.Println(res)
	}

}
