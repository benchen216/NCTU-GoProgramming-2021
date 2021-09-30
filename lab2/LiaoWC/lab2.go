package main

import (
	"fmt"
	"os"
)

func Sum(n int64) {
	var i int64
	if n <= 0 {
		os.Exit(0)
	}
	var cnt int64
	hasFirstTime := false
	for i = 1; i <= n; i++ {
		if i%7 == 0 {
			continue
		}
		cnt += i
		if !hasFirstTime {
			fmt.Print(i)
		} else {
			fmt.Print("+", i)
		}
		hasFirstTime = true
	}
	//fmt.Println(" = ", cnt)
	fmt.Print("=", cnt,"\n")
}

func main() {
	// Please complete the code to make this program be compiled without error.
	// Notice that you can only add code in this file.
	var n int64
	for {
		_, _ = fmt.Scanln(&n)
		Sum(n)
	}

}
