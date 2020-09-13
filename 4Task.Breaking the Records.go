
// https://www.hackerrank.com/challenges/breaking-best-and-worst-records/problem
// Complete the breakingRecords function below.
func breakingRecords(scores []int32) []int32 {
	var min,max,cmin,cmax,i int32
	var returnValue []int32

	for _,t := range scores {
		if i ==0 {
			min = t
			max = t
			i++
		}else{
			if t < min {
				cmin++
				min = t
			}
			if t > max{
				max = t
				cmax++
			}
		}
	}
	returnValue = append(returnValue, cmax, cmin)
	return returnValue
}