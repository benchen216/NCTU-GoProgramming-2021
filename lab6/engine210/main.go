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
	val, _ := strconv.ParseInt(str, 10, 64)
	fmt.Println(val)
	if big.NewInt(val).ProbablyPrime(0) {
		fmt.Println("Is a prime")
		js.Global().Get("answer").Set("innerHTML", "is prime")
	} else {
		fmt.Println("Not a prime")
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
