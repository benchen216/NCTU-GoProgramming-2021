package main

import (
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
	// fmt.Println("Using indented json")
	str := "["
	for i := 0; i < len(bookshelf); i++ {
		tmp_str := " { \"id\": \"" + strconv.Itoa(bookshelf[i].ID) + "\", \"name\": \"" + bookshelf[i].NAME + "\", \"pages\": \"" + strconv.Itoa(bookshelf[i].PAGES) + "\" }"
		str = str + tmp_str
	}

	str = str + " ]"

	c.String(http.StatusOK, str)
	// c.IndentedJSON(http.StatusOK, bookshelf)
}
func getBook(c *gin.Context) {
	// parse from url
	id := c.Param("id")
	// find the correspond id
	for i := 0; i < len(bookshelf); i++ {
		inVar, _ := strconv.Atoi(id)
		if bookshelf[i].ID == inVar {
			tmp_str := "{ \"id\": \"" + strconv.Itoa(bookshelf[i].ID) + "\", \"name\": \"" + bookshelf[i].NAME + "\", \"pages\": \"" + strconv.Itoa(bookshelf[i].PAGES) + "\" }"
			c.String(http.StatusOK, tmp_str)
			// c.IndentedJSON(http.StatusOK, bookshelf[i])
			return
		}
	}
	c.String(http.StatusOK, "{ \"message\": \"book not found\" }")
	// c.IndentedJSON(http.StatusOK, gin.H{"message": "book not found"})
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
			c.String(http.StatusOK, "{ \"message\": \"duplicate book id\" }")
			// c.IndentedJSON(http.StatusOK, gin.H{"message": "duplicate book id"})
			return
		}
	}

	inp := Book{json.ID, json.NAME, json.PAGES}
	bookshelf = append(bookshelf, inp)

	// print in the console
	tmp_str := "{ \"id\": \"" + strconv.Itoa(json.ID) + "\", \"name\": \"" + json.NAME + "\", \"pages\": \"" + strconv.Itoa(json.PAGES) + "\" }"
	c.String(http.StatusOK, tmp_str)
	// c.IndentedJSON(http.StatusOK, inp)
}
func deleteBook(c *gin.Context) {
	flag := 0
	// parse from url
	id := c.Param("id")
	inVar, _ := strconv.Atoi(id)
	// find the correspond id
	for i := 0; i < len(bookshelf); i++ {
		if bookshelf[i].ID == inVar {
			// return value
			json := bookshelf[i]
			tmp_str := "{ \"id\": \"" + strconv.Itoa(json.ID) + "\", \"name\": \"" + json.NAME + "\", \"pages\": \"" + strconv.Itoa(json.PAGES) + "\" }"
			if flag == 0 {
				c.String(http.StatusOK, tmp_str)
			}
			// c.IndentedJSON(http.StatusOK, bookshelf[i])
			// remove the element
			bookshelf = append(bookshelf[:i], bookshelf[i+1:]...) // remove
			flag = 1
			i--
		}
	}

	if flag == 0 {
		c.String(http.StatusOK, "{ \"message\": \"book not found\" }")
	}
	// c.IndentedJSON(http.StatusOK, gin.H{"message": "book not found"})
}
func updateBook(c *gin.Context) {
	// my update is update the correspond id
	var json Book
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// parse from url
	id := c.Param("id")
	inVar, _ := strconv.Atoi(id)
	// find the correspond id
	for i := 0; i < len(bookshelf); i++ {
		if bookshelf[i].ID == inVar {

			bookshelf[i].ID = json.ID
			bookshelf[i].NAME = json.NAME
			bookshelf[i].PAGES = json.PAGES

			// print in the console
			tmp_str := "{ \"id\": \"" + strconv.Itoa(json.ID) + "\", \"name\": \"" + json.NAME + "\", \"pages\": \"" + strconv.Itoa(json.PAGES) + "\" }"
			c.String(http.StatusOK, tmp_str)
			// c.IndentedJSON(http.StatusOK, bookshelf[i])
			return
		}
	}
	c.String(http.StatusOK, "{ \"message\": \"book not found\" }")
	// c.IndentedJSON(http.StatusOK, gin.H{"message": "book not found"})
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

// https://golang-lab7.herokuapp.com/
