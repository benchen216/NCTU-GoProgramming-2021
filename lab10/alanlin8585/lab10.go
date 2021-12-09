package main

import (
	"fmt"
	"sync"
	"time"
)

var doorStatus string
var handStatus string
var can_door sync.Mutex
var can_hand sync.Mutex

func hand() {
    can_hand.Lock()
	handStatus = "in"
	time.Sleep(time.Millisecond * 200)
	handStatus = "out"
	can_door.Unlock()
	wg.Done()
}

func door() {
    can_door.Lock()
	doorStatus = "close"
	time.Sleep(time.Millisecond * 200)
	if handStatus == "in" {
		fmt.Println("夾到手了啦！")
	} else {
		fmt.Println("沒夾到喔！")
	}
	doorStatus = "open"
	can_hand.Unlock()
	wg.Done()
}

var wg sync.WaitGroup

func main() {
    can_door.Lock()
	for i := 0; i < 50; i++ {
		wg.Add(2)
		go door()
		go hand()
		wg.Wait()
		time.Sleep(time.Millisecond * 200)
	}
}
