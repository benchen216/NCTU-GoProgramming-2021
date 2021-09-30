package main

import "fmt"

func Sum() {

}

func main() {
	// Please complete the code to make this program be compiled without error.
	// Notice that you can only add code in this file.
	var n int64 =1
	var sum,i int64
	for ; n!=0;{
		fmt.Scanln(&n)
		if n==0{
			continue
		}
		sum=1
		for i=1; i<=n; i++{
			if i==1{
				fmt.Printf("%d",i)
			}else if i%7==0{
				continue
			}else {
				sum+=i
				fmt.Printf("+%d",i)
			}
		}
		fmt.Printf("=%d\n",sum)
	}
}
