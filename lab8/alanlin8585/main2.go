package main

import (
	"database/sql"
	"log"
	"os"
    
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Book struct {
	// write your own struct
	// id's type is int
	id int
	name string
	pages string
}

func getBooks(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		/* [TODO] get all books data
		rows, err := db.Query("SELECT ???") */
        rows, err := db.Query("SELECT * FROM bookshelf")
		/* [TODO] scan the data one by one
		defer rows.Close()
		for rows.Next() {
			rows.Scan(&?? ,&?? ,&??)
			...
			...
		}*/
		bookshelf := []gin.H{}
		/*bookshelf = append(bookshelf, gin.H{
            "id": 1,
            "name": "A",
            "pages": "123",
		})*/
		defer rows.Close()
        for rows.Next() {
            a := Book{}
			rows.Scan(&a.id ,&a.name ,&a.pages)
            bookshelf = append(bookshelf, gin.H{
                "id": a.id,
                "name": a.name,
                "pages": a.pages,
            })
		}
		
		if (len(bookshelf) == 0) {
		    c.IndentedJSON(200, gin.H{
                "message": "book not found",
            })
            return
		}
		
		if (err == nil) {
            c.IndentedJSON(200, bookshelf)
		}
		//[TODO]send all data or error handling
	}
}
func getBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//rows, err := db.Query("SELECT ????? WHERE ??=$1", ???)
		rows, err := db.Query("SELECT * FROM bookshelf WHERE id=$1", c.Param("a"))
		if (err != nil) {
		    c.IndentedJSON(200, gin.H{
                "message": "error",
            })
            return
        }
		for rows.Next() {
            a := Book{}
			rows.Scan(&a.id ,&a.name ,&a.pages)
            c.IndentedJSON(200, gin.H{
                "id": a.id,
                "name": a.name,
                "pages": a.pages,
            })
            return
		}
		c.IndentedJSON(200, gin.H{
            "message": "book not found",
        })
	}
}

func addBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
        jsons := make(map[string]string)
        c.BindJSON(&jsons)
		//rows, err := db.Query("INSERT INTO ????? VALUES (??,??,??) RETURNING ??", ??, ??)
	    rows, err := db.Query("INSERT INTO bookshelf VALUES (DEFAULT,$1,$2) RETURNING *", jsons["NAME"], jsons["PAGES"])
		if (err != nil) {
		    c.IndentedJSON(200, gin.H{
                "message": "error",
            })
            return
        }
	    for rows.Next() {
            a := Book{}
			rows.Scan(&a.id ,&a.name ,&a.pages)
            c.IndentedJSON(200, gin.H{
                "id": a.id,
                "name": a.name,
                "pages": a.pages,
            })
            return
		}
		c.IndentedJSON(200, gin.H{
            "message": "error",
        })
	}
}

func updateBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
        jsons := make(map[string]string)
        c.BindJSON(&jsons)
		//rows, err := db.Query("?????????????????????????", ??, ??, ??)
		rows, err := db.Query("UPDATE bookshelf SET name=$1, pages=$2 WHERE id=$3 RETURNING *", jsons["NAME"], jsons["PAGES"], c.Param("a"))
		if (err != nil) {
		    c.IndentedJSON(200, gin.H{
                "message": "error",
            })
            return
        }
        for rows.Next() {
            a := Book{}
			rows.Scan(&a.id ,&a.name ,&a.pages)
            c.IndentedJSON(200, gin.H{
                "id": a.id,
                "name": a.name,
                "pages": a.pages,
            })
            return
		}
		c.IndentedJSON(200, gin.H{
            "message": "book not found",
        })
	}
}

func deleteBook(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//rows, err := db.Query("????????????????????????", ??)
		rows, err := db.Query("DELETE FROM bookshelf WHERE id=$1 RETURNING *", c.Param("a"))
		if (err != nil) {
		    c.IndentedJSON(200, gin.H{
                "message": "error",
            })
            return
        }
        for rows.Next() {
            a := Book{}
			rows.Scan(&a.id ,&a.name ,&a.pages)
            c.IndentedJSON(200, gin.H{
                "id": a.id,
                "name": a.name,
                "pages": a.pages,
            })
            return
		}
		c.IndentedJSON(200, gin.H{
            "message": "book not found",
        })
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
	r.GET("/bookshelf/:a", getBook(db))
	r.POST("/bookshelf", addBook(db))
	r.DELETE("/bookshelf/:a", deleteBook(db))
	r.PUT("/bookshelf/:a", updateBook(db))

	r.Run(":" + port)
}
