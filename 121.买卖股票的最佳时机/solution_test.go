package leetcode

import (
	"testing"
)

// leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int) int {
	minPoint, newMinPoint, profit := 0, 0, 0
	for i, price := range prices {
		// 如果还未产生可盈利股票,直接设置最小指针
		if profit == 0 && price < prices[minPoint] {
			minPoint = i
		}
		// 如果当前价格小于最小指针价格
		if price <= prices[minPoint] && price < prices[newMinPoint] {
			newMinPoint = i
		}
		// 如果当前新最小指针利润大于原最大利润,更新最小指针
		if price-prices[newMinPoint] > profit {
			minPoint = newMinPoint
			profit = price - prices[newMinPoint]
		}
		// 如果当前利润大于原最大利润,更新利润
		if price-prices[minPoint] > profit {
			profit = price - prices[minPoint]
		}
	}
	return profit
}

//leetcode submit region end(Prohibit modification and deletion)

func TestBestTimeToBuyAndSellStock(t *testing.T) {

}
