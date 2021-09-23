package main

import (
	"fmt"
)

func Add(x, y int64) int64 {
	return x + y
}
func Sub(x, y int64) int64 {
	return x - y
}
func Mul(x, y int64) int64 {
	return x * y
}
func Div(x, y int64) int64 {
	return x / y
}
func main() {
	// Please complete the code to make this program compiled without error.
	// Notice that you can only add code in this file.
	var action, a, b int64
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
