package main

import (
	"github.com/lukemorton/api"
	"gopkg.in/gin-gonic/gin.v1"
)

func App() *gin.Engine {
	app := gin.Default()

	app.GET("/status.json", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "good"})
	})

	app.GET("/authors.json", func(c *gin.Context) {
		c.JSON(200, api.Authors())
	})

	app.NoRoute(func(c *gin.Context) {
		c.JSON(400, api.Error("Bad request, check the docs."))
	})

	return app
}

func main() {
	App().Run(":3000")
}
