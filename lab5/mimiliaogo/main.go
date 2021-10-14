package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func gcd(a int, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}
func lab5(w http.ResponseWriter, r *http.Request) {
	// https://new-app-name.herokuapp.com/Lab5?op=gcd&num1=7&num2=3
	r.ParseForm()
	op := r.Form["op"][0]
	a, _ := strconv.Atoi(r.Form["num1"][0])
	b, _ := strconv.Atoi(r.Form["num2"][0])


	switch op {
	case "add":
		fmt.Fprintf(w, strconv.Itoa(a)+" + "+strconv.Itoa(b)+" = "+strconv.Itoa(a+b))
	case "sub":
		fmt.Fprintf(w, strconv.Itoa(a)+" - "+strconv.Itoa(b)+" = "+strconv.Itoa(a-b))
	case "mul":
		fmt.Fprintf(w, strconv.Itoa(a)+" * "+strconv.Itoa(b)+" = "+strconv.Itoa(a*b))
	case "div":
		fmt.Fprintf(w, strconv.Itoa(a)+" / "+strconv.Itoa(b)+" = "+strconv.Itoa(a/b)+", remainder = "+strconv.Itoa(a%b))
	case "gcd":
		result := gcd(a, b)
		fmt.Fprintf(w, "gcd of " + strconv.Itoa(a) + " and " + strconv.Itoa(b) + " is " + strconv.Itoa(result))

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
