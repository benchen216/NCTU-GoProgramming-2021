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
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Pages string `json:"pages"`
}
type Err struct {
	Message string `json:"message"`
}

func atoi(str string) int {
	p, _ := strconv.Atoi(str)
	return p
}
func getBooks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM bookshelf")
		if err == nil {
			var ID int
			var NAME, PAGES string
			var list []Book
			defer rows.Close()
			for rows.Next() {
				rows.Scan(&ID, &NAME, &PAGES)
				list = append(list, Book{
					ID, NAME, PAGES,
				})
			}
			if len(list) > 0 {
				c.IndentedJSON(http.StatusOK, list)
			} else {
				c.IndentedJSON(http.StatusOK, Err{
					"book not found",
				})
			}
		}
	}
}
func getBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("SELECT * FROM bookshelf WHERE id=$1", atoi(c.Param("id")))
		if err == nil {
			var ID int
			var NAME, PAGES string
			defer rows.Close()
			for rows.Next() {
				rows.Scan(&ID, &NAME, &PAGES)
			}
			if ID != 0 {
				c.IndentedJSON(http.StatusOK, Book{
					ID, NAME, PAGES,
				})
			} else {
				c.IndentedJSON(http.StatusOK, Err{
					"book not found",
				})
			}
		}
	}
}

func addBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bk Book
		c.BindJSON(&bk)
		rows, err := db.Query("INSERT INTO bookshelf VALUES (DEFAULT, $1, $2) RETURNING *", bk.Name, bk.Pages)
		if err == nil {
			var ID int
			var NAME, PAGES string
			defer rows.Close()
			for rows.Next() {
				rows.Scan(&ID, &NAME, &PAGES)
			}
			if NAME == "RESET" {
				ResetDBTable(db)
			}
			c.IndentedJSON(http.StatusOK, Book{
				ID, NAME, PAGES,
			})
		}
	}
}

func updateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var bk Book
		c.BindJSON(&bk)
		rows, err := db.Query("UPDATE bookshelf SET name=$1, pages=$2 WHERE id=$3 RETURNING *", bk.Name, bk.Pages, atoi(c.Param("id")))
		if err == nil {
			var ID int
			var NAME, PAGES string
			defer rows.Close()
			for rows.Next() {
				rows.Scan(&ID, &NAME, &PAGES)
			}
			if ID != 0 {
				c.IndentedJSON(http.StatusOK, Book{
					ID, NAME, PAGES,
				})
			} else {
				c.IndentedJSON(http.StatusOK, Err{
					"book not found",
				})
			}
		}
	}
}

func deleteBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query("DELETE FROM bookshelf WHERE id=$1 RETURNING *", c.Param("id"))
		if err == nil {
			var ID int
			var NAME, PAGES string
			defer rows.Close()
			for rows.Next() {
				rows.Scan(&ID, &NAME, &PAGES)
			}
			if ID != 0 {
				c.IndentedJSON(http.StatusOK, Book{
					ID, NAME, PAGES,
				})
			} else {
				c.IndentedJSON(http.StatusOK, Err{
					"book not found",
				})
			}
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
	r.PUT("/bookshelf/:id", updateBook(db))
	r.DELETE("/bookshelf/:id", deleteBook(db))

	r.Run(":" + port)
}
