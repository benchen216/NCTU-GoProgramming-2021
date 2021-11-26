package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
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
		// [TODO] get all books data
		rows, err := db.Query("SELECT * FROM bookshelf") 
		
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "db error",
			})
			return
		}
		bookshelfResults := make([]Book, 0)
		
		//[TODO] scan the data one by one
		defer rows.Close()
		for rows.Next() {
			book := Book{}			
			rows.Scan(&book.Id ,&book.Name ,&book.Pages)
			bookshelfResults = append(bookshelfResults, book)
		}

		//[TODO]send all data or error handling
		if (len(bookshelfResults) == 0) {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "book not found",
			})
		} else {
			c.IndentedJSON(http.StatusOK, bookshelfResults)
		}
		
	}
}
func getBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var oneBook Book
		id, _ :=  strconv.Atoi(c.Param("id"))
		log.Println(id)
		err := db.QueryRow("SELECT * FROM bookshelf WHERE id=$1", id).Scan(&oneBook.Id, &oneBook.Name, &oneBook.Pages)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "book not found",
			})
			return
		} else {
			c.IndentedJSON(http.StatusOK, oneBook)
		}

	}
}

func addBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var newBook Book
		if err := c.BindJSON(&newBook); err != nil {
			c.IndentedJSON(http.StatusBadRequest, "error binding")
		}
		err := db.QueryRow("INSERT INTO bookshelf VALUES (DEFAULT,$1,$2) RETURNING *", newBook.Name, newBook.Pages).Scan(&newBook.Id, &newBook.Name, &newBook.Pages)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "db error",
			})
			return
		}
		// rows.Scan(&id)
		// log.Println(id)
		c.IndentedJSON(http.StatusOK, newBook)
	}
}

func updateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ :=  strconv.Atoi(c.Param("id"))
		var newBook Book
		if err := c.BindJSON(&newBook); err != nil {
			c.IndentedJSON(http.StatusBadRequest, "error binding")
		}
		err := db.QueryRow("UPDATE bookshelf SET name=$1, pages=$2 WHERE id=$3 RETURNING *", newBook.Name, newBook.Pages, id).Scan(&newBook.Id, &newBook.Name, &newBook.Pages)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "book not found",
			})
			return
		}
		c.IndentedJSON(http.StatusOK, newBook)
	}
}

func deleteBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ :=  strconv.Atoi(c.Param("id"))
		var deletedBook Book
		if err := c.BindJSON(&deletedBook); err != nil {
			c.IndentedJSON(http.StatusBadRequest, "error binding")
		}
		err := db.QueryRow("DELETE FROM bookshelf WHERE id=$1 RETURNING *", id).Scan(&deletedBook.Id, &deletedBook.Name, &deletedBook.Pages)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"message": "book not found",
			})
			return
		}
		c.IndentedJSON(http.StatusOK, deletedBook)

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
	r.PUT("/bookshelf/:id", updateBook(db))
	r.DELETE("/bookshelf/:id", deleteBook(db))
	// [TODO] other method

	r.Run(":" + port)
}
