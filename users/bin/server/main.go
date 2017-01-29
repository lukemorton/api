package main

import (
	"github.com/lukemorton/api/users"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
)

func main() {
	HTTP().Run(":3000")
}

func HTTP() *gin.Engine {
	users := users.SQLUserStore()
	users.CreateStore()
	app := app{users}
	return app.Engine()
}

type app struct {
	Store users.UserStore
}

func (app *app) Engine() *gin.Engine {
	http := gin.Default()

	http.GET("/status.json", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "good"})
	})

	http.POST("/register.json", func(c *gin.Context) {
		var r users.RegisterUser
		c.BindJSON(&r)
		user, err := users.Register(app.Store, r)

		if err == nil {
			c.JSON(200, user)
		} else {
			c.JSON(422, gin.H{"error": err.Error()})
		}
	})

	http.POST("/verify.json", func(c *gin.Context) {
		var v users.VerifyUser
		c.BindJSON(&v)
		user, err := users.Verify(app.Store, v)

		if err == nil {
			c.JSON(200, user)
		} else {
			c.JSON(401, gin.H{"error": err.Error()})
		}
	})

	http.POST("/password/reset.json", func(c *gin.Context) {
		var r users.ResetPasswordUser
		c.BindJSON(&r)
		token, err := users.ResetPassword(app.Store, r)

		if err == nil {
			log.Println(gin.H{"token": token})
			c.JSON(200, gin.H{"message": "Reset token has been sent to your email address"})
		} else {
			c.JSON(422, gin.H{"error": err.Error()})
		}
	})

	http.NoRoute(func(c *gin.Context) {
		c.JSON(400, gin.H{"error": "Bad request, check the docs."})
	})

	return http
}
