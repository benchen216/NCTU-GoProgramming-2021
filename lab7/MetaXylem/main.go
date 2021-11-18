package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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
	c.JSON(http.StatusOK, bookshelf)
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
			c.JSON(http.StatusOK, v)
			flag = false
			break
		}
	}
	if flag {
		c.JSON(http.StatusOK, err)
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
			c.JSON(http.StatusOK, err)
			flag = false
			break
		}
	}
	if flag {
		bookshelf = append(bookshelf, b)
		c.JSON(http.StatusOK, b)
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
			c.JSON(http.StatusOK, bookshelf[i])
			bookshelf = append(bookshelf[:i], bookshelf[i+1:]...)
			flag = false
			break
		}
	}
	if flag {
		c.JSON(http.StatusOK, err)
	}
}
func main() {
	r := gin.Default()
	r.RedirectFixedPath = true
	r.GET("/bookshelf", getBooks)
	r.GET("/bookshelf/:id", getBook)
	r.POST("/bookshelf", addBook)
	r.DELETE("/bookshelf/:id", deleteBook)
	port := "8080"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	r.Run("localhost:" + port)
}
