package main

import (
	"fmt"
	"sync"
	"time"
)

/*
sync.WaitGroup用于阻塞等待一组Go协程的结束
WaitGroup内部实现了一个计数器，用来记录未完成的操作个数，它提供了三个方法
1.Add()用来添加计数
2.Done()用来在操作结束时调用，使计数减一
3.Wait()用来等待所有的操作结束，即计数变为0，该函数会在计数不为0时等待，在计数为0时立即返回
*/
func testWaitGroup() {
	var wg sync.WaitGroup

	wg.Add(2) // 因为有两个动作，所以增加2个计数
	go func() {
		defer wg.Done() // 操作完成，减少一个计数
		time.Sleep(time.Duration(1) * time.Second)
		fmt.Println("Goroutine 1")
	}()
	go func() {
		defer wg.Done() // 操作完成，减少一个计数
		time.Sleep(time.Duration(2) * time.Second)
		fmt.Println("Goroutine 2")
	}()

	wg.Wait() // 等待，直到计数为0
	fmt.Println("main routine exit.")
}
