package main

import (
	"fmt"
)

func main() {
	printVariablesWithInitializers()
	printShortVariableDeclarations()
	printMultipleValues()
	printConstants()
	printNamesWithDefaultFormat()
	printNamesWithWhiteSpaceAndNewLineAtEnd()
	printVariablesWithTypeAndValue()
}

func printVariablesWithInitializers() {
	var student1 string = "mazen" //type is string
	var student2 = "fteha"        //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
}

func printShortVariableDeclarations() {
	x := 2 //type is inferred
	fmt.Println(x)
}

func printMultipleValues() {
	var a, b, c, d int = 1, 3, 5, 7
	fmt.Println(a, b, c, d)
}

func printConstants() {
	const PI = 3.14
	fmt.Println(PI)
}

/*
Print()
Println()
Printf()
*/

// Print()
func printNamesWithDefaultFormat() {
	// The Print() function prints its arguments with their default format.
	var i, j string = "mazen", "fteha"

	fmt.Print(i)
	fmt.Print(j)

	fmt.Print(i, "\n", j, "\n")
}

// PrintIn()

func printNamesWithWhiteSpaceAndNewLineAtEnd() {
	firstName := "Mazen"
	secondName := "Fteha"

	fmt.Println(firstName, secondName)
}

// The Printf() Function
func printVariablesWithTypeAndValue() {
	var i string = "Hello"
	var j int = 15

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T", j, j)
}
