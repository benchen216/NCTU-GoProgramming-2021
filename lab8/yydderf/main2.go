package main

import (
	"database/sql"
	"log"
	"os"
    "net/http"
    "strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Book struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Pages string `json:"pages"`
}

func getBooks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
        var book Book
        var ret []Book
		rows, err := db.Query("SELECT * FROM bookshelf")
        if err != nil {
            panic(err)
        }
		defer rows.Close()
		for rows.Next() {
            err = rows.Scan(&book.ID ,&book.Name ,&book.Pages)
            ret = append(ret, book)
		}
        if len(ret) == 0 {
            c.IndentedJSON(http.StatusOK, gin.H{"message": "book not found"})
            return
        }
        c.IndentedJSON(http.StatusOK, ret)
	}
}
func getBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
        var book Book;
        id, err := strconv.Atoi(c.Param("id"))
		err = db.QueryRow("SELECT * FROM bookshelf WHERE id=$1", id).Scan(&book.ID, &book.Name, &book.Pages)
        if err != nil {
            c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not found"})
            return
        }
        c.IndentedJSON(http.StatusOK, book);
	}
}

func addBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
        var newBook Book;
        if err := c.BindJSON(&newBook); err != nil {
            return
        }

        err := db.QueryRow("INSERT INTO bookshelf VALUES (DEFAULT, $1, $2) RETURNING *", newBook.Name, newBook.Pages).Scan(&newBook.ID, &newBook.Name, &newBook.Pages)
        if err != nil {
            c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error"})
        }
        c.IndentedJSON(http.StatusOK, newBook)
	}
}

func updateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
        var newBook Book
        if err := c.BindJSON(&newBook); err != nil {
            panic(err)
            return
        }
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            panic(err)
            return
        }
        err = db.QueryRow("UPDATE bookshelf SET (name, pages) = ($1, $2) WHERE id = $3 RETURNING id", newBook.Name, newBook.Pages, id).Scan(&newBook.ID)
        if err != nil {
            c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
            return
        }
        c.IndentedJSON(http.StatusOK, newBook)
	}
}

func deleteBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
        var book Book
        id, err := strconv.Atoi(c.Param("id"))
        if err != nil {
            panic(err)
            return
        }
		err = db.QueryRow("DELETE FROM bookshelf WHERE id = $1 RETURNING *", id).Scan(&book.ID, &book.Name, &book.Pages)
        if err != nil {
            c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
            return
        }
        c.IndentedJSON(http.StatusOK, book)
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
        return
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

	r.Run(":" + port)
}
