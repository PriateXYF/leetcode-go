package leetcode

import (
	"sort"
	"strconv"
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func findRelativeRanks(scores []int) []string {
	type Athletes struct {
		index int
		score int
	}
	var athletes []Athletes
	for i, s := range scores {
		athletes = append(athletes, Athletes{i, s})
	}
	sort.Slice(athletes, func(i int, j int) bool { return athletes[i].score > athletes[j].score })
	res := make([]string, len(scores))
	reward := [3]string{"Gold Medal", "Silver Medal", "Bronze Medal"}
	for rank, athlete := range athletes {
		if rank < 3 {
			res[athlete.index] = reward[rank]
		} else {
			res[athlete.index] = strconv.Itoa(rank + 1)
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)

func TestRelativeRanks(t *testing.T) {

}
