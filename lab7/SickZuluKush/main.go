package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Book struct {
	// write your own struct
	ID string `json:"id"`
	NAME string `json:"name"`
	PAGE string `json:"pages"`
}

var bookshelf = []Book{
	// init data
	{ID: "1", NAME: "Blue Bird", PAGE: "500"},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, bookshelf)
}

func getBook(c *gin.Context) {
	id := c.Param("id")
	
	for _, x := range bookshelf {
		if x.ID == id {
			c.IndentedJSON(http.StatusOK, x)
			return
		}
	}
	
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func addBook(c *gin.Context) {
	var newBook Book
	
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	
	for _, x := range bookshelf {
		if newBook.ID == x.ID {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "duplicate book id"})
			return
		}
	}
	
	bookshelf = append(bookshelf, newBook)
	c.IndentedJSON(http.StatusOK, newBook)
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	var idx int
	flag := false
	
	for i, x := range bookshelf {
		if x.ID == id {
			idx = i
			c.IndentedJSON(http.StatusOK, x)
			flag = true
			break
		}
	}
	
	if flag {
		bookshelf = append(bookshelf[:idx], bookshelf[idx+1:]...)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
	}
}

func updateBook(c *gin.Context) {
	id := c.Param("id")
	var newBook Book
	
	if err := c.BindJSON(&newBook); err != nil {
		return
	}
	
	for i, x := range bookshelf {
		if id == x.ID {
			bookshelf[i] = newBook
			c.IndentedJSON(http.StatusOK, bookshelf[i])
			return
		}
	}
	
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
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
