package main

import (
	"fmt"
	"math/big"
	"syscall/js"
)

func CheckPrime(this js.Value, i []js.Value) interface{} {
	/* add code here */
	str := js.Global().Get("value").Get("value").String()

	z := new(big.Int)
	fmt.Sscan(str, z)
	if z.ProbablyPrime(20) {
		// fmt.Println(z, "is probably prime")
		js.Global().Get("answer").Set("innerHTML", "is prime")
	} else {
		// fmt.Println(z, "is not prime")
		js.Global().Get("answer").Set("innerHTML", "is not prime")
	}

	return 0
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
