package main

import (
	"database/sql"
	"log"
	"os"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Book struct {
	// write your own struct
	// id's type is int
	Id int `json:"id"`
	Name string `json:"name"`
	Pages string `json:"pages"`
}

func getBooks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var b Book
		var bookshelf []Book
		
		// [TODO] get all books data
		rows, _ := db.Query("SELECT * FROM bookshelf")
		
		// [TODO] scan the data one by one
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&b.Id ,&b.Name ,&b.Pages)
			bookshelf=append(bookshelf, b)
		}
		
		//[TODO]send all data or error handling
		if len(bookshelf)==0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		} else {
			c.IndentedJSON(http.StatusOK, bookshelf)
		}
	}
}
func getBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id:=c.Param("id")
		var b Book
		row := db.QueryRow("SELECT * FROM bookshelf WHERE id=$1", id)
		row.Scan(&b.Id ,&b.Name ,&b.Pages)
		
		if b.Id==0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		} else {
			c.IndentedJSON(http.StatusOK, b)
		}
	}
}

func addBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var b Book
		c.BindJSON(&b)
		row := db.QueryRow("INSERT INTO bookshelf VALUES (DEFAULT, $1, $2) RETURNING id", b.Name, b.Pages)
		row.Scan(&b.Id)
		c.IndentedJSON(http.StatusOK, b)
	}
}

func updateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id:=c.Param("id")
		var b Book
		c.BindJSON(&b)
		row := db.QueryRow("UPDATE bookshelf SET name=$1, pages=$2 WHERE id=$3 RETURNING id",b.Name, b.Pages, id)
		row.Scan(&b.Id)
		
		if b.Id==0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		} else {
			c.IndentedJSON(http.StatusOK, b)
		}
	}
}

func deleteBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id:=c.Param("id")
		var b Book
		row := db.QueryRow("DELETE FROM bookshelf WHERE id=$1 RETURNING *", id)
		row.Scan(&b.Id ,&b.Name ,&b.Pages)
		
		if b.Id==0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		} else {
			c.IndentedJSON(http.StatusOK, b)
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
	// [TODO] other method
	r.GET("/bookshelf/:id", getBook(db))
	r.POST("/bookshelf", addBook(db))
	r.PUT("/bookshelf/:id", updateBook(db))
	r.DELETE("/bookshelf/:id", deleteBook(db))
	
	
	r.Run(":" + port)
}
