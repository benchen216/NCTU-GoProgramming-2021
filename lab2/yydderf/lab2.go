package main

import (
    "fmt"
    "log"
)

func main() {
	var n, sum, i int64
	for {
        _, err := fmt.Scanln(&n)
        if err != nil {
            log.Fatal("Failed to get n\n")
        }
        if (n < 1) {
            break
        }
        sum = 0
        for i = 1; i <= n; i++ {
            if (i % 7 == 0) {
                continue
            }
            if (i != 1) {
                fmt.Print("+")
            }
            sum += i
            fmt.Print(i)
        }
        fmt.Printf("=%d\n", sum)
	}
}
