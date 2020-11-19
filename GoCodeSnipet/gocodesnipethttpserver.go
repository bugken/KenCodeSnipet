package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("this is root handler."))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("this is index handler."))
}

/*
使用默认DefaultServeMux作为处理器
HandleFunc函数向DefaultServeMux添加处理器。
*/
func startHTTPServer1() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/index", indexHandler)
	err := http.ListenAndServe("127.0.0.1:1986", nil)
	if err != nil {
		fmt.Printf("HTTP server start failed. err:%v", err)
		return
	}
}

/*
使用自定义的ServerMut作为处理器
HandleFunc函数向DefaultServeMux添加处理器。
*/
func startHTTPServer2() {
	serverMut := http.NewServeMux()
	serverMut.HandleFunc("/", rootHandler)
	serverMut.HandleFunc("/index", indexHandler)
	err := http.ListenAndServe("127.0.0.1:1986", serverMut)
	if err != nil {
		fmt.Printf("HTTP server start failed. err:%v", err)
		return
	}
}

/*
使用自定义的ServerMut作为处理器
Handle函数向DefaultServeMux添加处理器。
*/
type testGetHandler struct{}

//ServerHTTP 实现
func (h *testGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if params := r.URL.Query(); len(params) > 0 {
		fmt.Fprintln(w, "name:", params.Get("name"), "hobby:", params.Get("hobby"))
	} else {
		fmt.Fprintln(w, "请求成功")
	}
}

type postHandler struct{}

//Message ..
type Message struct {
	ID  int         `json:"id"`
	Arg interface{} `json:"arg"`
}

func (h *postHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	// r.ParseForm()
	// fmt.Println(r.PostForm) // 打印form数据
	// fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	// 2. 请求类型是application/json时从r.Body读取数据
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
	var msg Message
	json.Unmarshal(b, &msg)
	fmt.Printf("Ummarshal type=%T, id=%v, arg=%v \n", msg, msg.ID, msg.Arg)
	msgBody := msg.Arg
	for index, msg := range msgBody.(map[string]interface{}) {
		fmt.Printf("message infomation is Index=%v, msg=%v\n", index, msg)
	}
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}

func startHTTPServer3() {
	var testPostHandler postHandler
	serverMut := http.NewServeMux()
	serverMut.Handle("/", &testPostHandler)
	var err = http.ListenAndServe("127.0.0.1:1986", serverMut)
	if err != nil {
		fmt.Printf("http server failed, err: %v\n", err)
		return
	}
}

func testHTTPServer() {
	//startHTTPServer1()
	//startHTTPServer2()
	startHTTPServer3()
}
