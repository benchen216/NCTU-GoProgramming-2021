package main

import "fmt"

func Sum(n int64) {
	var a, i int64
	for i = 1; i < n; i++ {
		if i%7 != 0 {
			a += i
			fmt.Printf("%d+", i)
		}
	}
	fmt.Printf("%d=%d\n", n, a+n)
}

func main() {
	// Please complete the code to make this program be compiled without error.
	// Notice that you can only add code in this file.
	var n int64
	for {
		fmt.Scanln(&n)
		if n == 0 {
			break
		} else if n > 0 {
			Sum(n)
		}
	}
}
