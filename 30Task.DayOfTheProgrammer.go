package main

import "strconv"

func hightCalendar(year int32) bool {
	return year%4 == 0
}

func gregCalendar(year int32) bool {
	return (year%400 == 0) || ((year%4 == 0) && (year%100 != 0))
}

func dayOfProgrammer(year int32) string {

	var dayMonth string

	if year < 1918 {
		if hightCalendar(year) {
			dayMonth = "12.09."
		} else {
			dayMonth = "13.09."
		}
	} else if year == 1918 {
		dayMonth = "26.09."
	} else {
		if gregCalendar(year) {
			dayMonth = "12.09."
		} else {
			dayMonth = "13.09."
		}
	}

	return dayMonth + strconv.FormatInt(int64(year), 10)
}
