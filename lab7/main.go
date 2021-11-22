package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Book struct {
	// write your own struct
}

var bookshelf = []Book{
	// init data
}

func getBooks(c *gin.Context) {
}
func getBook(c *gin.Context) {
}
func addBook(c *gin.Context) {
}
func deleteBook(c *gin.Context) {
}
func updateBook(c *gin.Context) {
}
func main() {
	r := gin.Default()
	r.RedirectFixedPath = true
	r.GET("/bookshelf", getBooks)

	port := "8080"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	r.Run(":" + port)
}
