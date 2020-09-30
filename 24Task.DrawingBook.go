//https://www.hackerrank.com/challenges/drawing-book/problem
package main

import "fmt"

func min(a, b int32) int32 {
	if a < b {
		return a
	} else {
		return b
	}
}

func pageCount(n int32, p int32) int32 {

	var returnValue int32

	if n%2 == 0 {
		returnValue = min(p, n+1-p) / 2
		fmt.Println("returnValue div 2", returnValue)
	} else {
		returnValue = min(p, n-p) / 2
		fmt.Println("returnValue !div 2", returnValue)
	}

	return returnValue
}

func main() {
	var n, p int32 = 500, 47

	fmt.Println("result is :", pageCount(n, p))

}
