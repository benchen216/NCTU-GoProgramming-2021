package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

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

func callerr(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"message": "book not found",
	})
}

func parse(input string) (string, string) {
	var name, pages string
	temp := input[1 : len(input)-1]
	token := strings.Split(temp, ",")

	for index, value := range token {
		if index == 0 {
			tmp := strings.Split(value, ":")
			name = tmp[1][1:(len(tmp[1]) - 1)]
		} else if index == 1 {
			tmp := strings.Split(value, ":")
			pages = tmp[1][1:(len(tmp[1]) - 1)]
		} else {
			fmt.Println("Error input")
		}

	}
	return name, pages
}

func getBooks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		/* [TODO] get all books data
		rows, err := db.Query("SELECT ???") */
		var id int
		var name, pages string
		var books []Book
		rows, err := db.Query("SELECT * FROM bookshelf")
		if err != nil {
			log.Fatal(err)
			return
		}
		defer rows.Close()
		for rows.Next() {
			err = rows.Scan(&id, &name, &pages)
			if err != nil {
				log.Fatal(err)
				return
			}
			temp := Book{id, name, pages}
			books = append(books, temp)
		}
		if len(books) == 0 {
			callerr(c)
			return
		}
		c.IndentedJSON(200, books)

		/* [TODO] scan the data one by one
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&?? ,&?? ,&??)
			...
			...
		}*/

		//[TODO]send all data or error handling
	}
}
func getBook(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		ids := c.Param("id")
		var name, pages string
		//rows, err := db.Query("SELECT ????? WHERE ??=$1", ???)
		rows := db.QueryRow("SELECT name,pages FROM bookshelf WHERE id = $1", ids)
		err := rows.Scan(&name, &pages)
		if err != nil {
			callerr(c)
			return
		}
		var books []Book
		id, _ := strconv.Atoi(ids)
		temp := Book{id, name, pages}
		books = append(books, temp)

		if len(books) == 0 {
			callerr(c)
			return
		}
		c.IndentedJSON(200, temp)
	}
}

func addBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//rows, err := db.Query("INSERT INTO ????? VALUES (??,??,??) RETURNING ??", ??, ??)
		body := c.Request.Body
		value, err := ioutil.ReadAll(body)
		if err != nil {
			c.JSON(200, err.Error())
		}
		var id int
		name, pages := parse(string(value))
		fmt.Println(name, pages)
		errr := db.QueryRow("INSERT INTO bookshelf VALUES (DEFAULT,$1,$2) RETURNING id", name, pages).Scan(&id)

		if errr != nil {
			fmt.Println(errr)
			callerr(c)
			return
		}
		temp := Book{id, name, pages}
		c.IndentedJSON(200, temp)

	}
}

func updateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//rows, err := db.Query("?????????????????????????", ??, ??, ??)
		ids := c.Param("id")
		body := c.Request.Body
		value, err := ioutil.ReadAll(body)
		if err != nil {
			log.Fatal(err)
		}
		name, pages := parse(string(value))
		rows := db.QueryRow("UPDATE bookshelf SET name=$2,pages=$3 WHERE id = $1 RETURNING *", ids, name, pages)

		var i int
		var n, p string
		err = rows.Scan(&i, &n, &p)
		if err != nil {
			callerr(c)
			return
		}
		t := Book{i, n, p}
		c.IndentedJSON(200, t)

	}
}

func deleteBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//rows, err := db.Query("????????????????????????", ??)
		ids := c.Param("id")
		row := db.QueryRow("SELECT * FROM bookshelf WHERE id =$1", ids)
		var id int
		var name, pages string
		err := row.Scan(&id, &name, &pages)
		if err != nil {
			callerr(c)
			return
		}
		rows, err := db.Query("DELETE FROM bookshelf WHERE id = $1", ids)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer rows.Close()
		t := Book{id, name, pages}
		c.IndentedJSON(200, t)
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
