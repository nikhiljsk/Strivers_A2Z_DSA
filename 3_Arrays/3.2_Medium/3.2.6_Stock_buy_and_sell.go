// Approach 1
func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	mini, maxi := 100000, -10000
	for _, v := range prices {
		mini = min(mini, v)
		maxi = max(maxi, v-mini)
	}
	return maxi
}

// Approach 2
func maxProfit(prices []int) int {
	buy := 100000
	var maxProfit int

	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] && buy > prices[i] {
			buy = prices[i]
		}
		if maxProfit < prices[i]-buy {
			maxProfit = prices[i] - buy
		}
	}
	if maxProfit < prices[len(prices)-1]-buy {
		maxProfit = prices[len(prices)-1] - buy
	}
	return maxProfit
}