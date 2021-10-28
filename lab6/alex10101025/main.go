package main

import (
	"fmt"
	"math/big"
	"syscall/js"
)

func CheckPrime(this js.Value, i []js.Value) interface{} {
	/* add code here */
	z := new(big.Int)
	str := js.Global().Get("value").Get("value").String()
	fmt.Sscan(str, z)
	if z.ProbablyPrime(20) {
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
	select {}
	//need block the main thread forever
}
