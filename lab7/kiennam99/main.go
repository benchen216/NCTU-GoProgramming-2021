package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Book struct {
	// write your own struct
	ID    string `json:"id"`
	NAME  string `json:"name"`
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

func exist(s []Book, e string) (int, bool) {
	for index, err := range s {
		if err.ID == e {
			return index, true
		}
	}
	return 0, false
}

func parse(input string) (string, string, string) {
	var id, name, pages string
	temp := input[1 : len(input)-1]
	token := strings.Split(temp, ",")

	for index, value := range token {
		if index == 0 {
			tmp := strings.Split(value, ":")
			id = tmp[1][1:(len(tmp[1]) - 1)]
		} else if index == 1 {
			tmp := strings.Split(value, ":")
			name = tmp[1][1:(len(tmp[1]) - 1)]
		} else if index == 2 {
			tmp := strings.Split(value, ":")
			pages = tmp[1][1:(len(tmp[1]) - 1)]
		} else {
			fmt.Println("Error input")
		}

	}
	return id, name, pages
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(200, bookshelf)
}

func getBook(c *gin.Context) {
	id := c.Param("id")
	index, ex := exist(bookshelf, id)
	if ex {
		c.IndentedJSON(200, gin.H{
			"id":    id,
			"name":  bookshelf[index].NAME,
			"pages": bookshelf[index].PAGES,
		})
	} else {

		c.IndentedJSON(200, gin.H{
			"message": "book not found",
		})
	}
}
func addBook(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(200, err.Error())
	}

	id, name, pages := parse(string(value))
	_, errr := exist(bookshelf, id)
	if errr {

		c.IndentedJSON(200, gin.H{
			"message": "duplicate book id",
		})

	} else {
		ttmp := Book{id, name, pages}
		bookshelf = append(bookshelf, ttmp)
		c.IndentedJSON(200, bookshelf[len(bookshelf)-1])
	}

}
func deleteBook(c *gin.Context) {
	id := c.Param("id")
	index, ex := exist(bookshelf, id)
	if ex {
		c.IndentedJSON(200, gin.H{
			"id":    id,
			"name":  bookshelf[index].NAME,
			"pages": bookshelf[index].PAGES,
		})
		bookshelf = append(bookshelf[:index], bookshelf[index+1:]...)

	} else {

		c.IndentedJSON(200, gin.H{
			"message": "book not found",
		})
	}
}
func updateBook(c *gin.Context) {
	ids := c.Param("id")
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(200, err.Error())
	}

	id, name, pages := parse(string(value))
	index, errr := exist(bookshelf, ids)
	if !errr {

		c.IndentedJSON(200, gin.H{
			"message": "book not found",
		})

	} else {

		bookshelf[index].ID = id
		bookshelf[index].NAME = name
		bookshelf[index].PAGES = pages
		c.IndentedJSON(200, bookshelf[index])
	}
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
