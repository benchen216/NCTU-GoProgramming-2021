package main

import (
	"github.com/gin-gonic/gin"
	//  "net/http"
	"os"
    "strings"
)

type Book struct {
	// write your own struct
	id string
	name string
	pages string
}

var bookshelf = []Book{
	// init data
	{id: "1", name: "Blue Bird", pages: "500"},
}

func tran(s string) string {
    //re := ""
    s = strings.Replace(s, "\n", " ", -1)
    s = strings.Replace(s, "\t", " ", -1)
    for i := 0; i < 100; i++ {
        s = strings.Replace(s, "  ", " ", -1)
    }
    return s
}

func getBooks(c *gin.Context) {
    a := make([]gin.H, len(bookshelf))
    for top, i := range bookshelf {
        a[top] = gin.H{
            "id": i.id,
            "name": i.name,
            "pages": i.pages,
        }
    }
    c.IndentedJSON(200, a)
}
func getBook(c *gin.Context) {
    id := c.Param("a")
    for _, i := range bookshelf {
        if i.id == id {
            c.IndentedJSON(200, gin.H{
                "id": i.id,
                "name": i.name,
                "pages": i.pages,
            })
            return
        }
    }
    c.IndentedJSON(200, gin.H{
        "message": "book not found",
    })
}
func addBook(c *gin.Context) {
    jsons := make(map[string]string)
    c.BindJSON(&jsons)
    for _, i := range bookshelf {
        if i.id == jsons["ID"] {
            c.IndentedJSON(200, gin.H{
                "message": "duplicate book id",
            })
            return
        }
    }
    a := Book{}
    a.id = jsons["ID"]
    a.name = jsons["NAME"]
    a.pages = jsons["PAGES"]
    bookshelf = append(bookshelf, a)
    c.IndentedJSON(200, gin.H {
        "id": jsons["ID"],
        "name": jsons["NAME"],
        "pages": jsons["PAGES"],
    })
}
func deleteBook(c *gin.Context) {
    id := c.Param("a")
    for p, i := range bookshelf {
        if i.id == id {
            c.IndentedJSON(200, gin.H{
                "id": i.id,
                "name": i.name,
                "pages": i.pages,
            })
            bookshelf = append(bookshelf[:p], bookshelf[p+1:]...)
            return
        }
    }
    c.IndentedJSON(200, gin.H{
        "message": "book not found",
    })
}
func updateBook(c *gin.Context) {
    println(1)
    jsons := make(map[string]string)
    c.BindJSON(&jsons)
    a := Book{}
    a.id = jsons["ID"]
    a.name = jsons["NAME"]
    a.pages = jsons["PAGES"]
    for p, i := range bookshelf {
        if i.id == jsons["ID"] {
            bookshelf[p] = a;
            c.IndentedJSON(200, gin.H {
                "id": jsons["ID"],
                "name": jsons["NAME"],
                "pages": jsons["PAGES"],
            })
            return
        }
    }
    c.IndentedJSON(200, gin.H{
        "message": "book not found",
    })
}
func main() {
	r := gin.Default()
	r.RedirectFixedPath = true
	r.GET("/bookshelf/:a", getBook)
	r.GET("/bookshelf", getBooks)
	r.POST("/bookshelf", addBook)
	r.DELETE("/bookshelf/:a", deleteBook)
	r.PUT("/bookshelf", updateBook)
	r.PUT("/bookshelf/:a", updateBook)

	port := "8080"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}
	r.Run(":" + port)
}

// curl -X POST localhost:8080 -H 'Content-Type: application/json' -d '{"login":"my_login","password":"my_password"}'
