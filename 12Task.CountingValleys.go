//https://www.hackerrank.com/challenges/counting-valleys/problem
func countingValleys(steps int32, path string) int32 {
	var returnValue, positionCurren, positionPrevious int32

	for i:= 0; i < len(path); i++{
		positionPrevious = positionCurren
		if path[i] == 'D'{
			positionCurren--
		} else if path[i] == 'U'{
			positionCurren++
		}

		if positionPrevious < 0 && positionCurren >= 0{
			returnValue++
		}
	}

	return returnValue


	//inValleys := false
	//
	//for _,val := range path{
	//	if val == rune('D'){
	//		position--
	//	} else if val == rune('U'){
	//		position++
	//	}
	//	if inValleys && val == 0{
	//		inValleys = false
	//		returnValue++
	//	} else if !inValleys && pos < 0{
	//		inValleys = true
	//	}
	//
	//}

	//return returnValue
}