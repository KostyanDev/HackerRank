//https://www.hackerrank.com/challenges/chocolate-feast/problem
func chocolateFeast(n int32, c int32, m int32) int32 {
	var wrap,returnValue int32

	wrap = n/c
	returnValue = wrap

	for wrap >= m{
		wrap = wrap-m
		wrap++
		returnValue++
	}
	return returnValue
}
