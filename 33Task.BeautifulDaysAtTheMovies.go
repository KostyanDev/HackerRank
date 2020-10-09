package main

import "strconv"

func abs(n int32) int32 {
	if n >= 0 {
		return n
	} else {
		return -n
	}
}

// Complete the beautifulDays function below.
func abs(n int32) int32 {
	if n >= 0 {
		return n
	} else {
		return -n
	}
}


func beautifulDays(i int32, j int32, k int32) int32 {

	// Declare and initialize the variables
	var currentNubmer, reversNubmer string
	var currentReversNumber int32
	var returnValue int32 = 0

	// Do the loop
	for p := i; p <= j; p++ {
		currentNubmer = strconv.FormatInt(int64(p), 10)

		reversNubmer = ""
		for q := len(currentNubmer) - 1; q >= 0; q-- {
			reversNubmer += string(currentNubmer[q])
		}

		reversNumberInt64, err := strconv.ParseInt(reversNubmer, 10, 64)
		if err != nil {
			panic(err)
		}
		currentReversNumber = int32(reversNumberInt64)

		if abs(p - currentReversNumber) % k == 0 {
			returnValue++
		}
	}

	return returnValue
}
