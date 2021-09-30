package main

import "fmt"

func Sum(n int64) {
	var str string = "1"
	var total int64 = 1
	for i := int64(2); i <= n; i++ {
		str += "+" + fmt.Sprint(i)
		total += i
	}
	fmt.Printf(str + "=" + fmt.Sprint(total))
	fmt.Println()
}

func main() {
	// Please complete the code to make this program be compiled without error.
	// Notice that you can only add code in this file.
	var n int64
	for {
		fmt.Scanln(&n)
		if n == 0 {
			break
		} else {
			// PrintSum(n)
			Sum(n)
		}
		// fmt.Print(n)
		// fmt.Print((&n))
		// fmt.Printf("hihi")
	}
}
