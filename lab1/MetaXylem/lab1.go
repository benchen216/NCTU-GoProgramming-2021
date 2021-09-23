package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}
func minus(a, b int) int {
	return a - b
}
func times(a, b int) int {
	return a * b
}
func divide(a, b int) int {
	return a / b
}
func main() {
	cmd := 0
	var a, b int
	for _, err := fmt.Scanln(&cmd, &a, &b); err == nil; _, err = fmt.Scanln(&cmd, &a, &b) {
		if cmd == 1 {
			cmd = add(a, b)
		} else if cmd == 2 {
			cmd = minus(a, b)
		} else if cmd == 3 {
			cmd = times(a, b)
		} else {
			cmd = divide(a, b)
		}
		fmt.Println(cmd)
	}
}
