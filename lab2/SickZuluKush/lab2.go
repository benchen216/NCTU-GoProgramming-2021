package main

import (
	"fmt"
)

func main() {
	// Please complete the code to make this program be compiled without error.
	// Notice that you can only add code in this file.
	
	var n int64
	
	for {
		fmt.Scanln(&n)
		if n == 0 {
			break
		}
		
		var sum int64 = 0
		var first = true
		for i := int64(1); i <= n; i++ {
			if i % 7 != 0 {
				if !first {
					fmt.Print("+")
				} else {
					first = false				
				}
				
				sum += i
				fmt.Printf("%v", i)
			}
		}
		
		fmt.Printf("=%v\n", sum)
	}
}