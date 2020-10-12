//https://www.hackerrank.com/challenges/kangaroo/problem
func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {

	var returnValue string = "NO"
	var currentX1, currentX2 int32 = x1, x2

	if v1 > v2 {

		for currentX1 < currentX2 {
			currentX1 += v1
			currentX2 += v2
		}

		if currentX1 == currentX2 {
			returnValue = "YES"
		}
	}
	return returnValue
}

