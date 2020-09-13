package main


import (
"fmt"
)


func main() {

	var name string
	var surname string

	fmt.Scan(&name)


	if (len(name) <= 10) {
		fmt.Scan(&surname)
		if (len(surname) <= 10) {
			fmt.Printf("Hello %s %s! You just delved into Golang \n", name, surname)
		} else {
			fmt.Printf( "So long surname \n")
		}
	} else {
		fmt.Printf( "So long name \n")
	}
}