package main

import (
	"fmt"
)

func checkNumberType(number int) string {
	if number%2 == 0 {
		if number > 0 {
			return "Even and positive\n"
		} else if number < 0 {
			return "Even and negative\n"
		} else {
			return "This is zero\n"
		}
	} else {
		if number > 0 {
			return "Odd and positive\n"
		} else if number < 0 {
			return "Odd and negative\n"
		} else {
			return "This is zero\n"
		}
	}
}

func main() {
	// input Number here
	var num int = 0

	// check Number
	result := checkNumberType(num)
	fmt.Print(result)
}
