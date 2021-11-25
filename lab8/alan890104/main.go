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
	Id    int    `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Pages string `json:"page" binding:"required"`
}

func getBooks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		/* [TODO] get all books data*/
		rows, _ := db.Query("SELECT * FROM bookshelf")

		/* [TODO] scan the data one by one*/
		defer rows.Close()
		var books []Book
		for rows.Next() {
			var b Book
			rows.Scan(&b.Id, &b.Name, &b.Pages)
			books = append(books, b)
		}

		//[TODO]send all data or error handling
		c.IndentedJSON(http.StatusOK, books)

	}
}
func getBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		/* [TODO] get all books data*/
		rows, _ := db.Query("SELECT * FROM bookshelf WHERE id=?", id)

		/* [TODO] scan the data one by one*/
		defer rows.Close()
		var books []Book
		for rows.Next() {
			var b Book
			rows.Scan(&b.Id, &b.Name, &b.Pages)
			books = append(books, b)
		}

		//[TODO]send all data or error handling
		if len(books) == 0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		} else {
			c.IndentedJSON(http.StatusOK, books)
		}

	}
}

func addBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var post Book
		c.Bind(&post)
		db.QueryRow("INSERT INTO bookshelf VALUES (?,?,?) RETURNING id", post.Id, post.Name, post.Pages).
			Scan(&post.Id)
		c.IndentedJSON(http.StatusOK, post)
	}
}

func updateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var post Book
		c.Bind(&post)

		err := db.QueryRow("UPDATE bookshelf SET name=?, pages=? WHERE id=? RETURNING *", post.Name, post.Pages, id).
			Scan(&post.Id, &post.Name, &post.Pages)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"message": "book not found",
			})
		} else {
			c.IndentedJSON(http.StatusOK, post)
		}
	}
}

func deleteBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var ret Book
		err := db.QueryRow("DELETE FROM bookshelf WHERE id=? RETURNING *", id).Scan(&ret.Id, &ret.Name, &ret.Pages)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"message": "book not found",
			})
		} else {
			c.IndentedJSON(http.StatusOK, ret)
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
		log.Fatal("Cannot Load env")
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
	r.POST("/bookshelf/", addBook(db))
	r.PUT("/boolshelf/:id", updateBook(db))
	r.DELETE("/bookshelf/:id", deleteBook(db))

	r.Run(":" + port)
}
