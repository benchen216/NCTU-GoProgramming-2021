package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Book struct {
	// write your own struct
	ID    string `json:"id"`
	NAME  string `json:"name"`
	PAGES string `json:"pages"`
}

var bookshelf = []Book{
	// init data
	{ID: "1", NAME: "Blue Bird", PAGES: "500"},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, bookshelf)
}

func getBook(c *gin.Context) {
	for _, book := range bookshelf {
		if book.ID == c.Param("id") {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "book not found"})
	return
}
func addBook(c *gin.Context) {
	var submit Book
	c.ShouldBindJSON(&submit)
	for _, book := range bookshelf {
		if book.ID == submit.ID {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "duplicate book id"})
			return
		}
	}
	bookshelf = append(bookshelf, submit)
	c.IndentedJSON(http.StatusOK, submit)
	return
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	for i := 0; i < len(bookshelf); i++ {
		if bookshelf[i].ID == id {
			c.IndentedJSON(http.StatusOK, bookshelf[i])
			bookshelf = append(bookshelf[:i], bookshelf[i+1:]...)
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "book not found"})
	return
}
func updateBook(c *gin.Context) {
	var submit Book
	c.ShouldBindJSON(&submit)
	id := c.Param("id")
	for i := 0; i < len(bookshelf); i++ {
		if bookshelf[i].ID == id {
			bookshelf[i] = submit
			c.JSON(http.StatusOK, submit)
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"message": "book not found"})
	return
}
func main() {
	r := gin.Default()
	r.RedirectFixedPath = true
	r.GET("/bookshelf", getBooks)
	r.GET("/bookshelf/:id", getBook)
	r.POST("/bookshelf", addBook)
	r.DELETE("/bookshelf/:id", deleteBook)
	r.PUT("/bookshelf/:id", updateBook)

	port := "8080"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	r.Run(":" + port)
}
