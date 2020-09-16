// https://www.hackerrank.com/challenges/birthday-cake-candles/problem
func birthdayCakeCandles(candles []int) int {
	var returnValue, i, max int

	for _,val := range candles{
		if val > max{
			max = val
		}
	}
	for _,val := range candles {
		if val == max {
			returnValue++;
		}
	}
	return returnValue
}