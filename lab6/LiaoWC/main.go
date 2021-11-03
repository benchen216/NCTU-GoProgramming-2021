package main

import (
	"fmt"
	"math/big"
	"syscall/js"
)

func CheckPrime(this js.Value, i []js.Value) interface{} {
	/* add code here */
	str := js.Global().Get("value").Get("value").String()
	n := new(big.Int)
	n, _ = n.SetString(str, 10)

	if n.ProbablyPrime(0) {
		js.Global().Get("answer").Set("innerHTML", "is prime")
	} else {
		js.Global().Get("answer").Set("innerHTML", "is not prime")
	}

	return nil
}

func registerCallbacks() {
	js.Global().Set("CheckPrime", js.FuncOf(CheckPrime))
}

func main() {
	fmt.Println("Golang main function executed")
	registerCallbacks()

	//need block the main thread forever
	select {}
}
