package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Pages string `json:"pages"`
}

var bookshelf = []Book{
	{
		Id:    "1",
		Name:  "Blue Bird",
		Pages: "500",
	},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(200, bookshelf)
}

func getBook(c *gin.Context) {
	id := c.Param("id")
	for _, s := range bookshelf {
		if s.Id == id {
			c.IndentedJSON(200, s)
			return
		}
	}
	c.IndentedJSON(200, gin.H{"message": "book not found"})
}

func addBook(c *gin.Context) {
	tempBook := Book{}
	c.BindJSON(&tempBook)
	for _, s := range bookshelf {
		if s.Id == tempBook.Id {
			c.IndentedJSON(200, gin.H{"message": "duplicate book id"})
			return
		}
	}
	bookshelf = append(bookshelf, tempBook)
	c.IndentedJSON(200, tempBook)
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")
	for index, s := range bookshelf {
		if s.Id == id {
			c.IndentedJSON(200, s)
			bookshelf = append(bookshelf[:index], bookshelf[index+1:]...)
			return
		}
	}
	c.IndentedJSON(200, gin.H{"message": "book not found"})
}

func modifyBook(c *gin.Context) {
	id := c.Param("id")
	tempBook := Book{}
	c.BindJSON(&tempBook)
	for idx, _ := range bookshelf {
		if bookshelf[idx].Id == id {
			bookshelf[idx] = tempBook
			c.IndentedJSON(http.StatusOK, bookshelf[idx])
			return
		}
	}
	c.IndentedJSON(200, gin.H{"message": "book not found"})
}

func main() {
	r := gin.Default()
	r.RedirectFixedPath = true
	r.GET("/bookshelf", getBooks)
	r.GET("/bookshelf/:id", getBook)
	r.POST("/bookshelf", addBook)
	r.DELETE("/bookshelf/:id", deleteBook)
	r.PUT("/bookshelf/*id", modifyBook)

	port := "8080"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	r.Run(":" + port)
}
