package main

import (
	"fmt"
	"sync"
	"time"
)

var doorStatus string
var handStatus string

func hand() {
	mu.Lock()  // prevent break when is in

	handStatus = "in"
	time.Sleep(time.Millisecond * 200)
	handStatus = "out"

	mu.Unlock()  // before wg.Done prevent main process die

	wg.Done()

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

	mu.Unlock()
	
	wg.Done()
}

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	for i := 0; i < 50; i++ {
		wg.Add(2)
		go door()
		go hand()
		wg.Wait()
		time.Sleep(time.Millisecond * 200)
	}
}
