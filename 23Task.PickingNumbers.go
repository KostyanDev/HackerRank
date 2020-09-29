//https://www.hackerrank.com/challenges/picking-numbers/problem
func pickingNumbers(a []int32) int32 {

	var currentCount, returnValue int32
	elementCounter := make(map[int32]int32)

	for i := 0; i < len(a); i++ {
		_, isBe := elementCounter[a[i]]
		if !isBe {
			elementCounter[a[i]] = 1
		} else {
			elementCounter[a[i]]++
		}
	}

	for key, val := range elementCounter {
		nextValue, isBe := elementCounter[key+1]
		if isBe {
			currentCount = val + nextValue
		} else {
			currentCount = val
		}
		if currentCount > returnValue {
			returnValue = currentCount
		}
	}
	return returnValue
}


