// https://www.hackerrank.com/challenges/diagonal-difference/problem
func diagonalDifference(arr [][]int32) int32 {

	var len = len(arr)
	var firstDig, secondDig, returnValue int32

	for i := 0; i < len; i++ {
		firstDig += arr[i][i]
		secondDig += arr[i][len - 1 - i]
	}

	returnValue = firstDig - secondDig
	if returnValue < 0 {
		returnValue = -returnValue
	}

	return returnValue
}


