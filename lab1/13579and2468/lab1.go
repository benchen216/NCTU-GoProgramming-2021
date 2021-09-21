package main

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}
func Sub(a, b int) int {
	return a - b
}
func Mul(a, b int) int {
	return a * b
}
func Div(a, b int) int {
	return a / b
}
func main() {
	// Please complete the code to make this program compiled without error.
	// Notice that you can only add code in this file.

	fmt.Println("1) Add")
	fmt.Println("2) Sub")
	fmt.Println("3) Mul")
	fmt.Println("4) Div")
	var action int
	var a, b int
	fmt.Println("Please input your action: ")
	fmt.Scanln(&action)
	fmt.Println("Please input two numbers: ")
	fmt.Scanln(&a, &b)
	switch action {
	case 1:
		fmt.Println(Add(a, b))
	case 2:
		fmt.Println(Sub(a, b))
	case 3:
		fmt.Println(Mul(a, b))
	case 4:
		fmt.Println(Div(a, b))
	default:
		fmt.Println("Wrong input!!")
	}
}
