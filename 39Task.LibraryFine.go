//https://www.hackerrank.com/challenges/library-fine/problem

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {

	var returnValue int32

	if y1 < y2 {
		returnValue = 0
	} else if y1 == y2 {
		if m1 < m2 {
			returnValue = 0
		} else if m1 == m2 {
			if d1 <= d2 {
				returnValue = 0
			} else { // d1 > d2
				returnValue = 15 * (d1 - d2)
			}
		} else { // m1 > m2
			returnValue = 500 * (m1 - m2)
		}
	} else { // y1 > y2
		returnValue = 10000
	}

	return returnValue
}
