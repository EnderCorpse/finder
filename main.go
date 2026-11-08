package main

import (
	"net/http"

	"github.com/ComputerScienceHouse/finder/internal/logging"
	"github.com/gin-gonic/gin"
)

func main() {
	logging.Logger.Info("Hello World!")

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
		return
	})

	router.Run(":8080")
}
