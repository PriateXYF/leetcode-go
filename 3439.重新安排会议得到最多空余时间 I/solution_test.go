package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) (res int) {
	idleStart, meetingTime := 0, 0
	for meeting := range len(startTime) {
		if meeting >= k {
			idleTime := startTime[meeting] - idleStart - meetingTime
			idleStart = endTime[meeting-k]
			meetingTime -= endTime[meeting-k] - startTime[meeting-k]
			res = max(idleTime, res)
		}
		meetingTime += endTime[meeting] - startTime[meeting]
	}
	res = max(eventTime-idleStart-meetingTime, res)
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRescheduleMeetingsForMaximumFreeTimeI(t *testing.T) {
	maxFreeTime(10, 2, []int{0, 1, 2, 3, 4}, []int{1, 2, 3, 4, 5})
}
