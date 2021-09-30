package main

import "fmt"

func Sum(n int64) {
	if n == 1 {
		fmt.Printf("%d\n", 1)
		return
	}
	var a, i int64
	a = 1
	fmt.Printf("%d", 1)
	for i = 2; i < n; i++ {
		if i%7 != 0 {
			a += i
			fmt.Printf("+%d", i)
		}
	}
	if n%7 != 0 {
		fmt.Printf("+%d=%d\n", n, a+n)
	} else {
		fmt.Printf("=%d\n", a)
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
		} else if n > 0 {
			Sum(n)
		}
	}
}
