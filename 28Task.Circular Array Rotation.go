// https://www.hackerrank.com/challenges/circular-array-rotation/problem
func mod(a, b int32) int32 {
	return (a % b + b) % b
}

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {

	var returnValue []int32
	var length int32 = int32(len(a))


	for i := 0; i < len(queries); i++ {
		returnValue = append(returnValue, a[mod(queries[i] - k, length)])
	}
	return returnValue
}