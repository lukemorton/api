package main

import (
	"github.com/lukemorton/api/users"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	AppEngine().Run(":3000")
}

func AppEngine() *gin.Engine {
	users := users.ConnectUserStore()
	users.CreateStore()
	a := App{users}
	return a.Engine()
}

type App struct {
	Store *users.UserStore
}

func (app *App) Engine() *gin.Engine {
	e := gin.Default()

	e.GET("/status.json", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "good"})
	})

	e.POST("/register.json", func(c *gin.Context) {
		var user *users.User
		c.BindJSON(&user)
		err := users.Register(app.Store, user)

		if err == nil {
			c.JSON(200, user)
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
	})

	e.NoRoute(func(c *gin.Context) {
		c.JSON(400, gin.H{"error": "Bad request, check the docs."})
	})

	return e
}
