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
		fmt.Fprintf(w, "%d + %d = %d", a, b, a+b)
	case "sub":
		fmt.Fprintf(w, "%d - %d = %d", a, b, a-b)
	case "mul":
		fmt.Fprintf(w, "%d * %d = %d", a, b, a*b)
	case "div":
		fmt.Fprintf(w, "%d / %d = %d, remainder = %d", a, b, a/b, a%b)
	case "gcd":
		origin_a:=a
		origin_b:=b
		for ;a!=0&&b!=0; {
			if (a >= b) {
				a = a % b;
			} else if b > a {
				b = b % a;
			}
		}
		if a >= b {
			fmt.Fprintf(w,"gcd of %d and %d is %d", origin_a,origin_b,a)
		}		else {
			fmt.Fprintf(w,"gcd of %d and %d is %d", origin_a,origin_b,b)
		}
	}

}

func main() {
	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	http.HandleFunc("/", lab5)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
