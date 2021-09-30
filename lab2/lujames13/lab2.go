package main

import "fmt"

func Sum(n int64) int64 {
	sum, x := int64(1), int64(1)
	fmt.Printf("%d", x)
	for x=2; x<=n; x++ {
		if x%7 == 0{
			continue
		}
		sum += x
		fmt.Printf("+%d", x)
	}
	fmt.Printf("=%d\n", sum)
	return sum
}

func main() {
	// Please complete the code to make this program be compiled without error.
	// Notice that you can only add code in this file.
	var n int64
	for {
		fmt.Scanln(&n)
		if n>10000{
			continue
		}else if n<=0{
			break
		}else{
			Sum(n)
		}
	}
}
