//https://www.hackerrank.com/challenges/cats-and-a-mouse/problem
func abs(n int32) int32 {
	if n >= 0 {
		return n
	} else {
		return -n
	}
}

// Complete the catAndMouse function below.
func catAndMouse(x int32, y int32, z int32) string {

// Declare and initialize the variables
	var distanceAtoC int32 = abs(x - z)
	var distanseBtoC int32 = abs(y - z)
	var returnValue string

// Do the logic
	if distanceAtoC < distanseBtoC {
		returnValue = "Cat A"
	} else if distanceAtoC > distanseBtoC {
		returnValue = "Cat B"
	} else {
		returnValue = "Mouse C"
	}

return sOutput
}
