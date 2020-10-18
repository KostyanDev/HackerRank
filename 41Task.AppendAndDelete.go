// https://www.hackerrank.com/challenges/append-and-delete/problem
func appendAndDelete(s string, t string, k int32) string {
	var returnValue, min string
	var length,btw, count int32

	if len(s) <= len(t) {
		min = s
	} else {
		min = t
	}


	for i := 0; i < len(min); i++ {
		if s[i] != t[i] {
			break
		}
		count++
	}

	var lenS,lenT int32 = int32(len(s)), int32(len(t))
	btw = (lenS - count) + (lenT - count)
	length = lenS + lenT

	switch {
	case btw > k:
		returnValue = "No"
	case btw == k:
		returnValue = "Yes"
	case length % 2 == k%2:
		returnValue = "Yes"
	case length < k:
		returnValue = "Yes"
	default:
		returnValue = "No"
	}
	return returnValue
}