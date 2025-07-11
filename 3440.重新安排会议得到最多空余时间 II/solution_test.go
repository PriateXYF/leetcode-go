package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func maxFreeTime(eventTime int, startTime []int, endTime []int) (res int) {
	// 先标记是否可以完全移出该会议
	removable := make([]bool, len(startTime))
	idleStart, meetingTime := 0, 0
	// 从左往右遍历
	for i, maxDuration, maxIndex := 0, startTime[0], 0; i < len(startTime); i++ {
		// 如果可完全移出
		if endTime[i]-startTime[i] <= maxDuration && (maxIndex != i) {
			removable[i] = true
		}
		// 找到最大空闲时间
		if i >= 1 && startTime[i]-endTime[i-1] > maxDuration {
			maxIndex, maxDuration = i, startTime[i]-endTime[i-1]
		}
		// 找到相邻两空闲时间的最大值
		if i >= 1 {
			idleTime := startTime[i] - idleStart - meetingTime
			idleStart = endTime[i-1]
			meetingTime -= endTime[i-1] - startTime[i-1]
			res = max(idleTime, res)
		}
		meetingTime += endTime[i] - startTime[i]
	}
	res = max(eventTime-idleStart-meetingTime, res)
	// 从右往左遍历
	for i, maxDuration, maxIndex := len(startTime)-1, eventTime-endTime[len(endTime)-1], len(startTime)-1; i >= 0; i-- {
		// 如果可完全移出
		if endTime[i]-startTime[i] <= maxDuration && (maxIndex != i) {
			removable[i] = true
		}
		// 找到最大空闲时间
		if i < len(startTime)-1 && startTime[i+1]-endTime[i] > maxDuration {
			maxIndex, maxDuration = i, startTime[i+1]-endTime[i]
		}
		if removable[i] {
			if i == 0 {
				res = max(res, startTime[1])
			} else if i == len(startTime)-1 {
				res = max(res, eventTime-endTime[i-1])
			} else {
				res = max(res, startTime[i+1]-endTime[i-1])
			}

		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRescheduleMeetingsForMaximumFreeTimeIi(t *testing.T) {
	maxFreeTime(34, []int{0, 17}, []int{14, 19})
}
