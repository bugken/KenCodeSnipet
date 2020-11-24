package main

import (
	"fmt"
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

	// fmt.Println("...............http server test begin..................")
	// go testHTTPServer()
	// fmt.Println("...............http server test end..................")

	// fmt.Println("...............http client test begin..................")
	// testHTTPClient()
	// fmt.Println("...............http client test end..................")

	fmt.Println("...............http go routine return test begin..................")
	testGoRoutineReturn()
	fmt.Println("...............http go routine return test end..................")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
