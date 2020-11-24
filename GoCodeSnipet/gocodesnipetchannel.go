package main

import (
	"fmt"
	"time"
)

const (
	MAX_REQUEST_NUM = 10
	CMD_USER_POS    = 1
)

var (
	save chan bool
	quit chan bool
	req  chan *Request
)

type Request struct {
	CmdId uint16
	Data  interface{}
}

type UserPos struct {
	X int16
	Y int16
}

func testChannelExample() {
	req = make(chan *Request, MAX_REQUEST_NUM)
	save = make(chan bool)
	quit = make(chan bool)

	newReq := Request{
		CmdId: CMD_USER_POS,
		Data: UserPos{
			X: 10,
			Y: 20,
		},
	}

	go handler()
	req <- &newReq
	time.Sleep(2000 * time.Microsecond)

	save <- true
	//close(req)
	<-quit
}

func handler() {
	for {
		select {
		case <-save:
			saveGame()
		case r, ok := <-req:
			if ok {
				onReq(r)
			} else {
				fmt.Println("req chan closed.")
			}
		}
	}
}

func saveGame() {
	fmt.Println("......saveGame......")
	quit <- true
}

func onReq(r *Request) {
	pos := r.Data.(UserPos)
	fmt.Println(r.CmdId, pos)
}
