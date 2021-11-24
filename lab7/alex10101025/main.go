package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	//"net/http"
	"os"
)

type Book struct {
	// write your own struct
	Id    string `json:"id"`
	Name  string `json:"name"`
	Pages string `json:"pages"`
}

var bookshelf = []Book{{Id: "1", Name: "Blue Bird", Pages: "500"}} // init data

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, bookshelf)
}
func getBook(c *gin.Context) {
	Id := c.Param("id")
	_, err := strconv.Atoi(Id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id is not a number"})
		return
	}
	for _, book := range bookshelf {
		if book.Id == Id {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
	return
}
func addBook(c *gin.Context) {
	var newbook Book
	err := c.BindJSON(&newbook)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "error"})
	}
	_, err = strconv.Atoi(newbook.Id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id is not a number"})
		return
	}
	for _, oldbook := range bookshelf {
		if oldbook.Id == newbook.Id {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "duplicate book id"})
			return
		}
	}
	bookshelf = append(bookshelf, newbook)
	c.IndentedJSON(http.StatusOK, newbook)
	return
}
func deleteBook(c *gin.Context) {
	Id := c.Param("id")
	_, err := strconv.Atoi(Id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id is not a number"})
		return
	}
	for idx, book := range bookshelf {
		if book.Id == Id {
			c.IndentedJSON(http.StatusOK, book)
			bookshelf = append(bookshelf[:idx], bookshelf[idx+1:]...)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
func updateBook(c *gin.Context) {
	Id := c.Param("id")
	var newbook Book
	err := c.BindJSON(&newbook)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "error"})
	}
	_, err = strconv.Atoi(Id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id is not a number"})
		return
	}
	for idx, book := range bookshelf {
		if book.Id == Id {
			bookshelf[idx] = newbook
			c.IndentedJSON(http.StatusOK, newbook)
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
