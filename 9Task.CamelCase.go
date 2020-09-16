//https://www.hackerrank.com/challenges/camelcase/problem
func camelcase(s string) int32 {
	var returnValue int32 = 1

	for _,val := range s{
		if unicode.IsUpper(val) && unicode.IsLetter(val){
			returnValue += 1
		}
	}
	return returnValue
}
