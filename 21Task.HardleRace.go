// https://www.hackerrank.com/challenges/the-hurdle-race/problem
func hurdleRace(k int32, height []int32) int32 {

	var maxHeight int32 = height[0]
	var returnValue int32

	for _,val := range height{
		if val > maxHeight{
			maxHeight = val
		}
	}
	returnValue = maxHeight - k

	if returnValue < 0 {
		returnValue = 0
	}

	return returnValue
}

