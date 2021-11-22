package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Book struct {
    ID      string `json:"id"`
    Name    string `json:"name"`
    Pages   string `json:"pages"`
}

var bookshelf = []Book{
    {ID: "1", Name: "Blue Bird", Pages: "500"},
}

func getBooks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, bookshelf)
}

func getBook(c *gin.Context) {
    id := c.Param("id")

    for _, tmp := range bookshelf {
        if tmp.ID == id {
            c.IndentedJSON(http.StatusOK, tmp)
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

    for _, book := range bookshelf {
        if book.ID == newBook.ID {
            c.IndentedJSON(http.StatusConflict, gin.H{"message": "duplicate book id"})
            return
        }
    }

    bookshelf = append(bookshelf, newBook)
    c.IndentedJSON(http.StatusCreated, newBook)
}

func deleteBook(c *gin.Context) {
    id := c.Param("id")

    for i, book := range bookshelf {
        if book.ID == id {
            c.IndentedJSON(http.StatusOK, book)
            bookshelf[i] = bookshelf[len(bookshelf)-1]
            bookshelf = bookshelf[:len(bookshelf)-1]
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func updateBook(c *gin.Context) {
    var newBook Book
    if err := c.BindJSON(&newBook); err != nil {
        return
    }
    id := c.Param("id")

    for i, book := range bookshelf {
        if book.ID == id {
            bookshelf[i] = newBook
            c.IndentedJSON(http.StatusOK, newBook)
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
    r.PUT("/bookshelf/:id", updateBook)
    r.DELETE("/bookshelf/:id", deleteBook)

	port := "8080"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	r.Run(":" + port)
}
