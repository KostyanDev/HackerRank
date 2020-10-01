//https://www.hackerrank.com/challenges/jumping-on-the-clouds-revisited/problem
func jumpingOnClouds(c []int32, k int32) int32 {

	var n int = len(c)
	var currentPosition int = 0
	var returnValue int32 = 100
	var flag bool = true

	for flag {
		currentPosition = (currentPosition + int(k)) % n
		returnValue -= 2 * c[currentPosition] + 1
		if currentPosition == 0 {
			flag = false
		}
	}

	return returnValue
}
