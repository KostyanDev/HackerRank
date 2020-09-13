//https://www.hackerrank.com/challenges/compare-the-triplets/problem

func compareTriplets(a []int32, b []int32) []int32 {
	var alice,bob,i  int32
	var returnValue [] int32
	for i = 0; i < int32(len(a)); i ++ {

		if a[i] > b[i] {
			alice += 1
		} else if b[i] > a[i] {
			bob += 1
		}
	}
	returnValue = append(returnValue, alice, bob)
	return returnValue
}