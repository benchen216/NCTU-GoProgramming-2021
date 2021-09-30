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
	var n int64
	var total = 0
	var numbers_add_str string
	for {
		fmt.Scanln(&n)
		if n == 0 {
			break;
		}
		for index:=1; index <= int(n); index++ {
			if index % 7 != 0 {
				total += index
				numbers_add_str += strconv.Itoa(index) + "+"
			}
		}
		// strings.TrimSuffix(numbers_add_str, "+")
		numbers_add_str = numbers_add_str[:len(numbers_add_str)-1]
		numbers_add_str += "=" + strconv.Itoa(total)
		fmt.Println(numbers_add_str)
		numbers_add_str = ""
		total = 0
	}
}
