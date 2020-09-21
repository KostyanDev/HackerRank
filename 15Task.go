// https://www.hackerrank.com/challenges/electronics-shop/problem
func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {

	var maxPrice int32 = -1
	var currentPrice int32

	for _,val1 := range keyboards{
		for _,val2 := range drives{
			currentPrice = val1 + val2
			if currentPrice <= b && currentPrice > maxPrice {
				maxPrice = currentPrice
			}
		}
	}

	return maxPrice
}
