//https://www.hackerrank.com/challenges/sherlock-and-squares/problem
func squares(a int32, b int32) int32 {
	return int32(math.Sqrt(float64(b))) - int32(math.Sqrt(float64(a - 1)))
}
