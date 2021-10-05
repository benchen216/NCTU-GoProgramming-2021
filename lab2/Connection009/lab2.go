package main

import "fmt"

func Sum(n int64) {
	var i, sum int64
	fmt.Print(1)
	sum = 1
	for i=2; i<=n; i++ {
		if i % 7 == 0 { continue }
		fmt.Print(" + ", i)
		sum += i
	}
	fmt.Println(" =", sum)
}

func main() {
	// Please complete the code to make this program be compiled without error.
	// Notice that you can only add code in this file.
	var n int64
	for {
		fmt.Scanln(&n)
		if n > 0 && n <= 10000 { Sum(n) }
	}
}
