//https://www.hackerrank.com/challenges/utopian-tree/problem
func utopianTree(n int32) int32 {
	var height int32 = 1

	for i:= 1; i<= int(n); i++{
		if i % 2 == 1{
			height *= 2
		} else {
			height ++
		}
	}
	return height
}
