package main

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/ComputerScienceHouse/finder/internal/database"
	"github.com/ComputerScienceHouse/finder/internal/logging"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	if err := database.InitDatabase(); err != nil {
		logging.Logger.Fatalf("Error initializing database: %s", err)
	}
	//db, err := sql.Open("postgres", "host=postgres.csh.rit.edu port=5432 user=finder dbname=finder sslmode=disable")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer db.Close()
	//
	//// Verify the connection is alive
	//if err := db.Ping(); err != nil {
	//	log.Fatal(err)
	//}
	logging.Logger.Info("Hello World!")

	router := gin.Default()

	//router.GET("/users/:id", func(c *gin.Context) {
	//	var name string
	//	err := db.QueryRowContext(c.Request.Context(), "SELECT name FROM users WHERE id = $1", c.Param("id")).Scan(&name)
	//	if err == sql.ErrNoRows {
	//		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
	//		return
	//	}
	//}

	router.GET("/items", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
		defer cancel()

		items, err := database.GetAllItems(ctx)

		if err == sql.ErrNoRows {
			c.Status(http.StatusNoContent)
			return
		} else if err != nil {
			c.Status(http.StatusInternalServerError)
		}

		logging.Logger.Infof("Items: %v", items)

		c.JSON(http.StatusOK, items)
	})

	router.LoadHTMLGlob("templates/*")

	router.GET("t", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
		return
	})

	router.Run(":3000")
}
