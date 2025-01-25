package main

import "fmt"

/*
// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }
*/

func main() {
	basicForLoop()
	withNoInitValueAndPost()
}

// basic
func basicForLoop() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)
}

// with no init value and post
func withNoInitValueAndPost() {
	sum := 1

	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}
