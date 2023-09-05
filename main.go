package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Is this year leap? %t", isLeapYear(2000))
}

func isLeapYear(yearNumber int) bool {
	if yearNumber%4 == 0 && yearNumber%10 == 0 {
		return true
	} else {
		return false
	}
}
