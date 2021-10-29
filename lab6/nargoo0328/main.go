package main

import (
	"fmt"
	"syscall/js"
	"strconv"
)

func CheckPrime(this js.Value, i []js.Value) interface{} {
	/* add code here */
	str := js.Global().Get("value").Get("value").String()
	val, _ := strconv.Atoi(str)
	for j:=2; j*j<=val;j++{
		if val%j==0{
			js.Global().Get("answer").Set("innerHTML","is not prime")
			return nil
		}
	}
	js.Global().Get("answer").Set("innerHTML","is prime")
	return nil
}

func registerCallbacks() {
	js.Global().Set("CheckPrime", js.FuncOf(CheckPrime))
}

func main() {
	fmt.Println("Golang main function executed")
	registerCallbacks()
	select{}
	//need block the main thread forever
}
