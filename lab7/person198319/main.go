package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Book struct {
	// write your own struct
	ID    string `json:"id"`
	NAME  string `json:"name"`
	PAGES string `json:"pages"`
}

var bookshelf = []Book{
	// init data
	{
		ID:    "1",
		NAME:  "Blue Bird",
		PAGES: "500",
	},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, bookshelf)
}
func getBook(c *gin.Context) {
	x := c.Param("id")
	for _, v := range bookshelf {
		if v.ID == x {
			c.IndentedJSON(http.StatusOK, v)
			return
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "book not found"})
}
func addBook(c *gin.Context) {
	var json Book
	if err := c.ShouldBindJSON(&json); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _, v := range bookshelf {
		if v.ID == json.ID {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "duplicate book id"})
			return
		}
	}
	bookshelf = append(bookshelf, json)
	c.JSON(http.StatusOK, json)
}
func deleteBook(c *gin.Context) {
	x := c.Param("id")
	for y := 0; y < len(bookshelf); y++ {
		if bookshelf[y].ID == x {
			c.IndentedJSON(http.StatusOK, bookshelf[y])
			bookshelf = append(bookshelf[:y], bookshelf[y+1:]...)
			return
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "book not found"})
}
func updateBook(c *gin.Context) {
	var json Book
	if err := c.ShouldBindJSON(&json); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	x := c.Param("id")
	for y := 0; y < len(bookshelf); y++ {
		if bookshelf[y].ID == x {
			bookshelf[y] = json
			c.IndentedJSON(http.StatusOK, json)
			return
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "book not found"})
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
