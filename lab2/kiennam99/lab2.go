package main

import "fmt"

func Sum(n int64) {
	var i,total int64
	total = 0
	for i = 1;i<n-1;i++{
		if i % 7 == 0 {
			continue
		}
		fmt.Print(i)
		fmt.Print("+")
		total += i
	}
	if (n-1) % 7 == 0 {
		total+=n	
		fmt.Print(n,"=")
		fmt.Println(total)
		
	} else {
		total+=n-1
		if n % 7 == 0{
			fmt.Print(n-1,"=")
			fmt.Println(total)
		} else {
			total+=n
			fmt.Print(n-1,"+",n,"=")
			fmt.Println(total)
		}
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
		Sum(n)
	}
}
