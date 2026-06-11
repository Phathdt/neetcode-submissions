func maxProfit(prices []int) int {
	ans := 0
	minPrice := prices[0]

	for _, price := range prices {
		minPrice = min(minPrice, price)

		ans = max(ans, price-minPrice)
	}

	return ans
}
