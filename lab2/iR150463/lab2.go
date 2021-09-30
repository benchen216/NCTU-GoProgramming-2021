package main

import "fmt"

func Sum(n int64) {
	var sum int64
	for i := int64(1); i <= n; i++ {
		if i%7 != 0 {
			sum += i

			if i > 1 {
				fmt.Printf("+%d", i)
			} else {
				fmt.Printf("%d", i)
			}
		}
	}

	fmt.Printf("=%d\n", sum)
	return
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
