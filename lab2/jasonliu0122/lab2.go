package main

import "fmt"

func Sum(n int64)int64 {
	var total int64 = 0 
	var i int64 = 1
	for ; i <= n; i++{
		if i == 1{
			total += i
			fmt.Print(i)
		}else if i % 7 != 0{
			total += i
			fmt.Print("+",i)
		}
 	}
	fmt.Print("=")
	return total
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
		fmt.Println(Sum(n))

	}
}
