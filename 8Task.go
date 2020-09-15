// https://www.hackerrank.com/challenges/birthday-cake-candles/problem
func birthdayCakeCandles(candles []int) int {
	var returnValue, i, max int

	for _,val := range candles{
		if val > max{
			max = val
		}
	}
	for i = 0;i<len(candles);i++ {
		if candles[i] == max {
			returnValue++;
		}
	}
	return returnValue
}