// https://www.hackerrank.com/challenges/sock-merchant/problem
func sockMerchant(n int32, ar []int32) int32 {

	var returnValue, i int32 = 0, 0

	for index,_ := range ar {
		if ar[index] != 0 {
			for i = int32(index + 1); i < n; i++ {
				if ar[index] == ar[i] {
					returnValue++
					ar[i] = 0
					break
				}
			}
		}
	}
	return returnValue
}
