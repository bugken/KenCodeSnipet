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

	// fmt.Println("...............defer test begin..................")
	// testDefer()
	// fmt.Println("...............defer test end..................")

	fmt.Println("...............error test begin..................")
	isTrigger := true
	e := testErrors(isTrigger)
	if e != nil {
		fmt.Println("e:", e)
	}
	fmt.Println("...............error test end..................")

	var input string
	fmt.Scanln(&input)
}
