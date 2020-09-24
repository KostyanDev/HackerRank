//https://www.hackerrank.com/challenges/angry-professor/problem
func angryProfessor(k int32, a []int32) string {

	var returnValue string = "YES"
	var students int32 = 0

	for _,val := range a {
		if val <= 0 {
			students++
		}
		if students >= k {
			returnValue = "NO"
		}
	}

	return returnValue
}