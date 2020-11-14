package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

/*
读取文件有三种方式:1.将整个文件读入内存 2.按照字节数读取 3.按行读取
*/

//1.将这个文件读到内存
func readFile2Memory() {
	file, err := os.Open("./log.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	fmt.Println(string(content))
}
func readFile2Memory2() {
	filepath := "./log.txt"
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}

//2.按照字节读取文件，本方法读效率最高
func readFileByBytes() {
	filepath := "./log.txt"
	fi, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	r := bufio.NewReader(fi)
	chunks := make([]byte, 0)
	buf := make([]byte, 1024) //一次读取多少个字节
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	fmt.Println(string(chunks))
}
func readFileByBytes2() {
	file := "./log.txt"
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	chunks := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	fmt.Println(string(chunks))
}

//3.按照行读取文件
func readFileByLines() {
	filepath := "./log.txt"
	file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	fmt.Println("file size=", size)
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		fmt.Println(line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
	}
}

func testFileRead() {
	//readFile2Memory()
	//readFile2Memory2()
	//readFileByBytes()
	//readFileByBytes2()
	readFileByLines()
}

/*
写文件的几种方式:
1.ioutil.WriteFile
2.os
3.fsync
4.bufio
*/
//1.ioutil.WriteFile
func testFileWriteByWriteFile() {
	content := []byte("测试1\n测试2\n")
	err := ioutil.WriteFile("test.txt", content, 0644)
	if err != nil {
		panic(err)
	}
}

//判断文件是否存在
func checkFileExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

//2.os
func testFileWriteByOs() {
	var writeString = "测试3\n测试4\n"
	var filename = "./test.txt"
	var f *os.File
	var err1 error

	if checkFileExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	defer f.Close()
	n, err1 := io.WriteString(f, writeString) //写入文件(字符串)
	if err1 != nil {
		panic(err1)
	}
	fmt.Printf("写入 %d 个字节\n", n)
}

//3.fsync
func testFileWriteByFSync() {
	var str = "测试1\n测试2\n"
	var filename = "./test.txt"
	var f *os.File
	var err1 error
	if checkFileExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	defer f.Close()
	n, err1 := f.Write([]byte(str)) //写入文件(字节数组)
	fmt.Printf("写入 %d 个字节\n", n)
	n, err1 = f.WriteString(str) //写入文件(字符串)
	if err1 != nil {
		panic(err1)
	}
	fmt.Printf("写入 %d 个字节\n", n)
	f.Sync()
}

//4.bufid
func testFileWriteBufIO() {
	var str = "测试1\n测试2\n"
	var filename = "./test.txt"
	var f *os.File
	var err1 error
	if checkFileExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) //打开文件
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	defer f.Close()
	if err1 != nil {

		panic(err1)
	}
	w := bufio.NewWriter(f) //创建新的 Writer 对象
	n, _ := w.WriteString(str)
	fmt.Printf("写入 %d 个字节\n", n)
	w.Flush()
}

func testFileWrite() {
	//testFileWriteByWriteFile()
	//testFileWriteByOs()
	//testFileWriteByFSync()
	testFileWriteBufIO()
}

func testFileReadWrite() {
	//testFileRead()
	testFileWrite()
}
