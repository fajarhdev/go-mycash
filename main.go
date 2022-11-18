package main

import (
	"github.com/fajarhdev/go-mycash/controllers"
	"github.com/fajarhdev/go-mycash/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	r := gin.Default()

	// user
	r.POST("/signup", controllers.UserCreate)
	r.GET("/login", controllers.GetUser)
	
	r.Run()
}
