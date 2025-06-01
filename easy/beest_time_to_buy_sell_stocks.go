package easy

func maxProfit(prices []int) int {
	minPrice := 9999999
	ans := 0
	for _, val := range prices {
		minPrice = min(minPrice, val)
		ans = max(ans, val-minPrice)
	}
	return ans
}
