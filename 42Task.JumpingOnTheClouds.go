// https://www.hackerrank.com/challenges/jumping-on-the-clouds/problem
package main

func jumpingOnClouds(c []int32) int32 {

	var i int = 0
	var returnValue int32 = 0

	// Do the loop
	for i < len(c) {
		if i + 2 < len(c) {
			if c[i + 2] == 0 {
				i += 2
				returnValue++
			} else if c[i + 2] == 1 {
				i += 1
				returnValue++
			}
		} else if i + 1 < len(c) {
			if c[i + 1] == 0 {
				i += 1
				returnValue++
			}
		} else if i < len(c) {
			break
		}
	}

	return returnValue
}