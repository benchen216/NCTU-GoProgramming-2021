package main

import (
	"fmt"
	"syscall/js"
	"strconv"
	"math/big"
)

func CheckPrime(this js.Value, i []js.Value) interface{} {
	/* add code here */
	str := js.Global().Get("value").Get("value").String()
	val, _ := strconv.Atoi(str)

	if big.NewInt(int64(val)).ProbablyPrime(0) {
		js.Global().Get("answer").Set("innerHTML","is prime")
	} else {
		js.Global().Get("answer").Set("innerHTML","is not prime")
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
	select{}
}
