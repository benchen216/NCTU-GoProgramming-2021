package main

import (
	"fmt"
	"syscall/js"
)

func CheckPrime(this js.Value, i []js.Value) interface{} {
	/* add code here */
	n := this.Int()
	var i, j int
	var ans string = "is not prime."
	for i=2; i<n; i++ {
		for j=2; j<i; j++ {
			if i % j == 0 { ans = "is prime." }
		}
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
}
