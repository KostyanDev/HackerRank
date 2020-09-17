//https://www.hackerrank.com/challenges/time-conversion/problem

func timeConversion(s string) string {
	var returnValue string

	newReturnValue, _ := time.Parse("15:04:05PM", s)
	returnValue = newReturnValue.Format("15:04:05")

	return returnValue
}