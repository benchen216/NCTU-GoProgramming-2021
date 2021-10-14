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
		case "add" :
			fmt.Fprintf(w, "%d + %d = %d\n", a , b , add(a,b) )
		case "sub" :
			fmt.Fprintf(w, "%d - %d = %d\n", a , b , sub(a,b) )
		case "mul" :
			fmt.Fprintf(w, "%d * %d = %d\n", a , b , mul(a,b) )
		case "div" :
			ans1 ,ans2 := div(a,b)
			fmt.Fprintf(w, "%d / %d = %d, remainder = %d\n", a , b , ans1 , ans2 )
		case "gcd" :
			fmt.Fprintf(w, "gcd of %d and %d is %d", a , b , gcd(a,b) )
	}

}

func add(a, b int) int {
	return a+b
}

func sub(a, b int) int {
	return a-b
}

func mul(a, b int) int {
	return a*b
}

func div(a, b int) (int, int) {
	return a/b , a%b
}

func gcd(a,b int) int{
	for b != 0 {
		r := b
		b = a % b
		a = r
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
