package main

import (
	"github.com/lukemorton/api/users"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	AppEngine().Run(":3000")
}

func AppEngine() *gin.Engine {
	db := users.ConnectDB()
	a := App{db}
	return a.Engine()
}

type App struct {
	DB *users.DB
}

func (app *App) Engine() *gin.Engine {
	e := gin.Default()

	e.GET("/status.json", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "good"})
	})

	e.POST("/register.json", func(c *gin.Context) {
		var user *users.User
		c.BindJSON(&user)
		users.Register(app.DB, *user)
		c.JSON(200, user)
	})

	e.NoRoute(func(c *gin.Context) {
		c.JSON(400, gin.H{"error": "Bad request, check the docs."})
	})

	return e
}
