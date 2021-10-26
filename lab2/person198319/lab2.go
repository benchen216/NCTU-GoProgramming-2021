package main

import "fmt"

func Sum(n int) int {
	sum := 0
	//var i int
	for i := 1; i <= n; i++ {
		sum += i
		if i%7 == 0 {
			continue
		}
		fmt.Print(i)

		if i != n {
			fmt.Print("+")
		} else {
			fmt.Print("=")
		}
	}
	fmt.Println(sum)
	return sum
}

func main() {
	// Please complete the code to make this program be compiled without error.
	// Notice that you can only add code in this file.
	var n int
	for {
		fmt.Println("Please input the number n: ")
		fmt.Scanln(&n)
		Sum(n)
	}
}
