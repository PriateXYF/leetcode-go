package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func judgeCircle(moves string) bool {
	moveMap := map[rune][2]int{
		'U': [2]int{1, 0},
		'D': [2]int{-1, 0},
		'L': [2]int{0, -1},
		'R': [2]int{0, 1},
	}
	pos := []int{0, 0}
	for _, move := range moves {
		pos[0], pos[1] = pos[0]+moveMap[move][0], pos[1]+moveMap[move][1]
	}
	return pos[0] == 0 && pos[1] == 0
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRobotReturnToOrigin(t *testing.T) {

}
