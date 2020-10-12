//https://www.hackerrank.com/challenges/permutation-equation/problem
func permutationEquation(p []int32) []int32 {

	var arr = make(map[int32]int32)
	var returnValue []int32

	for i := 0; i < len(p); i++ {
		arr[p[i]] = int32(i) + 1
	}

	for i := 0; i < len(p); i++ {
		returnValue = append(returnValue, arr[arr[int32(i) + 1]])
	}

	return returnValue
}
