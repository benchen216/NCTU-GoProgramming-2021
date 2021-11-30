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
	ID int `json:"id"`
	NAME string `json:"name"`
	PAGES string `json:"pages"`
}

func getBooks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		/* [TODO] get all books data */
		rows, _ := db.Query("SELECT * FROM bookshelf")
		/* [TODO] scan the data one by one */
		defer rows.Close()
		bookshelf := []Book{}
		for rows.Next() {
			new_book := Book{}
			rows.Scan(&new_book.ID ,&new_book.NAME ,&new_book.PAGES)
			bookshelf = append(bookshelf,new_book)
		}
		//[TODO]send all data or error handling
		if len(bookshelf)==0{
			c.IndentedJSON(http.StatusOK, gin.H{"message": "book not found"})
		}else{
			c.IndentedJSON(http.StatusOK, bookshelf)
		}
	}
}
func getBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		i := c.Param("index")
		a_book := Book{}
		err := db.QueryRow("SELECT * FROM bookshelf WHERE id=$1", i).Scan(&a_book.ID,&a_book.NAME,&a_book.PAGES)
		if err != nil{
			c.IndentedJSON(http.StatusOK, gin.H{"message": "book not found"})
		}else{
			c.IndentedJSON(http.StatusOK, a_book)
		}
	}
}

func addBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var json Book
		if err2 := c.ShouldBindJSON(&json); err2 != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
			return
		}
		err := db.QueryRow("INSERT INTO bookshelf VALUES (DEFAULT,$1,$2) RETURNING *", json.NAME, json.PAGES).Scan(&json.ID, &json.NAME, &json.PAGES)
		if err != nil{
			c.IndentedJSON(http.StatusOK, gin.H{"message": "error"})
		}else{
			c.IndentedJSON(http.StatusOK, json)
		}
	}
}

func updateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		i := c.Param("index")
		var json Book
		if err2 := c.ShouldBindJSON(&json); err2 != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
			return
		}
		err := db.QueryRow("UPDATE bookshelf SET name=$1, pages=$2 WHERE id=$3 RETURNING *", json.NAME, json.PAGES, i).Scan(&json.ID, &json.NAME, &json.PAGES)
		if err != nil{
			c.IndentedJSON(http.StatusOK, gin.H{"message": "book not found"})
		}else{
			c.IndentedJSON(http.StatusOK, json)
		}
	}
}

func deleteBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		i := c.Param("index")
		var json Book
		err := db.QueryRow("DELETE FROM bookshelf WHERE id=$1 RETURNING *", i).Scan(&json.ID, &json.NAME, &json.PAGES)
		if err != nil{
			c.IndentedJSON(http.StatusOK, gin.H{"message": "book not found"})
		}else{
			c.IndentedJSON(http.StatusOK, json)
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
	r.GET("/bookshelf/:index", getBook(db))
	r.POST("/bookshelf", addBook(db))
	r.DELETE("/bookshelf/:index", deleteBook(db))
	r.PUT("/bookshelf/:index", updateBook(db))
	r.Run(":" + port)
}
