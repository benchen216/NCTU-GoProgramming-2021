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
	ID    int    `json:"id,string"`
	NAME  string `json:"name"`
	PAGES int    `json:"pages,string"`
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func return_rows(rows *sql.Rows) []Book {
	var tmp_bookshelf []Book
	for rows.Next() {
		var id int
		var name string
		var pages int
		err := rows.Scan(&id, &name, &pages)
		checkErr(err)

		tmp_book := Book{id, name, pages}
		tmp_bookshelf = append(tmp_bookshelf, tmp_book)
	}

	return tmp_bookshelf
}

func getBooks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// [TODO] get all books data
		rows, err := db.Query("SELECT * FROM bookshelf")
		checkErr(err)

		// [TODO] scan the data one by one
		defer rows.Close()
		return_results := return_rows(rows)
		if len(return_results) == 0 {
			c.String(http.StatusOK, "{ \"message\": \"book not found\" }")
		}

		str := "["
		for i := 0; i < len(return_results); i++ {
			tmp_str := " { \"id\": \"" + strconv.Itoa(return_results[i].ID) + "\", \"name\": \"" + return_results[i].NAME + "\", \"pages\": \"" + strconv.Itoa(return_results[i].PAGES) + "\" }"
			str = str + tmp_str
		}

		str = str + " ]"
		c.String(http.StatusOK, str)
	}
}
func getBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//rows, err := db.Query("SELECT ????? WHERE ??=$1", ???)
		// parse from url
		id := c.Param("id")
		inVar, _ := strconv.Atoi(id)
		rows, err := db.Query("SELECT * FROM bookshelf WHERE id=$1", inVar)
		checkErr(err)

		defer rows.Close()
		return_results := return_rows(rows)
		if len(return_results) == 0 {
			c.String(http.StatusOK, "{ \"message\": \"book not found\" }")
			return
		}

		bookshelf := return_results[0]
		tmp_str := "{ \"id\": \"" + strconv.Itoa(bookshelf.ID) + "\", \"name\": \"" + bookshelf.NAME + "\", \"pages\": \"" + strconv.Itoa(bookshelf.PAGES) + "\" }"
		c.String(http.StatusOK, tmp_str)
	}
}

func addBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//rows, err := db.Query("INSERT INTO ????? VALUES (??,??,??) RETURNING ??", ??, ??)
		var book_from_json Book
		if err := c.ShouldBindJSON(&book_from_json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// // duplicate book
		// rows, err := db.Query("SELECT * FROM bookshelf WHERE id=$1", book_from_json.ID)
		// checkErr(err)
		// defer rows.Close()
		// return_results := return_rows(rows)
		// if len(return_results) != 0 {
		// 	c.String(http.StatusOK, "{ \"message\": \"duplicate book id\" }")
		// 	return
		// }
		// add book

		lastInsertId := 0
		err := db.QueryRow("INSERT INTO bookshelf (name, pages) VALUES ($1, $2) RETURNING id", book_from_json.NAME, book_from_json.PAGES).Scan(&lastInsertId)
		// rows, err = db.Query("INSERT INTO bookshelf (name, pages) VALUES ($1, $2)", book_from_json.NAME, book_from_json.PAGES)
		checkErr(err)

		bookshelf := book_from_json
		tmp_str := "{ \"id\": \"" + strconv.Itoa(lastInsertId) + "\", \"name\": \"" + bookshelf.NAME + "\", \"pages\": \"" + strconv.Itoa(bookshelf.PAGES) + "\" }"
		c.String(http.StatusOK, tmp_str)
	}
}

func updateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//rows, err := db.Query("?????????????????????????", ??, ??, ??)
		var book_from_json Book
		if err := c.ShouldBindJSON(&book_from_json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id := c.Param("id")
		inVar, _ := strconv.Atoi(id)
		// duplicate book
		rows, err := db.Query("SELECT * FROM bookshelf WHERE id=$1", inVar)
		checkErr(err)
		defer rows.Close()
		return_results := return_rows(rows)
		if len(return_results) == 0 {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "book not found"})
			return
		}

		_, err = db.Query("UPDATE bookshelf SET name=$2, pages=$3 WHERE id=$1", inVar, book_from_json.NAME, book_from_json.PAGES)
		checkErr(err)

		bookshelf := book_from_json
		tmp_str := "{ \"id\": \"" + strconv.Itoa(inVar) + "\", \"name\": \"" + bookshelf.NAME + "\", \"pages\": \"" + strconv.Itoa(bookshelf.PAGES) + "\" }"
		c.String(http.StatusOK, tmp_str)
	}
}

func deleteBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//rows, err := db.Query("????????????????????????", ??)

		// parse from url
		id := c.Param("id")
		inVar, _ := strconv.Atoi(id)

		// no book
		rows, err := db.Query("SELECT * FROM bookshelf WHERE ID=$1", inVar)
		checkErr(err)
		defer rows.Close()
		return_results := return_rows(rows)
		if len(return_results) == 0 {
			c.IndentedJSON(http.StatusOK, gin.H{"message": "book not found"})
			return
		}

		_, err = db.Query("DELETE FROM bookshelf WHERE id=$1", inVar)
		checkErr(err)

		bookshelf := return_results[0]
		tmp_str := "{ \"id\": \"" + strconv.Itoa(bookshelf.ID) + "\", \"name\": \"" + bookshelf.NAME + "\", \"pages\": \"" + strconv.Itoa(bookshelf.PAGES) + "\" }"
		c.String(http.StatusOK, tmp_str)
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
