package main

import (
	"database/sql"
	"log"
	"os"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Book struct {
	ID    int `json:"id"`
	Name  string `json:"name"`
	Pages string `json:"pages"`
}

func sti(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("Convert string to int error")
		return -1
	}
	return i
}

func getBooks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		/* [TODO] get all books data */
		rows, err := db.Query("SELECT * FROM bookshelf")
		if err != nil {
			log.Fatal("Error query")
			return
		}
		/* [TODO] scan the data one by one */
		defer rows.Close()
		bookshelf_list := make([]Book, 0)
		var id, name, pages string
		for rows.Next() {
			rows.Scan(&id ,&name ,&pages)
			bookshelf_list = append(bookshelf_list, Book{sti(id), name, pages})
		}
		if len(bookshelf_list) == 0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
			return
		}
		//[TODO]send all data or error handling
		c.IndentedJSON(200, bookshelf_list)
	}
}
func getBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Param("id"))
		row := db.QueryRow("SELECT * FROM bookshelf WHERE id=$1", c.Param("id"))
		var id, name, pages string
		err := row.Scan(&id ,&name ,&pages)
		if err != nil {
			log.Println("Wrong id query. ", err)
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
			return
		}
		c.IndentedJSON(http.StatusOK, Book{sti(id), name, pages})
	}
}

func addBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var json Book
		if err := c.ShouldBindJSON(&json); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		row := db.QueryRow("SELECT * FROM bookshelf WHERE name=$1", json.Name)
		var id, name, pages string
		err := row.Scan(&id ,&name ,&pages)
		if err == nil { // book found, duplicate id
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "duplicate book name"})
			return
		}

		row = db.QueryRow("INSERT INTO bookshelf (id, name, pages) VALUES (DEFAULT, $1, $2) RETURNING *", json.Name, json.Pages)
		err = row.Scan(&id ,&name ,&pages)
		if err != nil {
			log.Println("Wrong addbook query. ", err)
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Unknown error"})
			return
		}
		c.IndentedJSON(http.StatusOK, Book{sti(id), name, pages})
	}
}

func updateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json Book
		if err := c.ShouldBindJSON(&json); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		row := db.QueryRow("SELECT * FROM bookshelf WHERE id=$1", c.Param("id"))
		var id, name, pages string
		err := row.Scan(&id ,&name ,&pages)
		if err != nil { // book not found
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
			return
		}

		row = db.QueryRow("UPDATE bookshelf SET name=$1, pages=$2 WHERE id=$3 RETURNING *", json.Name, json.Pages, id)
		err = row.Scan(&id ,&name ,&pages)
		if err != nil {
			log.Println("Wrong addbook query. ", err)
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Unknown error"})
			return
		}
		c.IndentedJSON(http.StatusOK, Book{sti(id), name, pages})
	}
}

func deleteBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(c.Param("id"))
		row := db.QueryRow("DELETE FROM bookshelf WHERE id=$1 RETURNING *", c.Param("id"))
		var id, name, pages string
		err := row.Scan(&id ,&name ,&pages)
		if err != nil {
			log.Println("Wrong id query. ", err)
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
			return
		}
		c.IndentedJSON(http.StatusOK, Book{sti(id), name, pages})
	}
}

func ResetDBTable(db *sql.DB) {
	if _, err := db.Exec("DROP TABLE IF EXISTS bookshelf"); err != nil {
		log.Fatal("Drop table failed. ", err)
		return
	}
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS bookshelf (id SERIAL PRIMARY KEY, name VARCHAR(100), pages VARCHAR(10))"); err != nil {
		log.Fatal("Create table failed. ", err)
		return
	}
	// db.Exec("INSERT INTO bookshelf (id, name, pages) VALUES (DEFAULT, 'testbook1', '100')")
	// db.Exec("INSERT INTO bookshelf (id, name, pages) VALUES (DEFAULT, 'testbook2', '200')")
}

func main() {
	if err := godotenv.Load(); err != nil {
		//Do nothing
		log.Fatal("Error loading .env file")
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
