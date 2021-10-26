package main

import (
	"fmt"
	"math/big"
	"syscall/js"
)

func CheckPrime(this js.Value, i []js.Value) interface{} {
	/* add code here */
	source := js.Global().Get(i[0].String()).Get("value").String()
	target := js.Global().Get(i[1].String())

	bi := big.NewInt(0)
	n, _ := bi.SetString(source, 10)
	if bi.ProbablyPrime(0) {
		fmt.Println(n, "is prime")
		target.Set("innerHTML", "is prime")
		return true
	} else {
		fmt.Println(n, "is not prime")
		target.Set("innerHTML", "is not prime")
		return false
	}
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
