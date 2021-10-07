package main

import "fmt"

func main() {
	// Please complete the code to make this program be compiled without error.
	// Notice that you can only add code in this file.
	var n int64

	for {
		fmt.Scanln(&n)
		if n == 0 {
			break
		}
		fmt.Print("1")
		sum := 0
		for x := 2; int64(x) <= n; x++ {
			if x%7 == 0 {
				continue
			}
			sum += x
			fmt.Print("+", x)
		}
		fmt.Print("=", sum+1)
	}

}
