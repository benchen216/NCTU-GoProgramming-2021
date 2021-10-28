package main

import (
	"fmt"
	"syscall/js"
)

func CheckPrime(this js.Value, i []js.Value) interface{} {
	/* add code here */
	n := js.Global().Get("value").Get("value").Int()
	var i int
	var ans string = "is prime."
	for i=2; i<n/2; i++ {
		if n % i == 0 { ans = "is not prime." }
	}

	js.Global().Get("answer").Set("innerHTML", ans)
	return js.Global().Get("answer")
}

func registerCallbacks() {
	js.Global().Set("CheckPrime", js.FuncOf(CheckPrime))
}

func main() {
	fmt.Println("Golang main function executed")
	registerCallbacks()

	//need block the main thread forever
	select{}
}
