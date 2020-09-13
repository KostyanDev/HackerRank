//https://www.hackerrank.com/challenges/plus-minus/problem

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
	var l,positive,negative,zeros int32
	l = int32(len(arr))

	for _,t := range arr {
		if t > 0 {
			positive++
		}else if t <0 {
			negative++
		}else {
			zeros++
		}
	}

	positivef := float64(positive)/float64(l)
	negativef := float64(negative)/float64(l)
	zerosf := float64(zeros)/float64(l)

	fmt.Printf("%.6f\n",positivef)
	fmt.Printf("%.6f\n",negativef)
	fmt.Printf("%.6f\n",zerosf)

}
