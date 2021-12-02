package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Book struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Pages string `json:"pages"`
}

func getBooks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		/* [TODO] get all books data */
		rows, _ := db.Query("SELECT * FROM bookshelf")

		/* [TODO] scan the data one by one */
		bookshelf := []Book{}
		defer rows.Close()
		for rows.Next() {
			book := Book{}
			rows.Scan(&book.Id, &book.Name, &book.Pages)
			bookshelf = append(bookshelf, book)
		}

		//[TODO]send all data or error handling
		if len(bookshelf) != 0 {
			c.IndentedJSON(http.StatusOK, bookshelf)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
	}
}

func getBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		book := Book{}
		err := db.QueryRow("SELECT * FROM bookshelf WHERE id=$1", id).Scan(&book.Id, &book.Name, &book.Pages)

		if err == nil {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
	}
}

func addBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book Book
		c.BindJSON(&book)
		db.QueryRow("INSERT INTO bookshelf VALUES (DEFAULT,$1,$2) RETURNING id", book.Name, book.Pages).Scan(&book.Id)
		c.IndentedJSON(http.StatusOK, book)
	}
}

func updateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var book Book
		c.BindJSON(&book)

		err := db.QueryRow("UPDATE bookshelf SET name=$1, pages=$2 WHERE id=$3 RETRUNING id", book.Name, book.Pages, id).Scan(&book.Id)
		if err == nil {
			c.IndentedJSON(http.StatusOK, book)
			return
		}
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found", "error": err})
	}
}

func deleteBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var book Book
		c.BindJSON(&book)

		err := db.QueryRow("DELETE FROM bookshelf WHERE id=$1 RETRUNING id", id).Scan(&book.Id)
		if err == nil {
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
	// [TODO] other method
	r.GET("/bookshelf/:id", getBook(db))
	r.POST("/bookshelf", addBook(db))
	r.DELETE("/bookshelf/:id", deleteBook(db))
	r.PUT("/bookshelf/:id", updateBook(db))

	r.Run(":" + port)
}
