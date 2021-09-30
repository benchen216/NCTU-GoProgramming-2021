package main

import "fmt"

func Sum(n int64) {
	var tmp = 0
	for i := 1; i <= int(n); i++ {
		if int64(i)%7 != 0 {
			if i != 1 {
				fmt.Printf("+%d", i)
			} else {
				fmt.Printf("%d", i)
			}
			tmp += i
		}
	}
	fmt.Printf("=%d\n", tmp)
}

func main() {
	// Please complete the code to make this program be compiled without error.
	// Notice that you can only add code in this file.
	var n int64
	for {
		fmt.Scanln(&n)
		if n == 0 {
			return
		}
		Sum(n)
	}
}
