package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
// 计算小球对应的编号
func calcBallNumber(ball int) (num int) {
	for ball != 0 {
		num += ball % 10
		ball /= 10
	}
	return num
}

// 计算盒子中小球的最大数量
func countBalls(lowLimit int, highLimit int) (maxNum int) {
	table := map[int]int{}
	for i := lowLimit; i <= highLimit; i++ {
		num := calcBallNumber(i)
		if _, exists := table[num]; exists {
			table[num]++
		} else {
			table[num] = 1
		}
		maxNum = max(table[num], maxNum)
	}
	return maxNum
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMaximumNumberOfBallsInABox(t *testing.T) {

}
