package main

import (
	"fmt"
	"log"
)

func main() {
	var n, total, i int64
	for {
		_, err := fmt.Scanln(&n)
		if err != nil {
			log.Fatal("Failed\n")
		}
		if n < 1 {
			break
		}
		total = 0
		for i = 1; i <= n; i++ {
			if i%7 == 0 {
				continue
			}
			if i != 1 {
				fmt.Print("+")
			}
			total += i
			fmt.Print(i)
		}
		fmt.Printf("=%d\n", total)
	}
}
