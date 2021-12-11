package main

import (
	"fmt"
	"sync"
	"time"
)

var doorStatus string
var handStatus string

func hand() {
    m.Lock()
	handStatus = "in"
	time.Sleep(time.Millisecond * 200)
	handStatus = "out"
	wg.Done()
    m.Unlock()
}

func door() {
    m.Lock()
	doorStatus = "close"
	time.Sleep(time.Millisecond * 200)
	if handStatus == "in" {
		fmt.Println("夾到手了啦！")
	} else {
		fmt.Println("沒夾到喔！")
	}
	doorStatus = "open"
    m.Unlock()
	wg.Done()
}

var wg sync.WaitGroup
var m sync.Mutex

func main() {
	for i := 0; i < 50; i++ {
		wg.Add(2)
		go door()
		go hand()
		wg.Wait()
		time.Sleep(time.Millisecond * 200)
	}
}
