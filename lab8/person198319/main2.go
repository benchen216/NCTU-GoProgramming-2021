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
	ID    int    `json:"id"`
	NAME  string `json:"name"`
	PAGES string `json:"pages"`
}

func getBooks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, _ := db.Query("SELECT * FROM bookshelf")

		defer rows.Close()

		books := []Book{}
		for rows.Next() {
			book := Book{}

			if rows.Scan(&book.ID, &book.NAME, &book.PAGES) != nil {
				log.Fatal(rows)
			}

			books = append(books, book)
		}
		//[TODO]send all data or error handling
		if len(books) == 0 {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "book not found"})
		} else {
			c.IndentedJSON(http.StatusOK, books)
		}
	}
}
func getBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		i := c.Param("id")

		book := Book{}
		err := db.QueryRow("SELECT * FROM bookshelf WHERE id=$1", i).Scan(&book.ID, &book.NAME, &book.PAGES)
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Not Found"})
		} else {
			c.IndentedJSON(http.StatusOK, book)
		}
	}
}

func addBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var submit Book
		if err2 := c.ShouldBindJSON(&submit); err2 != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
			return
		}
		err := db.QueryRow("INSERT INTO bookshelf VALUES (DEFAULT,$1,$2) RETURNING *", submit.NAME, submit.PAGES).Scan(&submit.ID, &submit.NAME, &submit.PAGES)
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "error"})
		} else {
			c.IndentedJSON(http.StatusOK, submit)
		}

	}
}
func updateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		i := c.Param("id")

		var submit Book
		if err2 := c.ShouldBindJSON(&submit); err2 != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
			return
		}
		err := db.QueryRow("UPDATE bookshelf SET name=$1, pages=$2 WHERE id=$3 RETURNING *", submit.NAME, submit.PAGES, i).Scan(&submit.ID, &submit.NAME, &submit.PAGES)
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Not Found"})
		} else {
			c.IndentedJSON(http.StatusOK, submit)
		}
	}
}

func deleteBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		i := c.Param("id")

		var book Book
		err := db.QueryRow("DELETE FROM bookshelf WHERE id=$1 RETURNING *", i).Scan(&book.ID, &book.NAME, &book.PAGES)
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "Not Found"})
		} else {
			c.IndentedJSON(http.StatusOK, book)
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

	// [TODO] other method

	r.Run(":" + port)
}
