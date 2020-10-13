//https://www.hackerrank.com/challenges/cut-the-sticks/problem
func cutTheSticks(arr []int32) []int32 {

	// Declare and initialize the variables
	var returnValue []int32

	// Do the loop
	for len(arr) > 0 {
		returnValue = append(returnValue, int32(len(arr)))

		var minElevent int32 = arr[0]
		for i := 1; i < len(arr); i++ {
			if arr[i] < minElevent {
				minElevent = arr[i]
			}
		}

		var tmpStick []int32
		var lowStickLength  int32


		for index,_ := range arr{
			lowStickLength = arr[index] - minElevent
			if lowStickLength > 0 {
				tmpStick = append(tmpStick, lowStickLength)
			}
		}
		arr = tmpStick
	}

	return returnValue
}
