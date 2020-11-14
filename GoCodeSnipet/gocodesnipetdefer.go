package main

import "fmt"

/*
defer和go一样都是Go语言提供的关键字。defer用于资源的释放，会在函数返回之前进行调用。
关键字defer允许我们推迟到函数返回之前（或任意位置执行return语句之后）一刻才执行某个语句或函数（为什么要在返回之后才执行这些语句？因为return语句同样可以包含一些操作，而不是单纯地返回某个值）。
关键字defer的用法类似于面向对象编程语言Java和C#的finally语句块，它一般用于释放某些已分配的资源。
通常我们会将一些函数的收尾工作通过defer执行，使代码结构更清晰。
*/
/*使用defer实现代码跟踪，或者函数耗时计算*/
func deferTrace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func deferUntrace(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer deferUntrace(deferTrace("a"))
	fmt.Println("in a")
}

func b() {
	defer deferUntrace(deferTrace("b"))
	fmt.Println("in b")
	a()
}

//关键字defer测试
func testDefer() {
	b()
}
