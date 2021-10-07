package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func cal(ctx *gin.Context) {
	op := ctx.Param("op")
	a, _ := strconv.Atoi(ctx.Param("num1"))
	b, _ := strconv.Atoi(ctx.Param("num2"))
	switch op {
	case "add":
		ctx.String(http.StatusOK,
			fmt.Sprintf("%d + %d = %d", a, b, add(a, b)))
	case "sub":
		ctx.String(http.StatusOK,
			fmt.Sprintf("%d - %d = %d", a, b, sub(a, b)))
	case "mul":
		ctx.String(http.StatusOK,
			fmt.Sprintf("%d * %d = %d", a, b, mul(a, b)))
	case "div":
		q, r := div(a, b)
		ctx.String(http.StatusOK,
			fmt.Sprintf("%d / %d = %d, remainder = %d", a, b, q, r))
	default:
		ctx.String(http.StatusOK, "Bad parameter")
	}
}

func add(a, b int) int {
	return a + b
}

func div(a, b int) (int, int) {
	return a / b, a % b
}

func mul(a, b int) int {
	return a * b
}

func sub(a, b int) int {
	return a - b
}

func main() {
	port := "12345"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	server := gin.Default()
	server.StaticFile("/favicon.ico", "./favicon.ico")
	server.NoRoute(func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world!")
	})
	server.GET("/:op/:num1/:num2", cal)
	if server.Run(fmt.Sprintf(":%s", port)) != nil {
		log.Fatal("Server cannot start on port", port)
	}
}
