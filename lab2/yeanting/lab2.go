package main

import "fmt"
import "strconv"

func Sum(n int64) string{
	var sum int64 = 0
		var s,s2 string = "", ""
		for i:=int64(1); i<=n; i++{
			if i%7 == 0{
				continue
			}
			sum += i
			s = s + strconv.FormatInt(i, 10) + string('+')
		}
		if len(s) != 0{
			s2 = s[:(len(s)-1)] // remove last plus symbol
		}
		s2 = s2 + string('=') + strconv.FormatInt(sum, 10)
	return s2
}

func main() {
	// Please complete the code to make this program be compiled without error.
	// Notice that you can only add code in this file.
	var n int64
	for {
		fmt.Scanln(&n)
		if n == 0{
			break
		}
		s := Sum(n)
		
		fmt.Printf("%s\n",s)
	}
}
