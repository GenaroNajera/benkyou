package main

import (
	"database/sql"
	"log"

	"github.com/GenaroNajera/benkyou/database"
	"github.com/GenaroNajera/benkyou/handlers"
	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	db, err := sql.Open("sqlite", "./database/benkyou.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	kanji := &database.Kanji{Db: db}
	data := kanji.SelectAll()

	router.GET("/", handlers.Home)
	router.GET("/study", handlers.Study)
	router.GET("/api", func(c *gin.Context) {
		handlers.Api(c, data)
	})

	router.Run("localhost:8080")
}
