//https://www.hackerrank.com/challenges/bon-appetit/problem
func bonAppetit(bill []int32, k int32, b int32) {
	var sum,moneyBack int32

	for i := 0; i < len(bill); i++ {
		sum += bill[i]
	}

	moneyBack = b - (sum - bill[k]) / 2

	if moneyBack == 0 {
		fmt.Printf("Bon Appetit")
	} else {
		fmt.Printf("%d", moneyBack)
	}
}
