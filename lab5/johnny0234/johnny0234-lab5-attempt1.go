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
	num1, _ := strconv.Atoi(r.Form["num1"][0])
	num2, _ := strconv.Atoi(r.Form["num2"][0])

	if op == "gcd" {
		fmt.Fprintf(w, "gcd of %d and %d is %d", num1, num2, gcd(num1, num2))
	} else {
		fmt.Fprintf(w, "error 404")
	}

}
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
func main() {
	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	http.HandleFunc("/Lab5", lab5)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
