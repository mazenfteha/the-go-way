package main

import (
	"fmt"
	"reflect"
)

func main() {
	integer()
}

func integer() {
	//declaration
	var a int

	b := 15

	fmt.Printf("a has value: %v and type: %T\n", a, a)
	fmt.Printf("b has value: %v\n", b)
	fmt.Println("b has type of", reflect.TypeOf(b))
}
