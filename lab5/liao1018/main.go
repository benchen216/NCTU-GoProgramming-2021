package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func findGcd(a int,b int) int {

	for b != 0 {
		t := b
		b = a%b
		a = t	
	}

	return a
}

func lab5(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	op := r.Form["op"][0]
	num1, _ := strconv.Atoi(r.Form["num1"][0])
	num2, _ := strconv.Atoi(r.Form["num2"][0])

	switch op {
	case "add":
		result := num1+num2
		output := strconv.Itoa(num1) + " + " + strconv.Itoa(num2) + " = " + strconv.Itoa(result)
		fmt.Fprintf(w,output)

	case "sub":
		result := num1-num2
		output := strconv.Itoa(num1) + " - " + strconv.Itoa(num2) + " = " + strconv.Itoa(result)
		fmt.Fprintf(w,output)

	case "mul":
		result := num1*num2
		output := strconv.Itoa(num1) + " * " + strconv.Itoa(num2) + " = " + strconv.Itoa(result)
		fmt.Fprintf(w,output)

	case "div":
		result := num1/num2
		output := strconv.Itoa(num1) + " / " + strconv.Itoa(num2) + " = " + strconv.Itoa(result)
		output += ", remainder = " + strconv.Itoa(num1%num2)
		fmt.Fprintf(w,output)

	case "gcd":
		result := findGcd(num1,num2)
		output := "gcd of " + strconv.Itoa(num1) + " and " + strconv.Itoa(num2) + " is " + strconv.Itoa(result)
		fmt.Fprintf(w,output)

	default:
		fmt.Fprintf(w,"Hello World")

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