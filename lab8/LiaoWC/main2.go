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
	// write your own struct
	// id's type is int
	ID    int    `json:"id"`
	NAME  string `json:"name"`
	PAGES string `json:"pages"`
}

func getBooks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		/* [TODO] get all books data
		rows, err := db.Query("SELECT ???") */
		rows, err := db.Query("SELECT * FROM bookshelf")
		if err != nil {
			log.Fatal(err)
			//panic("Select error when get all books data.")
		}

		/* [TODO] scan the data one by one
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&?? ,&?? ,&??)
			...
			...
		}*/
		var books []Book
		numRows := 0
		for rows.Next() {
			var book Book
			if rows.Scan(&book.ID, &book.NAME, &book.PAGES) != nil {
				log.Fatal(rows)
				//panic("Scan error when get a book.")
			}
			books = append(books, book)
			numRows++
		}

		//[TODO]send all data or error handling
		if numRows > 0 {
			c.IndentedJSON(http.StatusOK, books)
		} else {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "book not found"})
		}
		return
	}
}
func getBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//rows, err := db.Query("SELECT ????? WHERE ??=$1", ???)
		var book Book
		if err := db.QueryRow("SELECT * FROM bookshelf WHERE id=$1", c.Param("id")).Scan(&book.ID, &book.NAME, &book.PAGES); err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "book not found"})
			return
		}
		c.IndentedJSON(http.StatusOK, book)
		return
	}
}

func addBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//rows, err := db.Query("INSERT INTO ????? VALUES (??,??,??) RETURNING ??", ??, ??)
		var submit Book
		err := c.ShouldBindJSON(&submit)
		if err != nil {
			log.Fatal(err)
		}

		var id int
		err = db.QueryRow("INSERT INTO bookshelf(name, pages) VALUES ($1,$2) RETURNING id", submit.NAME, submit.PAGES).Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "duplicate book id"})
			return
		}
		book := Book{id, submit.NAME, submit.PAGES}
		c.IndentedJSON(http.StatusOK, book)
		return
	}
}

func updateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//rows, err := db.Query("?????????????????????????", ??, ??, ??)
		id := c.Param("id")
		var submit Book
		err := c.ShouldBindJSON(&submit)
		if err != nil {
			log.Fatal(err)
		}

		//db.QueryRow("UPDATE bookshelf SET name=$1, pages=$2 WHERE id=$3", submit.NAME, submit.PAGES, id)
		res, err := db.Exec("UPDATE bookshelf SET name=$1, pages=$2 WHERE id=$3", submit.NAME, submit.PAGES, id)
		if err != nil {
			log.Fatal(err)
		}
		affected, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		if affected != 1 {
			c.JSON(http.StatusOK, gin.H{"message": "book not found"})
			return
		}
		intId, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, Book{intId, submit.NAME, submit.PAGES})
		return
	}
}

func deleteBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//rows, err := db.Query("????????????????????????", ??)
		id := c.Param("id")
		rows, err := db.Query("DELETE FROM bookshelf WHERE id=$1 RETURNING *", id)
		if err != nil {
			log.Fatal(err)
		}
		numDeleted := 0
		var book Book
		for rows.Next() {
			err := rows.Scan(&book.ID, &book.NAME, &book.PAGES)
			if err != nil {
				log.Fatal(err)
			}
			numDeleted++
			break
		}
		if numDeleted == 0 {
			c.JSON(http.StatusOK, gin.H{"message": "book not found"})
		} else {
			c.IndentedJSON(http.StatusOK, book)
		}
		return
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
