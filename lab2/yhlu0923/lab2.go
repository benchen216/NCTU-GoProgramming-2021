package main

import "fmt"

func Sum(n int64) {
	arr := []int64{}
	tmp := int64(7)
	for i := int64(1); tmp < 10000; i++ {
		arr = append(arr, tmp)
		tmp = 7 * (i + 1)
	}

	var str string = "1"
	var total int64 = 1
	for i := int64(2); i <= n; i++ {

		var flag bool = false
		for j := 0; j < len(arr) && int64(arr[j]) <= i; j++ {
			if int64(arr[j]) == i {
				flag = true
				break
			}
		}
		if flag == false {
			str += "+" + fmt.Sprint(i)
			total += i
		}
	}
	fmt.Printf(str + "=" + fmt.Sprint(total))
	fmt.Println()
}

func main() {
	// Please complete the code to make this program be compiled without error.
	// Notice that you can only add code in this file.
	var n int64
	for {
		fmt.Scanln(&n)
		if n == 0 {
			break
		} else {
			// PrintSum(n)
			Sum(n)
		}
		// fmt.Print(n)
		// fmt.Print((&n))
		// fmt.Printf("hihi")
	}
}
