package main

import "fmt"

func Sum(n int64) {
	var i, sum int64
	for i = 1; i < n; i++ {
		if i % 7 != 0 {
			fmt.Printf("%d", i)
			sum += i
			i += 1
			break
		}
	}
	for ; i <= n; i++ {
		if i % 7 != 0 {
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
		if n > 0 {
			Sum(n)
		} else {
			break
		}
	}
}
