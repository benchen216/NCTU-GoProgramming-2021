package main

import (
	"fmt"
)

func Sum(a int64) {
	fmt.Printf("1")
	var i, sum int64
	sum = 1
	for i = 2; i < a+1; i++ {
		if i%7 != 0 {
			fmt.Printf("+%d", i)
			sum += i
		}
	}
	fmt.Printf("=%d\n", sum)
}

func main() {
	// Please complete the code to make this program be compiled without error.
	// Notice that you can only add code in this file.
	var n int64
	for {
		fmt.Scanln(&n)
		if n < 1 {
			break
		}
		Sum(n)
	}
}
