// https://www.hackerrank.com/challenges/taum-and-bday/problem
func taumBday(b int32, w int32, bc int32, wc int32, z int32) int64 {

	var returnValie int32

	if (bc > wc + z) {
		returnValie = (w+b)*wc+(b*z)
	} else if (wc > bc + z) {
		returnValie = (w+b)*bc+(w*z)
	} else {
		returnValie = (w*wc)+(b*bc)
	}

	return int64(returnValie)
}