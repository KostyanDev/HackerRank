//https://www.hackerrank.com/challenges/between-two-sets/problem
// подсмотрел вариант решения через Euclidean algorithm
func biggestMultiple(a, b int32) int32 {
	for a % b != 0 {
		a, b = b, (a % b)
		fmt.Println("gcd B: ", b)
	}
	return b
}

func biggestMultipleArray(arr []int32) int32 {
	var returnValue int32 = arr[0]
	for i := 1; i < len(arr); i++ {
		returnValue = biggestMultiple(returnValue, arr[i])
	}
	return returnValue
}

func littlesMultipe(a, b int32) int32 {
	return a * b / biggestMultiple(a, b)
}

func littlesMultipleArray(arr []int32) int32 {
	var returnValue int32 = arr[0]
	for i := 1; i < len(arr); i++ {
		returnValue = littlesMultipe(returnValue, arr[i])
	}
	return returnValue
}

func getTotalX(a []int32, b []int32) int32 {

	var min int32 = littlesMultipleArray(a)
	var max int32 = biggestMultipleArray(b)
	var current int32 = min
	var returnValue int32 = 0

	for current <= max {
		if max % current == 0 {
			returnValue++
		}
		current += min
	}

	return returnValue
}