package leetcode

import (
	"fmt"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func earliestAndLatest(n int, firstPlayer int, secondPlayer int) []int {
	n1, n2 := n, n
	left1, mid1, right1 := firstPlayer-1, secondPlayer-firstPlayer-1, n-secondPlayer
	left2, mid2, right2 := left1, mid1, right1
	res := []int{1, 1}
	// 最早对战的情况
	for ; left1 != right1; res[0]++ {
		eliminate := n1 / 2            // 第 n 轮需要淘汰的人数
		if right1-left1 >= eliminate { // 右边严重偏多
			right1 -= eliminate
		} else if left1-right1 >= eliminate { // 左边严重偏多
			left1 -= eliminate
		} else { // 两边相差不大
			var diff int
			if left1 < right1 {
				diff = left1 - right1
				right1 = left1
			} else {
				diff = right1 - left1
				left1 = right1
			}
			remainder := eliminate - diff
			if remainder%2 == 0 || (remainder%1 == 0 && mid1 > 0) { // 如果余下来的是偶数或中间有数可以抵消奇数
				continue
			}
			// 中间无数且余下是奇数
			left1, right1 = left1-remainder/2, right1-remainder/2+1
		}
		n1 -= eliminate
	}
	// 最晚对战的情况
	for ; left2 != right2; res[1]++ {
		mid := 0
		if n2%2 == 1 {
			mid = n2/2 + 1 // 中间位轮空
		}
		fmt.Println(n2, left2, mid2, right2)
		eliminate := n2 / 2 // 第 n 轮需要淘汰的人数
		if left2 < right2 { // 左边小于右边
			if left2 >= eliminate {
				left2 -= eliminate
			} else if (left2 + mid2) >= eliminate {
				if mid != 0 && mid <= (left2+mid2+1) { // 中间位轮空
					mid2 -= eliminate - left2 - 1
					right2--
				} else {
					mid2 -= eliminate - left2
				}
				left2 = 0
			} else {
				right2 -= eliminate - (left2 + mid2)
				mid2, left2 = 0, 0
			}
		} else { // 右边小于左边
			if right2 >= eliminate {
				right2 -= eliminate
			} else if (right2 + mid2) >= eliminate {
				if mid != 0 && mid >= (right2+mid2+1) { // 中间位轮空
					mid2 -= eliminate - right2 - 1
					right2--
				} else {
					mid2 -= eliminate - right2
				}
				right2 = 0
			} else {
				left2 -= eliminate - (right2 + mid2)
				mid2, right2 = 0, 0
			}
		}
		n2 -= eliminate
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestTheEarliestAndLatestRoundsWherePlayersCompete(t *testing.T) {
	res := earliestAndLatest(28, 1, 26)
	fmt.Println(res)
	res = earliestAndLatest(28, 1, 27)
	fmt.Println(res)
}
