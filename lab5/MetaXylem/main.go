package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func hacked(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your computer is hacked\n")
}

func hello(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 1000; i++ {
		fmt.Fprintf(w, "HIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHI")
		fmt.Fprintf(w, "HIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHIHI\n")
	}
}

func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	} else {
		return a
	}
}

func lab5(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	op := r.Form["op"][0]
	num1, _ := strconv.Atoi(r.Form["num1"][0])
	num2, _ := strconv.Atoi(r.Form["num2"][0])
	switch op {
	case "add":
		fmt.Fprintf(w, strconv.Itoa(num1)+" + "+strconv.Itoa(num2)+" = "+strconv.Itoa(num1+num2))
	case "sub":
		fmt.Fprintf(w, strconv.Itoa(num1)+" - "+strconv.Itoa(num2)+" = "+strconv.Itoa(num1-num2))
	case "mul":
		fmt.Fprintf(w, strconv.Itoa(num1)+" * "+strconv.Itoa(num2)+" = "+strconv.Itoa(num1*num2))
	case "div":
		fmt.Fprintf(w, strconv.Itoa(num1)+" / "+strconv.Itoa(num2)+" = "+strconv.Itoa(num1/num2)+", remainder = "+strconv.Itoa(num1%num2))
	case "gcd":
		fmt.Fprintf(w, "("+strconv.Itoa(num1)+", "+strconv.Itoa(num2)+") = "+strconv.Itoa(gcd(num1, num2)))
	}
}

func main() {
	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	http.HandleFunc("/", hello)
	http.HandleFunc("/Hacked", hacked)
	http.HandleFunc("/Lab5", lab5)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
