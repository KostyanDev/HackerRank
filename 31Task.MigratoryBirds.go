//https://www.hackerrank.com/challenges/migratory-birds/problem

func migratoryBirds(arr []int32) int32 {

	count := [1000]int{}

	for i := 0; i < len(arr); i++ {
		count[int(arr[i])]++
	}

	var returnValue int = 1
	var maxValue int = count[returnValue]

	for i := 2; i < len(count); i++ {
		if count[i] > maxValue {
			returnValue, maxValue = i, count[i]
		}
	}
	return int32(returnValue)
}
