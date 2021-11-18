package main

import (
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
	var idx = c.Param("id")
	for _, s := range bookshelf {
		if s.Id == idx {
			c.IndentedJSON(200, s)
			return
		}
	}
	c.IndentedJSON(200, gin.H{"message": "book not found"})
}

func addBook(c *gin.Context) {
	var b Book
	c.BindJSON(&b)
	for _, s := range bookshelf {
		if s.Id == b.Id {
			c.IndentedJSON(200, gin.H{"message": "duplicate book id"})
			return
		}
	}
	bookshelf = append(bookshelf, b)
	c.IndentedJSON(200, b)
}

func deleteBook(c *gin.Context) {
	var idx int = 0
	var target = c.Param("id")
	for id, s := range bookshelf {
		if s.Id == target {
			c.IndentedJSON(200, s)
			idx = id
			break
		}
	}
	if idx != 0 {
		bookshelf = append(bookshelf[:idx], bookshelf[idx+1:]...)
	} else {
		c.IndentedJSON(200, gin.H{"message": "book not found"})
	}
}

func modifyBook(c *gin.Context) {
	var target = c.Param("id")
	var b Book
	c.BindJSON(&b)
	for idx, _ := range bookshelf {
		if bookshelf[idx].Id == target {
			bookshelf[idx].Id = b.Id
			bookshelf[idx].Name = b.Name
			bookshelf[idx].Pages = b.Pages
			c.IndentedJSON(200, bookshelf[idx])
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
	r.PUT("/bookshelf/:id", modifyBook)

	port := "8080"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	r.Run(":" + port)
}
