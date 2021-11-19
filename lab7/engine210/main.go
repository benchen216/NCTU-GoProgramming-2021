package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"os"
	"net/http"
	"strconv"
)

type Book struct {
	ID 		int 	`json:"id"`
	Name 	string 	`json:"name"`
	Pages 	int 	`json:"pages"`
}

var bookshelf map[int]Book

func getBooks(c *gin.Context) {
	//fmt.Println(len(bookshelf))
	//fmt.Println(bookshelf)
	bookshelf_list := make([]Book, 0)
	for _, value := range bookshelf {
		bookshelf_list = append(bookshelf_list, value)
	}
	//fmt.Println(bookshelf_list)
	c.IndentedJSON(200, bookshelf_list)
}

func getBook(c *gin.Context) {
	//fmt.Println(c.Param("id"))
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid ID format")
		return
	}
	if _, ok := bookshelf[id]; !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, bookshelf[id])
}
func addBook(c *gin.Context) {
	var json Book
	if err := c.ShouldBindJSON(&json); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _, ok := bookshelf[json.ID]; ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "duplicate book id"})
		return
	}
	//fmt.Println(json.Pages)
	bookshelf[json.ID] = json
	c.IndentedJSON(http.StatusOK, json)
}

func deleteBook(c *gin.Context) {	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid ID format")
		return
	}
	if _, ok := bookshelf[id]; !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, bookshelf[id])
	delete(bookshelf, id)
}

func updateBook(c *gin.Context) {	
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid ID format")
		return
	}
	if _, ok := bookshelf[id]; !ok {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	var json Book
	if err := c.ShouldBindJSON(&json); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	bookshelf[id] = json
	c.IndentedJSON(http.StatusOK, bookshelf[id])
}

func main() {
	bookshelf = map[int]Book{
		1: {ID: 1, Name: "Blue Bird", Pages: 500,},
	}
	fmt.Println(bookshelf)
	
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
