package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)


func lab5(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	op := r.Form["op"][0]
	a, _ := strconv.Atoi(r.Form["num1"][0])
	b, _ := strconv.Atoi(r.Form["num2"][0])
	switch op {
		case "add":
			fmt.Fprintf(w,"%d + %d = %d",a,b,a+b)
		case "sub":
			fmt.Fprintf(w,"%d - %d = %d",a,b,a-b)
		case "mul":
			fmt.Fprintf(w,"%d * %d = %d",a,b,a*b)
		case "div":
			fmt.Fprintf(w,"%d / %d = %d, remainder = %d",a,b,a/b,a%b)
		case "gcd":
			var gcd_temp int =1
			if a>=b{
				for i := 2; i <= b; i++{
					if a%i==0 && b%i==0{
						gcd_temp=i
					}
				}
			}else{
	
				for i := 2; i <= a; i++{
					if a%i==0 && b%i==0{
						gcd_temp=i
					}
				}
			}
			fmt.Fprintf(w,"gcd of %d and %d is %d",a,b,gcd_temp)
		default:
			fmt.Fprintf(w, "hello world!")
	}

}

func main() {
	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	http.HandleFunc("/Lab5", lab5)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}