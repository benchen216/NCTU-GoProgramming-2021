package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "[::1]:59658 表示: 靠北"
	fmt.Println(strings.Contains(a, "靠北"))

}
