package main

import (
	"github.com/gin-gonic/gin"
	"os"
	// "fmt"
	// "encoding/json"
)

type Book struct {
	// write your own struct
	Id string `json:"id"`
	Name string `json:"name"`
	Pages string `json:"pages"`
}

var bookshelf = []Book{
	// init data
	Book {
		Id: "1",
		Name: "Blue Bird",
		Pages: "500",
	},
	
}

func getBooks(c *gin.Context) {
	// v_json, _ := json.Marshal(bookshelf)
	c.IndentedJSON(200, bookshelf)
}
func getBook(c *gin.Context) {
	id :=  c.Param("id")
	for _, v := range bookshelf {
		if v.Id == id {
			c.IndentedJSON(200, v)
			return 
		}
	}
	// error handling
	c.IndentedJSON(404, gin.H{
		"message": "book not found",
	})

}
func addBook(c *gin.Context) {
	var newBook Book
	c.BindJSON(&newBook)
	// error handling
	for _, v := range bookshelf {
		if newBook.Id == v.Id {
			c.IndentedJSON(404, "duplicate book id")
			return
		}
	}
	
	bookshelf = append(bookshelf, newBook)
	c.IndentedJSON(200, newBook)
}
func deleteBook(c *gin.Context) {
	id := c.Param("id")
	for i, v := range bookshelf {
		if id == v.Id {
			c.IndentedJSON(200, v)
			bookshelf = append(bookshelf[:i], bookshelf[i+1:]...)
			return
		}
	}
}
func main() {
	r := gin.Default()
	r.RedirectFixedPath = true
	r.GET("/bookshelf", getBooks)
	r.GET("/bookshelf/:id", getBook)
	r.POST("bookshelf/", addBook)
	r.DELETE("bookshelf/:id", deleteBook)
	port := "8080"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	r.Run(":" + port)
}
