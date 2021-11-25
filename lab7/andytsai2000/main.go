package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Book struct {
	// write your own struct
	id     string `json:"id"`
	title  string `json:"title"`
	author string `json:"author"`
}

var bookshelf = []Book{
	// init data
	{
		id:     "0001",
		title:  "aa bird",
		author: "cat",
	},
	{
		id:     "0002",
		title:  "aa bird 2",
		author: "cat 2",
	},
	{
		id:     "0003",
		title:  "aa bird 3",
		author: "cat 3",
	},
	{
		id:     "0004",
		title:  "aa bird 4",
		author: "cat 4",
	},
	{
		id:     "0005",
		title:  "aa bird 5",
		author: "cat 5",
	},
}

func getBooks(c *gin.Context) {
	c.IntentedJSON(http.StatusOK, bookshelf)
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
