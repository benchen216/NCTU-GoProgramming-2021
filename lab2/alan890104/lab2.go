package main

import "fmt"

func Sum() {

}

func main() {
	// Please complete the code to make this program be compiled without error.
	// Notice that you can only add code in this file.
	var n, i, s int64
	for {
		fmt.Scanln(&n)
		if n == 0 {
			break
		} else {
			s = 0
			if n%7 == 0 {
				n = n - 1
			}
			for i = 1; i <= n; i++ {
				if i%7 != 0 {
					fmt.Printf("%d", i)
					s += i
					if i != n {
						fmt.Print("+")
					}
				}
			}
			fmt.Printf("=%d\n", s)
		}
	}
}
