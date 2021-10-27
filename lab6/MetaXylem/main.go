package main

import (
	"fmt"
	"math/big"
	"strconv"
	"syscall/js"
)

func CheckPrime(this js.Value, i []js.Value) interface{} {
	num, _ := strconv.Atoi(js.Global().Get("value").Get("value").String())
	var str string
	if big.NewInt(int64(num)).ProbablyPrime(0) {
		str = "is prime."
	} else{
		str = "is not prime."
	}
	js.Global().Get("answer").Set("innerHTML", str)
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
