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
	i := c.Param("id")
	for _,v := range bookshelf {
		if v.ID == i {
			c.IndentedJSON(http.StatusOK,v)
			return
		}
	}
	c.IndentedJSON(http.StatusOK,gin.H{"message": "book not found"})
}
func addBook(c *gin.Context) {
	var json Book
	if err := c.ShouldBindJSON(&json); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for _,v := range bookshelf {
		if v.ID == json.ID {
			c.IndentedJSON(http.StatusOK,gin.H{"message": "duplicate book id"})
			return
		}
	}
	bookshelf = append(bookshelf,json)
	c.JSON(http.StatusOK,json)
}
func deleteBook(c *gin.Context) {
	i := c.Param("id")
	for j := 0; j < len(bookshelf); j++ {
		if bookshelf[j].ID == i{
			c.IndentedJSON(http.StatusOK,bookshelf[j])
			bookshelf = append(bookshelf[:j],bookshelf[j+1:]...)
			return
		}
	}
	c.IndentedJSON(http.StatusOK,gin.H{"message": "book not found"})
}
func updateBook(c *gin.Context) {
	var json Book
	if err := c.ShouldBindJSON(&json); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	i := c.Param("id")
	for j := 0; j < len(bookshelf); j++ {
		if bookshelf[j].ID == i{
			bookshelf[j] = json
			c.IndentedJSON(http.StatusOK,json)
			return
		}
	}
	c.IndentedJSON(http.StatusOK,gin.H{"message": "book not found"})
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
  