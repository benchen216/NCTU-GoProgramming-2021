package main

import "fmt"

func Sum(n int64) {
	var sum = int64(0)

	for a := int64(1); a <= n; a++{
		if a%7 == 0 {
			continue
		} else {
			if(a == 1){
				fmt.Print(a)
			} else {
				fmt.Print("+",a)
			}
			sum = sum + a
		}
	} 
	fmt.Print("=")
	fmt.Println(sum)
}

func main() {
	// Please complete the code to make this program be compiled without error.
	// Notice that you can only add code in this file.
	var n int64
	for {
		fmt.Println("Enter a int.")
		fmt.Scanln(&n)

		if n==0{
			break
		}
		if n!=0 {
			Sum(n)
		}

	}
}