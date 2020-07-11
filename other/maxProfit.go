package other

func maxProfit(prices []int) int {
	lP := len(prices)
	if lP == 0 || lP == 1 {
		return 0
	}
	prevMin := prices[0]
	profit := 0

	for i := 1; i <= lP-1; i++ {
		current := prices[i]
		less := current - prevMin
		if less > profit {
			profit = less
		}
		if prevMin > prices[i] {
			prevMin = prices[i]
		}
	}
	return profit

}

func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	sum := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			sum = sum + prices[i] - prices[i-1]
		}
	}
	return sum
}
