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

	r.GET("/", controllers.UserCreate)
	r.Run()
}
