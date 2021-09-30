package main

import (
	"fmt"
	"strconv"
)

func Sum() {

}

func main() {
	// Please complete the code to make this program be compiled without error.
	// Notice that you can only add code in this file.
	var n int
	for {
		fmt.Scanln(&n)
		if n == 0 {
			break
		}
		var ans string = "1"
		var sum int = 1
		for i := 2; i <= n; i++ {
			if i%7 != 0 {
				ans += "+" + strconv.Itoa(i)
				sum += i
			}
		}
		ans += "=" + strconv.Itoa(sum)

		fmt.Println(ans)
	}
}
