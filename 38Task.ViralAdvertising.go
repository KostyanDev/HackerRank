// https://www.hackerrank.com/challenges/strange-advertising/problem
func viralAdvertising(n int32) int32 {

	var received, returnValue, i int32 = 5, 0, 0

	for i = 0; i < n;  i++ {
		received /= 2;
		returnValue += received;
		received *= 3;
	}
	return returnValue
}


