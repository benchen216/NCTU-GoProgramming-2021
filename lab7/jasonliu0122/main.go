package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Book struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Pages string `json:"pages"`
}

var bookshelf = []Book{
	// init data
	{
		ID		: "1",
		Name	: "Blue Bird",
		Pages	: "500",
	},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, bookshelf)
}
func getBook(c *gin.Context) {
	id := c.Param("id")

	for i := range bookshelf {
		if bookshelf[i].ID == id {
			c.IndentedJSON(http.StatusOK, bookshelf[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
func addBook(c *gin.Context) {
	var newbook Book

	if err := c.BindJSON(&newbook); err != nil {
		c.IndentedJSON(http.StatusNotFound, "error")
	}
	for _, v := range bookshelf {
		if newbook.ID == v.ID {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "duplicate book id"})
			return
		}
	}

	bookshelf = append(bookshelf, newbook)
	c.IndentedJSON(http.StatusOK, bookshelf[len(bookshelf)-1])
}

func deleteBook(c *gin.Context) {
	id := c.Param("id")

	for i := range bookshelf {
		if bookshelf[i].ID == id {
			c.IndentedJSON(http.StatusOK, bookshelf[i])
			bookshelf = append(bookshelf[:i],bookshelf[i+1:]...)
			return
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "book not found"})


}
func updateBook(c *gin.Context) {
	var book Book
	id := c.Param("id")
	c.BindJSON(&book)

	for i := range bookshelf {
		if bookshelf[i].ID == id {	
			bookshelf[i] = book
			c.IndentedJSON(http.StatusOK, book)
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
