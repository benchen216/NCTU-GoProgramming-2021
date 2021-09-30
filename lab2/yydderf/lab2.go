package main

import (
    "fmt"
    "log"
)

func main() {
    var num, sum int
    for {
        _, err := fmt.Scanf("%d", &num)
        if err != nil {
            log.Fatal("Failed to scan the input\n")
        }
        if (num == 0) {
            break
        }
        sum = 0
        for i := 1; i <= num; i++ {
            if (i % 7 != 0) {
                if (i != 1) {
                    fmt.Printf("+");
                }
                fmt.Printf("%d", i);
                sum += i
            }
        }
        fmt.Printf("=%d\n", sum);
    }
}
