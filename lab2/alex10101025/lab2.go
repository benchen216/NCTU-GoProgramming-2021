package main

import (
	"fmt"
)

func Sum(n int64) int64 {
	var total, i int64
	for i = 1; i <= n; i++ {
		if i%7 == 0 {
			continue
		}
		if i > 1 {
			fmt.Printf("+%d", i)
		} else {
			fmt.Printf("%d", i)
		}
		total += i
	}
	return total
}

func main() {
	// Please complete the code to make this program be compiled without error.
	// Notice that you can only add code in this file.
	var n int64
	for {
		fmt.Scanln(&n)
		if n > 0 {
			total := Sum(n)
			fmt.Printf("=%d\n", total)
		} else {
			break
		}
	}
}
