package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Book struct {
	// write your own struct
	// id's type is int
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Pages string `json:"pages"`
}

func getBooks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// [TODO] get all books data
		rows, err := db.Query("SELECT * from bookshelf")
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "error1"})
		}
		// [TODO] scan the data one by one
		defer rows.Close()
		var bookshelf = []Book{}
		for rows.Next() {
			var book Book
			rows.Scan(&book.Id, &book.Name, &book.Pages)
			bookshelf = append(bookshelf, book)
		}
		if len(bookshelf) == 0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		} else {
			c.IndentedJSON(http.StatusOK, bookshelf)
		}
		//[TODO]send all data or error handling
	}
}
func getBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT * from bookshelf WHERE id=$1", c.Param("id"))
		defer rows.Close()
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id is not a number"})
		}
		var book Book
		for rows.Next() {
			rows.Scan(&book.Id, &book.Name, &book.Pages)
			c.IndentedJSON(http.StatusOK, book)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
	}
}

func addBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newbook Book
		err := c.BindJSON(&newbook)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "error"})
		}
		rows, err := db.Query("INSERT INTO bookshelf VALUES (DEFAULT,$1,$2) RETURNING id", newbook.Name, newbook.Pages)
		defer rows.Close()
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "error2"})
			return
		}
		for rows.Next() {
			rows.Scan(&newbook.Id)
			c.IndentedJSON(http.StatusOK, newbook)
		}
	}
}

func updateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//rows, err := db.Query("?????????????????????????", ??, ??, ??)
		Id := c.Param("id")
		var newbook Book
		err := c.BindJSON(&newbook)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "error1"})
		}
		rows, err := db.Query("UPDATE bookshelf SET name=$2,pages=$3 WHERE id=$1 RETURNING *", Id, newbook.Name, newbook.Pages)
		defer rows.Close()
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "error2"})
		}
		for rows.Next() {
			rows.Scan(&newbook.Id, &newbook.Name, &newbook.Pages)
			c.IndentedJSON(http.StatusOK, newbook)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
	}
}

func deleteBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		Id := c.Param("id")
		rows, err := db.Query("DELETE FROM bookshelf WHERE id=$1 RETURNING *", Id)
		defer rows.Close()
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "error"})
			return
		}
		var book Book
		for rows.Next() {
			rows.Scan(&book.Id, &book.Name, &book.Pages)
			c.IndentedJSON(http.StatusOK, book)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
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
	// [TODO] other method

	r.Run(":" + port)
}
