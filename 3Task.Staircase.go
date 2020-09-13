// https://www.hackerrank.com/challenges/staircase/problem

func staircase(n int32) {

var i int32
var j int32
	for i =0;i<n;i++{
		for j =0;j<n-i-1;j++{
			fmt.Print(" ")
		}
		for j =0;j<=i;j++{
			fmt.Print("#")
		}
	fmt.Println()
	}

}
