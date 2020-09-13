
// https://www.hackerrank.com/challenges/a-very-big-sum/problem
// https://www.hackerrank.com/challenges/simple-array-sum/problem - same
func aVeryBigSum(ar []int64) int64 {
	var sum int64
	for _, val := range ar {
		sum += val
	}
	return sum
}
