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

/*
slice用法，包括slice创建和slice传值方式
切片的创建有多种方式，可以可以使用make创建，可以使用数组创建
slice是按照类似传引用的放入传值的
slice定义如下:
type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
*/
func testSlice() {
	//直接创建slice
	var slice1 []int
	slice1 = append(slice1, 0)
	slice1 = append(slice1, 1)
	fmt.Println("slice1:", slice1)
	//直接创建并初始化
	slice2 := []int{1, 2, 3, 4, 5} // 注意这里中括号中没有长度，如果有长度就变成数组的创建了
	fmt.Println("slice2:", slice2)
	//使用make创建
	slice3 := make([]int, 5)
	slice3[0] = 0
	slice3[1] = 1
	fmt.Println("slice3:", slice3)
	//使用数组创建
	var array1 [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // 基于数组创建一个数组切片
	var slice4 []int = array1[:5]
	fmt.Println("slice4:", slice4)
}

//关键字defer测试
func testDefer() {

}
