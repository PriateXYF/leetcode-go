package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func countDays(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	if len(meetings) == 1 {
		return days - (meetings[0][1] - meetings[0][0] + 1)
	}
	meetingDays := 0
	for i := 1; i < len(meetings); i++ {
		// 如果有重叠则合并到下一个数组中
		if meetings[i-1][1] >= meetings[i][0] {
			meetings[i][0] = meetings[i-1][0]
			meetings[i][1] = max(meetings[i-1][1], meetings[i][1])
		} else {
			meetingDays += meetings[i-1][1] - meetings[i-1][0] + 1
		}
	}
	meetingDays += meetings[len(meetings)-1][1] - meetings[len(meetings)-1][0] + 1
	return days - meetingDays
}

//leetcode submit region end(Prohibit modification and deletion)

func TestCountDaysWithoutMeetings(t *testing.T) {
	fmt.Println(countDays(14, [][]int{{6, 11}, {7, 13}, {8, 9}, {5, 8}, {3, 13}, {11, 13}, {1, 3}, {5, 10}, {8, 13}, {3, 9}}))
}
