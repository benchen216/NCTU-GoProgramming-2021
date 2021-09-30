package main

import "fmt"

func Sum(n int64) int64 {
	sum, x := int64(0), int64(1)
	for x=1; x<n; x++ {
		if x%7 == 0{
			continue
		}
		sum += x
		fmt.Printf("%d + ", x)
	}
	sum += x
	fmt.Printf("%d = %d\n", x, sum)
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
