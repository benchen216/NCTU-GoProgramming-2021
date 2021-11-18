package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Pages string `json:"pages"`
	// write your own struct
}

var bookshelf = []Book{
	// init data
	{
		ID:    "1",
		Name:  "Blue Bird",
		Pages: "500",
	},
}

func getBooks(c *gin.Context) {
	c.JSON(200, bookshelf)
}

func getBook(c *gin.Context) {
	ID := c.Param("id")

	for i := 0; i < len(bookshelf); i++ {
		if bookshelf[i].ID == ID {
			c.JSON(200, bookshelf[i])
			return
		}
	}

	c.JSON(404, gin.H{
		"message": "book not found",
	})
}

func addBook(c *gin.Context) {
	var json Book

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i := 0; i < len(bookshelf); i++ {
		if json.ID == bookshelf[i].ID {
			c.JSON(404, gin.H{
				"message": "duplicate book id",
			})

			return
		}
	}

	bookshelf = append(bookshelf, json)
	c.JSON(200, json)
}

func deleteBook(c *gin.Context) {
	ID := c.Param("id")

	for i := 0; i < len(bookshelf); i++ {
		if bookshelf[i].ID == ID {
			c.JSON(200, bookshelf[i])
			bookshelf = append(bookshelf[0:i], bookshelf[i+1:]...)
			return
		}
	}

	c.JSON(404, gin.H{
		"message": "book not found",
	})
}

func updateBook(c *gin.Context) {
	ID := c.Param("id")
	var json Book

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i := 0; i < len(bookshelf); i++ {
		if ID == bookshelf[i].ID {
			bookshelf[i] = json
			c.JSON(200, bookshelf[i])
			return
		}
	}

	c.JSON(404, gin.H{
		"message": "book not found",
	})
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
