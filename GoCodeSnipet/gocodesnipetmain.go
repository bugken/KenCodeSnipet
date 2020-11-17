package main

import (
	"os"
	"os/signal"
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

	// fmt.Println("...............error test begin..................")
	// isTrigger := true
	// e := testErrors(isTrigger)
	// if e != nil {
	// 	fmt.Println("e:", e)
	// }
	// fmt.Println("...............error test end..................")

	// fmt.Println("...............defer test begin..................")
	// testLogger()
	// fmt.Println("...............defer test end..................")

	// fmt.Println("...............file read write test begin..................")
	// testFileReadWrite()
	// fmt.Println("...............file read write test end..................")

	// fmt.Println("...............wait group test begin..................")
	// testWaitGroup()
	// fmt.Println("...............wait group test end..................")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
