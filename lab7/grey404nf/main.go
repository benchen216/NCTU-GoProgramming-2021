package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Book struct {
	// write your own struct
	Id string `json:"id"`
	Name string `json:"name"`
	Pages string `json:"pages"`
}

var bookshelf = []Book{
	// init data
	{
		Id: "1",
		Name: "Blue Bird",
		Pages: "500",
	},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, bookshelf)
}

func getBook(c *gin.Context) {
	id:=c.Param("id")
	for i:=0;i<len(bookshelf);i++ {
		if id==bookshelf[i].Id {
			c.IndentedJSON(http.StatusOK, bookshelf[i])
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func addBook(c *gin.Context) {
	var b Book
	
	c.BindJSON(&b)
	for i:=0;i<len(bookshelf);i++ {
		if b.Id==bookshelf[i].Id {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "duplicate book id"})
			return
		}
	}
	bookshelf=append(bookshelf, b)
	c.IndentedJSON(http.StatusOK, b)
}

func deleteBook(c *gin.Context) {
	id:=c.Param("id")
	for i:=0;i<len(bookshelf);i++ {
		if id==bookshelf[i].Id {
			c.IndentedJSON(http.StatusOK, bookshelf[i])
			bookshelf=append(bookshelf[:i], bookshelf[i+1:]...)
			return
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "book not found"})
}

func updateBook(c *gin.Context) {
	id:=c.Param("id")
	var b Book
	c.BindJSON(&b)
	for i:=0;i<len(bookshelf);i++ {
		if id==bookshelf[i].Id {
			bookshelf[i]=b
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
