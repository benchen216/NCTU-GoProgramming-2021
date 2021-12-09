package main

import (
	"fmt"
	"sync"
	"time"
)

var doorStatus string
var handStatus string
var mu sync.Mutex

func hand() {
	mu.Lock()
	handStatus = "in"
	time.Sleep(time.Millisecond * 200)
	handStatus = "out"
	wg.Done()
	mu.Unlock()
}

func door() {
	mu.Lock()
	doorStatus = "close"
	time.Sleep(time.Millisecond * 200)
	if handStatus == "in" {
		fmt.Println("夾到手了啦！")
	} else {
		fmt.Println("沒夾到喔！")
	}
	doorStatus = "open"
	wg.Done()
	mu.Unlock()
}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 50; i++ {
		wg.Add(2)
		go door()
		go hand()
		wg.Wait()
		time.Sleep(time.Millisecond * 200)
	}
}
