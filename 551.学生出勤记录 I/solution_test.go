package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func checkRecord(s string) bool {
	absentTimes, lateTimes := 0, 0
	for _, status := range s {
		switch status {
		case 'A':
			lateTimes = 0
			absentTimes++
			if absentTimes >= 2 {
				return false
			}
		case 'L':
			lateTimes++
			if lateTimes >= 3 {
				return false
			}
		case 'P':
			lateTimes = 0
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)

func TestStudentAttendanceRecordI(t *testing.T) {

}
