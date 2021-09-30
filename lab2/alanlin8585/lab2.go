package main

import "fmt"

func Sum(n int64) {
    var i, ans int64
    ans = 0
    for i = 1; i <= n; i++ {
        if i % 7 == 0 {
            continue
        }
        if i != 1 {
            fmt.Print("+")
        }
        fmt.Print(i)
        ans += i
    }
    fmt.Print("=")
    fmt.Println(ans)
}

func main() {
	// Please complete the code to make this program be compiled without error.
	// Notice that you can only add code in this file.
	var n int64
	for {
		fmt.Scanln(&n)
        if n == 0 {
            break;
        }
        Sum(n)
	}
}
