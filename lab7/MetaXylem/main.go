package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

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

func modifier(v []byte) string {
	s := string(v)
	s = strings.Replace(s, "[", "[ ", -1)
	s = strings.Replace(s, "{", "{ ", -1)
	s = strings.Replace(s, ",", ", ", -1)
	s = strings.Replace(s, ":", ": ", -1)
	s = strings.Replace(s, "]", " ]", -1)
	s = strings.Replace(s, "}", " }", -1)
	return s
}

func getBooks(c *gin.Context) {
	str, _ := json.Marshal(bookshelf)
	c.String(http.StatusOK, modifier(str))
}
func getBook(c *gin.Context) {
	type Err struct {
		Message string `json:"message"`
	}
	err := Err{
		"book not found",
	}
	ID := c.Param("id")
	flag := true
	for _, v := range bookshelf {
		if v.Id == ID {
			str, _ := json.Marshal(v)
			c.String(http.StatusOK, modifier(str))
			flag = false
			break
		}
	}
	if flag {
		str, _ := json.Marshal(err)
		c.String(http.StatusOK, modifier(str))
	}
}
func addBook(c *gin.Context) {
	type Err struct {
		Message string `json:"message"`
	}
	err := Err{
		"duplicate book id",
	}
	var b Book
	c.BindJSON(&b)
	flag := true
	for _, v := range bookshelf {
		if v.Id == b.Id {
			str, _ := json.Marshal(err)
			c.String(http.StatusOK, modifier(str))
			flag = false
			break
		}
	}
	if flag {
		bookshelf = append(bookshelf, b)
		str, _ := json.Marshal(b)
		c.String(http.StatusOK, modifier(str))
	}
}
func deleteBook(c *gin.Context) {
	type Err struct {
		Message string `json:"message"`
	}
	err := Err{
		"book not found",
	}
	ID := c.Param("id")
	flag := true
	for i := 1; i < len(bookshelf); i++ {
		if bookshelf[i].Id == ID {
			str, _ := json.Marshal(bookshelf[i])
			c.String(http.StatusOK, modifier(str))
			bookshelf = append(bookshelf[:i], bookshelf[i+1:]...)
			flag = false
			break
		}
	}
	if flag {
		str, _ := json.Marshal(err)
		c.String(http.StatusOK, modifier(str))
	}
}
func updateBook(c *gin.Context) {
	type Err struct {
		Message string `json:"message"`
	}
	err := Err{
		"book not found",
	}
	var b Book
	c.BindJSON(&b)
	ID := c.Param("id")
	flag := true

	for i := 1; i < len(bookshelf); i++ {
		if bookshelf[i].Id == ID {
			bookshelf[i] = b
			str, _ := json.Marshal(bookshelf[i])
			c.String(http.StatusOK, modifier(str))
			flag = false
			break
		}
	}
	if flag {
		str, _ := json.Marshal(err)
		c.String(http.StatusOK, modifier(str))
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