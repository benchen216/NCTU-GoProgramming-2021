package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	id int
	name string
	pages int
}

var bookshelf = []Book{
	// init data
	{1, "Blue Bird", 500},
}

func getBooks(c *gin.Context) {
	res := "[ "
	for index, element := range bookshelf {
		if index != (len(bookshelf)-1) {
			res = res + `{ "id": "` + strconv.Itoa(element.id) + `", ` +
				`"name": "` + element.name + `", ` +
				`"pages": "` + strconv.Itoa(element.pages) + `" },`
			} else {
			res = res + `{ "id": "` + strconv.Itoa(element.id) + `", ` +
				`"name": "` + element.name + `", ` +
				`"pages": "` + strconv.Itoa(element.pages) + `" }`
			}
		}
	res += " ]"
	res_json := []byte(res)
	//c.JSON(http.StatusOK, res)
	c.Data(http.StatusOK, "application/json", res_json)
}
func getBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, element := range bookshelf {
		if id == element.id {
			c.IndentedJSON(http.StatusOK, gin.H{
				"id": strconv.Itoa(element.id),
				"name": element.name,
				"pages": strconv.Itoa(element.pages),
			})
			return
		}
	}
	// Error Handling
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "book not found",
	})
}
func addBook(c *gin.Context) {
	var i map[string]interface{}
	_ = c.Bind(&i)
	lower_i := make(map[string]interface{}, len(i))
	for key, val := range i {
		lower_i[strings.ToLower(key)] = val
	}
	i = lower_i

	id, _ := strconv.Atoi(i["id"].(string))
	name := i["name"].(string)
	pages, _ := strconv.Atoi(i["pages"].(string))
	
	for _, element := range bookshelf {
		if id == element.id {
			// Error Handling
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "duplicate book id",
			})
			return
		}
	}
	// add to bookshelf
	book := Book{id, name, pages}
	bookshelf = append(bookshelf, book)
	// response message
	c.IndentedJSON(http.StatusOK, gin.H{
		"id": strconv.Itoa(id),
		"name": name,
		"pages": strconv.Itoa(pages),
	})
}
func deleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for index, element := range bookshelf {
		if id == element.id {
			//send the removing book information
			c.IndentedJSON(http.StatusOK, gin.H{
				"id": strconv.Itoa(element.id),
				"name": element.name,
				"pages": strconv.Itoa(element.pages),
			})
			// remove book from bookshelf slice
			bookshelf = append(bookshelf[:index], bookshelf[index+1:]...)
			return
		}
	}
}
func updateBook(c *gin.Context) {
	var i map[string]interface{}
	_ = c.Bind(&i)
	lower_i := make(map[string]interface{}, len(i))
	for key, val := range i {
		lower_i[strings.ToLower(key)] = val
	}
	i = lower_i

	id, _ := strconv.Atoi(c.Param("id"))

	id_new, _ := strconv.Atoi(i["id"].(string))
	name := i["name"].(string)
	pages, _ := strconv.Atoi(i["pages"].(string))

	for index, element := range bookshelf {
		if id == element.id {
			// update information of book
			bookshelf[index].id = id_new
			bookshelf[index].name = name
			bookshelf[index].pages = pages
			//send the updated book information
			c.IndentedJSON(http.StatusOK, gin.H{
				"id": strconv.Itoa(id_new),
				"name": name,
				"pages": strconv.Itoa(pages),
			})
			return
		}
	}
	// Error Handling
	c.IndentedJSON(http.StatusOK, gin.H{
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
