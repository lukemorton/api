package main

import (
	"github.com/lukemorton/api/authors"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	AppEngine().Run(":3000")
}

func AppEngine() *gin.Engine {
	a := App{[]string{"Luke Morton", "Bob"}}
	return a.Engine()
}

type App struct {
	AuthorNames []string
}

func (app *App) Engine() *gin.Engine {
	e := gin.Default()

	e.GET("/status.json", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "good"})
	})

	e.GET("/authors.json", func(c *gin.Context) {
		c.JSON(200, api.Authors(app.AuthorNames))
	})

	e.NoRoute(func(c *gin.Context) {
		c.JSON(400, api.Error("Bad request, check the docs."))
	})

	return e
}
