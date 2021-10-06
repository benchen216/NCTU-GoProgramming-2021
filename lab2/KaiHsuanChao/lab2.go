package main

import "fmt"

func Sum(n int64) {
	sum := int64(0)
	if n%7 != 0 {
		for a := int64(1); a <= n; a++ {
			if a%7 == 0 {
				continue
			}
			if a != n {
				fmt.Print(a, "+")
			}
			sum = sum + a
		}
		fmt.Println(n, "=", sum)
	}
	if n%7 == 0 {
		for a := int64(1); a <= n; a++ {
			if a%7 == 0 {
				continue
			}
			if a != n && (a+1) != n {
				fmt.Print(a, "+")
			} else {
				fmt.Print(a)
			}
			sum = sum + a
		}
		fmt.Println("=", sum)
	}
}

func main() {
	// Please complete the code to make this program be compiled without error.
	// Notice that you can only add code in this file.
	var n int64
	for {
		fmt.Scanln(&n)

		if n == 0 {
			break
		}
		if n != 0 {
			Sum(n)
		}

	}
}
