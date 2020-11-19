package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

/*
客户端常用方法:
resp, err := http.Get("http://example.com/")
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
resp, err := http.PostForm("http://example.com/form", url.Values{"key": {"Value"}, "id": {"123"}})
*/

//简单的get请求
func testHTTPClientGet() {
	resp, err := http.Get("http://127.0.0.1:1986")
	if err != nil {
		fmt.Printf("http.Get error, err:%v\n", err)
		return
	}
	defer resp.Body.Close() //必须关闭

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read resp error, err:%v\n", err)
		return
	}
	fmt.Println(string(body))
}

//带参数的Get请求
func testHTTPClientGetWithParams() {
	//设置请求参数
	params := url.Values{}
	params.Set("name", "ken")
	params.Set("hobby", "jogging")

	//设置请求url
	rawURL := "http://127.0.0.1:1986"
	reqURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		fmt.Printf("ParseRequestURI error, err:%v\n", err)
		return
	}

	//整合请求URL和参数, Encode方法将请求参数编码为url编码格式("bar=baz&foo=quux")，编码时会以键进行排序。
	reqURL.RawQuery = params.Encode()

	//发送HTTP请求, reqURL.String() String将URL重构为一个合法URL字符串。
	resp, err := http.Get(reqURL.String())
	if err != nil {
		fmt.Printf("http.Get error, err:%v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read resp error, err:%v\n", err)
		return
	}
	fmt.Println(string(body))

}

//post请求
func testHTTPClientPost() {
	url := "http://127.0.0.1:1986/post"
	//表单数据
	// contentType := "application/x-www-form-urlencoded"
	// data := "name=prince&age=18"
	//json数据
	contentType := "application/json"
	data := `{"id":900005,"arg":{"userid":234, "message":"this is a message"}}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read all from resp body error, err:%v\n", err)
		return
	}
	fmt.Println(string(body))
}

func testHTTPClient() {
	// testHTTPClientGet()
	// testHTTPClientGetWithParams()
	testHTTPClientPost()
}
