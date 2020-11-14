package main

import (
	"fmt"
)

func main() {
	// fmt.Println("...............timer test begin..................")
	// testCodeTimer()
	// fmt.Println("...............timer test end..................")
	// fmt.Println("...............slice test begin..................")
	// testSlice()
	// fmt.Println("...............slice test end..................")
	fmt.Println("...............defer test begin..................")
	testDefer()
	fmt.Println("...............defer test end..................")

	var input string
	fmt.Scanln(&input)
}
