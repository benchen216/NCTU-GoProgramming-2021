package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	// write your own struct
	ID    int    `json:"id,string"`
	NAME  string `json:"name"`
	PAGES int    `json:"pages,string"`
}

var bookshelf = []Book{
	// init data
	{
		ID:    1,
		NAME:  "Blue Bird",
		PAGES: 500,
	},
}

func getBooks(c *gin.Context) {
	fmt.Println("Using indented json")
	c.IndentedJSON(http.StatusOK, bookshelf)
}
func getBook(c *gin.Context) {
	// parse from url
	id := c.Param("id")
	// find the correspond id
	for i := 0; i < len(bookshelf); i++ {
		inVar, _ := strconv.Atoi(id)
		if bookshelf[i].ID == inVar {
			c.IndentedJSON(http.StatusOK, bookshelf[i])
			return
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "book not found"})
}
func addBook(c *gin.Context) {
	// add book into data structure
	var json Book
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i := 0; i < len(bookshelf); i++ {
		if bookshelf[i].ID == json.ID {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "duplicate book id"})
			return
		}
	}

	inp := Book{json.ID, json.NAME, json.PAGES}
	bookshelf = append(bookshelf, inp)

	// print in the console
	c.IndentedJSON(http.StatusOK, inp)
}
func deleteBook(c *gin.Context) {
	// parse from url
	id := c.Param("id")
	// find the correspond id
	for i := 0; i < len(bookshelf); i++ {
		inVar, _ := strconv.Atoi(id)
		if bookshelf[i].ID == inVar {
			// return value
			c.IndentedJSON(http.StatusOK, bookshelf[i])
			// remove the element
			bookshelf = append(bookshelf[:i], bookshelf[i+1:]...) // remove
			return
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "book not found"})
}
func updateBook(c *gin.Context) {
	var json Book
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// parse from url
	// id := c.Param("id")
	// find the correspond id
	for i := 0; i < len(bookshelf); i++ {
		// inVar, _ := strconv.Atoi(id)
		if bookshelf[i].ID == json.ID {

			bookshelf[i].ID = json.ID
			bookshelf[i].NAME = json.NAME
			bookshelf[i].PAGES = json.PAGES

			// print in the console
			c.IndentedJSON(http.StatusOK, bookshelf[i])
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

// curl -v -X GET http://localhost:8080/bookshelf \
//   -H 'content-type: application/json' \
//   -d '{"fisrtname":"eric", "lastname":"lu"}'
