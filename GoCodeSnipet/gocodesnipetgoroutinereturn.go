package main

import (
	"fmt"
	"sync"
)

/*通过channel拿到返回值有两种处理形式
一种是发送给一个for channel或select channel的独立goroutine中,由该独立的goroutine来处理函数的返回值.
一种是将所有goroutine的返回值都集中到当前函数, 然后统一返回给调用函数.*/

/*方法一:发送给一个for channel或select channel的独立goroutine中,由该独立的goroutine来处理函数的返回值
begin*/
var responseChannel = make(chan string, 5) //超过5个的时候，会阻塞

func httpGetMode1(url int, limiter chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	responseChannel <- fmt.Sprintf("Hello Go %d", url)
	fmt.Println("httpGetMode1 ", url)
	// time.Sleep(1 * time.Second)
	<-limiter
}

func repsonseControl() {
	// time.Sleep(1 * time.Second)
	//如果responseChannel的缓冲区为空，这个协程会一直阻塞，除非被channel被close
	for rc := range responseChannel {
		fmt.Println("reponse ", rc)
		// time.Sleep(1 * time.Second)
	}
}

func testGoRoutineReturnMode1() {
	//启动一个独立的goroutine
	go repsonseControl()
	//用于协程间的同步
	wg := &sync.WaitGroup{}
	//控制并发数
	limiter := make(chan bool, 2)
	defer close(limiter)
	//启动协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		limiter <- true
		go httpGetMode1(i, limiter, wg)
	}
	wg.Wait()
	close(responseChannel)
	fmt.Println("goroutine done.")
}

/*
方法一:发送给一个for channel或select channel的独立goroutine中,由该独立的goroutine来处理函数的返回值
end
*/

/*
方法二:所有goroutine的返回值都集中到当前函数, 然后统一返回给调用函数
begin
*/

func httpGetMode2(url int, limiter chan bool, responseChannel chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	responseChannel <- fmt.Sprintf("Hello_Go_%d", url)
	fmt.Println("httpGetMode2 ", url)
	// time.Sleep(1 * time.Second)
	<-limiter
}

func collect(urls []int) []string {
	var results []string

	wg := &sync.WaitGroup{}

	//控制并发数
	limiter := make(chan bool, 5)
	defer close(limiter)

	// 函数内的局部变量channel, 专门用来接收函数内所有goroutine的结果
	responseChannel := make(chan string, 10)
	//保证控制器内的所有值都已经正确处理完毕, 才能结束
	wgResponse := &sync.WaitGroup{}
	go func() {
		wgResponse.Add(1)
		for rc := range responseChannel {
			results = append(results, rc)
		}
		// 当responseChannel被关闭时且channel中所有的值都已经被处理完毕后, 将执行到这一行
		wgResponse.Done()
	}()

	for url := range urls {
		wg.Add(1)
		limiter <- true
		go httpGetMode2(url, limiter, responseChannel, wg)
	}

	wg.Wait()
	fmt.Println("goroutine done.")
	close(responseChannel) //关闭responseChannel，返回协程的结果
	wgResponse.Wait()
	return results
}

func testGoRoutineReturnMode2() {
	urls := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	results := collect(urls)
	fmt.Println("results:", results)
}

/*
方法二:所有goroutine的返回值都集中到当前函数, 然后统一返回给调用函数
end
*/

func testGoRoutineReturn() {
	// testGoRoutineReturnMode1()
	testGoRoutineReturnMode2()
}
