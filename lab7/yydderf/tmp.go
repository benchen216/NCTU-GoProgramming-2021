package main

import (
        "github.com/gin-gonic/gin"
        "net/http"
        "os"
)

type Book struct {
    ID      string `json:"id"`
    Name    string `json:"name"`
    Pages   string `json:"pages"`
}

var bookshelf = []Book{
    {ID: "1", Name: "Blue Bird", Pages: "500"},
}

func getBooks(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, bookshelf)
}

func getBook(c *gin.Context) {
    ID := c.Param("id")

    for _,a:= range bookshelf{
        if a.ID== ID{
            c.IndentedJSON(http.StatusOK,a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func addBook(c *gin.Context) {
    var newBook Book

    if err := c.BindJSON(&newBook); err != nil {
        return
    }

    for _,a := range bookshelf{
        if a.ID == newBook.ID{
            c.IndentedJSON(http.StatusConflict, gin.H{"message": "duplicate book id"})
            return
        }
    }

    bookshelf = append(bookshelf, newBook)
    c.IndentedJSON(http.StatusCreated, newBook)
}

func deleteBook(c *gin.Context) {
    ID := c.Param("id")

    for i,a:= range bookshelf{
        if a.ID == ID{
            // implement book removal
            c.IndentedJSON(http.StatusOK,a)
            bookshelf[i] = bookshelf[len(bookshelf)-1]
            bookshelf = bookshelf[:len(bookshelf)-1]
            return
        }
    }

    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

func updateBook(c *gin.Context) {
    var newBook Book
    if err:= c.BindJSON(&newBook); err != nil{
        return
    }
    id := c.Param("id")

    for i,a:= range bookshelf{
        if a.ID == id{
            // implement book update
            bookshelf[i] = newBook
            c.IndentedJSON(http.StatusOK, newBook)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message" :"book not found"})
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
