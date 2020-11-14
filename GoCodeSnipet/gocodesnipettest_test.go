package main

import (
	"strconv"
	"testing"
)

/*
go test命令只能在一个相应的目录下执行所有文件

1.需要创建一个名称以 _test.go 结尾的文件
2.该文件包含 测试用例 (如：TestXxx 函数)

测试用例有四种形式：
1.TestXxxx(t *testing.T) // 基本测试用例
2.BenchmarkXxxx(b *testing.B) // 压力测试的测试用例
3.Example_Xxx() // 测试控制台输出的例子
4.TestMain(m *testing.M) // 测试 Main 函数

函数中通过调用testing.T的Error, Errorf, FailNow, Fatal, FatalIf方法，说明测试不通过，调用Log方法用来记录测试的信息
*/

// 基本测试用例 testing.T 普通测试用例
func TestJoin(t *testing.T) {
	var (
		expected = "ab" // 期望结果
	)
	actual := testJoin("a", "b")
	if actual != expected {
		t.Errorf("业务结果没有达到期望的结果! \n 业务结果： %s，期望的结果： %s \n", actual, expected)
	} else {
		t.Log("测试通过了") // 记录信息
	}
}

// 压力测试的测试用例 testing.B 压力测试
func BenchmarkJoin(t *testing.B) {
	for i := 0; i < t.N; i++ {
		testJoin(strconv.Itoa(i))
	}
}

// 基本测试用例 testing.T 普通测试用例
func TestDivision(t *testing.T) {
	if i, e := testDivision(6, 2); i != 3 || e != nil { //try a unit test on function
		t.Error("除法函数测试没通过") // 如果不是如预期的那么就报错
	} else {
		t.Log("第一个测试通过了") // 记录一些你期望记录的信息
	}
}

// 压力测试的测试用例 testing.B 压力测试
func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ { //use b.N for looping
		testDivision(4, 5)
	}
}

// 压力测试的测试用例 testing.B 压力测试
func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() // 调用该函数停止压力测试的时间计数

	// 做一些初始化的工作,例如读取文件数据,数据库连接之类的,
	// 这样这些时间不影响我们测试函数本身的性能

	b.StartTimer() // 重新开始时间
	for i := 0; i < b.N; i++ {
		testDivision(4, 5)
	}
}
