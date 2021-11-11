package main

import (
	"fmt"
	"math/big"
	"strconv"
	"syscall/js"
)

func CheckPrime(this js.Value, i []js.Value) interface{} {
	str := js.Global().Get("value").Get("value").String()
	number, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	if big.NewInt(number).ProbablyPrime(0) {
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

	select {} // block the main thread forever
}
