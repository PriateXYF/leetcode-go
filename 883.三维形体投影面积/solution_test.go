package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func projectionArea(grid [][]int) int {
	xy, yz, zx := 0, 0, 0
	yzs, zxs := make([]int, len(grid)), make([]int, len(grid))
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[x]); y++ {
			h := grid[x][y]
			if h != 0 {
				xy++
			}
			if h > yzs[y] {
				yz = yz - yzs[y] + h
				yzs[y] = h
			}
			if h > zxs[x] {
				zx = zx - zxs[x] + h
				zxs[x] = h
			}
		}
	}
	return xy + yz + zx
}

//leetcode submit region end(Prohibit modification and deletion)

func TestProjectionAreaOf3dShapes(t *testing.T) {

}
