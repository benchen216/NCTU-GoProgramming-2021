package main

import "fmt"

func Sum(n int64) int64 {
	var sum, i int64
	sum = 1
	fmt.Print(1)
	for i = 2; i <= n; i++ {
		if i%7 != 0 {
			fmt.Print("+", i)
			sum += i
		}
	}
	fmt.Print("=")
	return sum
}

func main() {
	// Please complete the code to make this program be compiled without error.
	// Notice that you can only add code in this file.
	var n int64
	for {
		fmt.Scanln(&n)
		if n > 0 {
			fmt.Println(Sum(n))
		} else {
			break
		}
	}
}
