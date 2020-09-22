//https://www.hackerrank.com/challenges/apple-and-orange/problem
func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {

	var currentPosition int32
	var apple, orange int


	for _,val := range apples{
		currentPosition = a + val
		if (currentPosition >= s) && (currentPosition <= t) {
			apple++
		}
	}
	for _,val := range oranges{
		currentPosition = b + val
		if (currentPosition >= s) && (currentPosition <= t) {
			orange++
		}
	}
	fmt.Println(apple)
	fmt.Println(orange)
}

