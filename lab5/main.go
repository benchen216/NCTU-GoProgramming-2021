package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func gcd(a, b int) int {
	// b小 a大
	if b == 1 {
		return 1
	}
	return gcd(b, a%b)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world! Lab5\n")
}

func lab5(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	op := r.Form["op"][0]
	num1, _ := strconv.Atoi(r.Form["num1"][0])
	num2, _ := strconv.Atoi(r.Form["num2"][0])
	switch op {
	case "add":
		fmt.Fprintf(w, "%d + %d = %d", num1, num2, num1+num2)
	case "sub":
		fmt.Fprintf(w, "%d - %d = %d", num1, num2, num1-num2)
	case "mul":
		fmt.Fprintf(w, "%d * %d = %d", num1, num2, num1*num2)
	case "div":
		fmt.Fprintf(w, "%d / %d = %d, remainder = %d", num1, num2, num1/num2, num1%num2)
	case "gcd":
		ans := 0
		if num1 > num2 {
			ans = gcd(num1, num2)
		} else {
			ans = gcd(num2, num1)
		}
		fmt.Fprintf(w, "gcd of %d and %d is %d", num1, num2, ans)
		// fmt.Fprintf(w, "%d", ans)
	}
}

func main() {
	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/Lab5", lab5)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
