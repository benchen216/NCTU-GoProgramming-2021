package main

import (
	"fmt"
	"syscall/js"
	"math/big"
	"strconv"
)

func CheckPrime(this js.Value, i []js.Value) interface{} {
	/* add code here */
	n := js.Global().Get("value").Get("value").String()
	nint,_ := strconv.ParseInt(n,10,64)
	if big.NewInt(nint).ProbablyPrime(0) {
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
	
	select {}
	//need block the main thread forever
}
