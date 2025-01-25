package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	ifElseStatement()
	switchStatement()
	fmt.Println(pow(3, 2, 10))
}

/*
  - Syntax :
    if condition {
    code to be executed if condition is true
    } else {
    something else to happen
    }

    // if statement initialization
    if  var declaration;  condition {
    // code to be executed if condition is true
    }
    }
*/

func ifElseStatement() {
	if 10%2 == 0 {
		fmt.Println("10 is even")
	} else {
		fmt.Println("10 is odd")
	}

	if 8%2 == 0 || 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	// if statement initialization
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

/*
- Switch Case
- Syntax
	switch expression {
	case condition:

	}
*/

func switchStatement() {
	today := time.Now()

	switch today.Day() {
	case 5:
		fmt.Println("Today is 5th. Clean your house.")
	case 10:
		fmt.Println("Today is 10th. Buy some wine.")
	case 15:
		fmt.Println("Today is 15th. Visit a doctor.")
	case 20:
		fmt.Println("Today is 20th. Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}
}

// If with a short statement
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
