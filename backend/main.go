package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Static("/static/css", "frontend/build/static/css")
	r.Static("/static/js", "frontend/build/static/js")

	r.LoadHTMLFiles("frontend/build/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	_ = r.Run(":8080")
}
