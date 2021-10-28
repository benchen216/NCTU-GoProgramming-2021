package main

import (
	"fmt"
	"math/big"
	"strconv"
	"syscall/js"
)

func CheckPrime(this js.Value, i []js.Value) interface{} {
	/* add code here */
	str := js.Global().Get("value").Get("value").String()
	n, _ := strconv.Atoi(str)
	if big.NewInt(int64(n)).ProbablyPrime(0) {
		js.Global().Get("answer").Set("innerHTML", "is prime")
	} else {
		js.Global().Get("answer").Set("innerHTML", "is not prime")
	}
	return js.Global().Get("answer")
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
