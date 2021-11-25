package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Book struct {
	// write your own struct
	// id's type is int
	ID    int `json:"id"`
	Name  string `json:"name"`
	Pages string `json:"pages"`
}

func getBooks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// [TODO] get all books data
		rows, _ := db.Query("SELECT * FROM bookshelf") 

		// [TODO] scan the data one by one
		defer rows.Close()
		var book Book
		var bookList []Book
		for rows.Next() {
			rows.Scan(&book.ID ,&book.Name ,&book.Pages)
			bookList = append(bookList,book)
		}

		//[TODO]send all data or error handling
		if len(bookList) > 0 {
			c.IndentedJSON(200, bookList)
		} else {
			c.IndentedJSON(404, gin.H{"message": "book not found"})
		}
	}
}
func getBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		n, _ := strconv.Atoi(c.Param("id"))
		rows := db.QueryRow("SELECT * FROM bookshelf WHERE id=$1", n)
		var book Book
		rows.Scan(&book.ID ,&book.Name ,&book.Pages)
		if book.ID != 0{
			c.IndentedJSON(200, book)
		} else {
			c.IndentedJSON(404, gin.H{"message": "book not found"})
		}
	}
}

func addBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var book Book
		c.BindJSON(&book)
		db.QueryRow("INSERT INTO bookshelf VALUES (DEFAULT,$1,$2) RETURNING id", book.Name, book.Pages).Scan(&book.ID)
		c.IndentedJSON(200, book)
	}
}

func updateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		n, _ := strconv.Atoi(c.Param("id"))
		var book,oldBook Book
		c.BindJSON(&book)
		rows := db.QueryRow("SELECT * FROM bookshelf WHERE id=$1", n)
		rows.Scan(&oldBook.ID ,&oldBook.Name ,&oldBook.Pages)
		
//		db.QueryRow("UPDATE bookshelf SET name=$1, pages=$2 WHERE id=$3", book.Name, book.Pages, n).Scan(&book.ID)
		if oldBook.ID != 0{
			db.QueryRow("UPDATE bookshelf SET name=$1, pages=$2 WHERE id=$3", book.Name, book.Pages, n)
			book.ID = n
			c.IndentedJSON(200, book)
		} else {
			c.IndentedJSON(404, gin.H{"message": "book not found"})
		}
	}
}

func deleteBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		n, _ := strconv.Atoi(c.Param("id"))
		var book Book
		rows := db.QueryRow("SELECT * FROM bookshelf WHERE id=$1", n)
		rows.Scan(&book.ID ,&book.Name ,&book.Pages)
		if book.ID != 0{
			c.IndentedJSON(200, book)		
			db.Query("DELETE FROM bookshelf WHERE id=$1 RETURNING *", n)
		} else {
			c.IndentedJSON(404, gin.H{"message": "book not found"})
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
	r.DELETE("/bookshelf/:id", deleteBook(db))
	r.PUT("/bookshelf/:id", updateBook(db))

	r.Run(":" + port)
}
