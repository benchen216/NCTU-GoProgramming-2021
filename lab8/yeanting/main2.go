package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"
	"strings"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Book struct {
	id int
	name string
	pages string
}

func getBooks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM bookshelf")

		defer rows.Close()

		var Books []Book
		var book Book
		if err==nil {
			for rows.Next() {
				rows.Scan(&book.id ,&book.name ,&book.pages)
				Books = append(Books, book)
			}
			if len(Books)==0 {
				// Error Handling
				c.IndentedJSON(http.StatusOK, gin.H{
					"message": "book not found",
				})
			} else {
				res := "[ "
				for index, element := range Books {
					if index != (len(Books)-1) {
						res = res + `{ "id": "` + strconv.Itoa(element.id) + `", ` +
							`"name": "` + element.name + `", ` +
							`"pages": "` + element.pages + `" },`
						} else {
						res = res + `{ "id": "` + strconv.Itoa(element.id) + `", ` +
							`"name": "` + element.name + `", ` +
							`"pages": "` + element.pages + `" }`
						}
					}
				res += " ]"
				res_json := []byte(res)
				//c.JSON(http.StatusOK, res)
				c.Data(http.StatusOK, "application/json", res_json)
			}
		} else {
			// Error Handling
			c.IndentedJSON(http.StatusOK, gin.H{
				//"message": "duplicate book id",
				"message": "sql error",
			})
		}
	}
}
func getBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		
		rows, err := db.Query("SELECT * FROM bookshelf WHERE id=$1", id)
		if err != nil {
			// Error Handling
			c.IndentedJSON(http.StatusOK, gin.H{
				//"message": "duplicate book id",
				"message": "sql error",
			})
		}
		var book Book
		for rows.Next() {
			rows.Scan(&book.id ,&book.name ,&book.pages)
		}
		if book.id==0 {
			// Error Handling
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "book not found",
			})
		} else{
			c.IndentedJSON(http.StatusOK, gin.H{
				"id": strconv.Itoa(book.id),
				"name": book.name,
				"pages": book.pages,
			})
		}
	}
}

func addBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var i map[string]interface{}
		_ = c.Bind(&i)
		lower_i := make(map[string]interface{}, len(i))
		for key, val := range i {
			lower_i[strings.ToLower(key)] = val
		}
		i = lower_i

		name := i["name"].(string)
		pages, _ := i["pages"].(string)
		
		rows, err := db.Query("INSERT INTO bookshelf (name, pages) VALUES ($1, $2) RETURNING *", name, pages)
		if err != nil {
			// Error Handling
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "sql error",
			})
			return
		}
		var book Book
		for rows.Next() {
			rows.Scan(&book.id ,&book.name ,&book.pages)
		}
		// response message
		c.IndentedJSON(http.StatusOK, gin.H{
			"id": strconv.Itoa(book.id),
			"name": book.name,
			"pages": book.pages,
		})
	}
}

func updateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var i map[string]interface{}
		_ = c.Bind(&i)
		lower_i := make(map[string]interface{}, len(i))
		for key, val := range i {
			lower_i[strings.ToLower(key)] = val
		}
		i = lower_i

		id := c.Param("id")
		name := i["name"].(string)
		pages, _ := strconv.Atoi(i["pages"].(string))
		
		rows, err := db.Query("UPDATE bookshelf SET name=$1, pages=$2 WHERE id=$3 RETURNING *", name, pages, id)

		if err != nil {
			// Error Handling
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "sql error",
			})
			return
		}
		var book Book
		for rows.Next() {
			rows.Scan(&book.id ,&book.name ,&book.pages)
		}
		if book.id==0 {
			// Error Handling
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "book not found",
			})
		} else {
			//send the updated book information
			c.IndentedJSON(http.StatusOK, gin.H{
				"id": strconv.Itoa(book.id),
				"name": book.name,
				"pages": book.pages,
			})
		}
	}
}

func deleteBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		rows, err := db.Query("DELETE FROM bookshelf WHERE id=$1 RETURNING *", id)
		if err != nil {
			// Error Handling
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "sql error",
			})
			return
		}
		var book Book
		for rows.Next() {
			rows.Scan(&book.id ,&book.name ,&book.pages)
		}
		if book.id==0 {
			// Error Handling
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "book not found",
			})
		} else {
			//send the removing book information
			c.IndentedJSON(http.StatusOK, gin.H{
				"id": strconv.Itoa(book.id),
				"name": book.name,
				"pages": book.pages,
			})
		}
	}
}

func ResetDBTable(db *sql.DB) {
	if _, err := db.Exec("DROP TABLE IF EXISTS bookshelf"); err != nil {
		return
	}
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS bookshelf (id SERIAL PRIMARY KEY, name VARCHAR(100), pages VARCHAR(10))"); err != nil {
		return
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		//Do nothing
	}
	port := "8080"
	if v := os.Getenv("PORT"); len(v) > 0 {
		port = v
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	ResetDBTable(db)

	r := gin.Default()
	r.RedirectFixedPath = true
	r.GET("/bookshelf", getBooks(db))
	r.GET("/bookshelf/:id", getBook(db))
	r.POST("/bookshelf", addBook(db))
	r.DELETE("/bookshelf/:id", deleteBook(db))
	r.PUT("/bookshelf/:id", updateBook(db))

	r.Run(":" + port)
}
