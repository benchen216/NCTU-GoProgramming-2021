package main

import (
	"fmt"
)

func main() {
	// Please complete the code to make this program compiled without error.
	// Notice that you can only add code in this file.
	var (
		a, b   int
		action string
		Add    int = a + b
		Sub    int = a - b
		Mul    int = a * b
		Div    int = a / b
	)

	fmt.Println("1) Add")
	fmt.Println("2) Sub")
	fmt.Println("3) Mul")
	fmt.Println("4) Div")

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
