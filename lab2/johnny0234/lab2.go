package main

import "fmt"

func main() {
	// Please complete the code to make this program be compiled without error.
	// Notice that you can only add code in this file.
	var n int64

	for {
		fmt.Scanln(&n)
		sum := 0
		fmt.Print("1")
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
