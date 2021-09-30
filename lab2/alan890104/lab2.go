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
			for i = 1; i <= n; i++ {
				if i%7 != 0 {
					if i != n {
						fmt.Printf("%d+", i)
					} else {
						fmt.Printf("%d", i)
					}
					s += i
				}
			}
			fmt.Printf("=%d\n", s)
		}
	}
}
