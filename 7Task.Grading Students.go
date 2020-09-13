//https://www.hackerrank.com/challenges/grading/problem

func gradingStudents(grades []int32) []int32 {
	// Write your code here
	var result []int32

	for _, tmp := range grades {
		if tmp<38{
			result = append(result, int32(tmp))
		}else{
			if tmp % 5 > 2 {
				// round our number
				fmt.Println(((tmp/5)+1)*5)
				i := ((tmp/5)+1)*5
				result = append(result, int32(i))
			}else {
				result = append(result, int32(tmp))
			}
		}
	}
	return result
}