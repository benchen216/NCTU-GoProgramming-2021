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
	switch op

}

func main() {
	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	http.HandleFunc("/", lab5)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
