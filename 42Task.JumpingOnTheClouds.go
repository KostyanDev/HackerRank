// https://www.hackerrank.com/challenges/jumping-on-the-clouds/problem
package main

func jumpingOnClouds(c []int32) int32 {

	// Declare and initialize the variables
	var i int = 0
	var i32HoopCounter int32 = 0

	// Do the loop
	for i < len(c) {
		if i + 2 < len(c) {
			if c[i + 2] == 0 {
				i += 2
				i32HoopCounter++
			} else if c[i + 2] == 1 {
				i += 1
				i32HoopCounter++
			}
		} else if i + 1 < len(c) {
			if c[i + 1] == 0 {
				i += 1
				i32HoopCounter++
			}
		} else if i < len(c) {
			break
		}
	}

	return i32HoopCounter
}