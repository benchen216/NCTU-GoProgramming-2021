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
	PAGE string `json:"pages"`
}

func getBooks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		/* [TODO] get all books data
		rows, err := db.Query("SELECT ???") */
		rows, err := db.Query("SELECT * FROM bookshelf")
		if err != nil {
			log.Fatalf("%q", err)
			return
		}

		/* [TODO] scan the data one by one
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&?? ,&?? ,&??)
			...
			...
		}*/
		defer rows.Close()
		
		var bookshelf []Book
		for rows.Next() {
			var book Book
			if err := rows.Scan(&book.ID, &book.NAME, &book.PAGE); err != nil {
				log.Fatalf("%q", err)
				return
			}
			
			bookshelf = append(bookshelf, book)
		}
		
		if err = rows.Err(); err != nil {
			log.Fatalf("%q", err)
			return
		}

		//[TODO]send all data or error handling
		if len(bookshelf) == 0 {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		} else {
			c.IndentedJSON(http.StatusOK, bookshelf)
		}
	}
}
func getBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//rows, err := db.Query("SELECT ????? WHERE ??=$1", ???)
		var book Book
		if err := db.QueryRow("SELECT * FROM bookshelf WHERE id=$1", 
			c.Param("id")).Scan(&book.ID, &book.NAME, &book.PAGE); err != nil {
			if err == sql.ErrNoRows {
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
			} else {
				log.Fatalf("%q", err)				
			}
			return
		}
		
		c.IndentedJSON(http.StatusOK, book)
	}
}

func addBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//rows, err := db.Query("INSERT INTO ????? VALUES (??,??,??) RETURNING ??", ??, ??)
		var newBook Book
		if err := c.BindJSON(&newBook); err != nil {
			return
		}
		
		err := db.QueryRow("INSERT INTO bookshelf VALUES(DEFAULT, $1, $2) RETURNING id", 
			newBook.NAME, newBook.PAGE).Scan(&newBook.ID)
		
		if err != nil {
			log.Fatalf("%q", err)
			return
		}
		
		c.IndentedJSON(http.StatusOK, newBook)
	}
}

func updateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//rows, err := db.Query("?????????????????????????", ??, ??, ??)
		var book Book
		if err := c.BindJSON(&book); err != nil {
			return
		}
		
		id := -1
		err := db.QueryRow("UPDATE bookshelf SET name = $2, pages = $3 WHERE id = $1 RETURNING id",
			c.Param("id"), book.NAME, book.PAGE).Scan(&id)
		
		if err != nil {
			if err == sql.ErrNoRows {
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
			} else {
				log.Fatalf("%q", err)
			}
			return
		}

		book.ID = id
		c.IndentedJSON(http.StatusOK, book)
		
	}
}

func deleteBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//rows, err := db.Query("????????????????????????", ??)
		var book Book
		
		err := db.QueryRow("DELETE FROM bookshelf WHERE id = $1 RETURNING *", 
			c.Param("id")).Scan(&book.ID, &book.NAME, &book.PAGE)
		
		if err != nil {
			if err == sql.ErrNoRows {
				c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
			} else {
				log.Fatalf("%q", err)
			}
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
