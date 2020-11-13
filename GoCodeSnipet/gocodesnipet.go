package main

import (
	"fmt"
	"time"
)

/*
Ticker是间隔特定时间触发
*/
func testTicker() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("ticker timeout...")
		}
	}
}

/*
Timer只触发一次，如果还需继续触发需要使用Reset进行重置
Timer可以通过NewTimer创建或者使用AfterFunc创建
Timer类型定义如下:
type Timer struct {
    C <-chan Time     // The channel on which the time is delivered.
    r runtimeTimer
}
*/
func testTimer() {
	duration := time.Duration(time.Second * 2)
	timer := time.NewTimer(duration)
	defer timer.Stop()
	//第一种实现方法
	for i := 0; i < 5; i++ {
		// 等待channel C中的信号
		<-timer.C
		fmt.Println("timer timeout...")
		// 重置定时器，因为Timer是一次触发
		// Reset 会先调用 stopTimer 再调用 startTimer，类似于废弃之前的定时器，重新启动一个定时器
		timer.Reset(time.Second * 2)
	}
	//第二种实现方法
	time.AfterFunc(time.Second*2, func() {
		fmt.Println("timer trigger in AfterFunc mode...")
		go testTimer()
	})
}

/*
*定时器测试代码
*goland的定时器有两种实现方式:分别为Timer和Ticker
 */
func testCodeTimer() {
	//testTimer()
	testTicker()
}

func main() {
	fmt.Println("...............timer test begin..................")
	testCodeTimer()
	fmt.Println("...............timer test end..................")

	var input string
	fmt.Scanln(&input)
}
