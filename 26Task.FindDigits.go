//https://www.hackerrank.com/challenges/find-digits/problem?h_r=next-challenge&h_v=zen
func findDigits(n int32) int32 {

	var returnValue, number int32 = 0 , n
	var current int32

	for number != 0 {

		current = number % 10

		if current != 0 {
			if n % current == 0 {
				returnValue++
			}
		}
		number /= 10
	}
	return returnValue
}