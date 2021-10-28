package main

import (
	"fmt"
	"syscall/js"
	"math/big"
)

func CheckPrime(this js.Value, i []js.Value) interface{} {
	/* add code here */
	if js.Global().Get("value").Get("value").String() == "" {
		js.Global().Get("answer").Set("innerHTML", "")
		return "no input"
	}
	
	var v big.Int
	_, ok := v.SetString(js.Global().Get("value").Get("value").String(), 10)
	if !ok {
		js.Global().Get("answer").Set("innerHTML", "invalid input")
		return "not an integer"
	} 
	
	if v.ProbablyPrime(20) {
		js.Global().Get("answer").Set("innerHTML", "is prime")		
	} else {
		js.Global().Get("answer").Set("innerHTML", "is not prime")				
	}
	return "integer"
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
