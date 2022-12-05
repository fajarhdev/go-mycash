package routes

import (
	"github.com/fajarhdev/go-mycash/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine{
	r := gin.Default()
	user := r.Group("/user")
	{
		user.GET("/", controllers.GetUsers)			// for fetch all user
		user.POST("/", controllers.CreateUser)		// for register
		user.POST("/login", controllers.FindUser)	// for login
		user.GET("/:id", controllers.GetUserByID)	// for select spesific user with id
		user.PUT("/:id", controllers.UpdateUser)	// for update user 
		user.DELETE("/:id", controllers.DeleteUser)	// fpr deletting user
	}
	
	income := r.Group("/income")
	{
		income.POST("/", controllers.PostIncome)	// for add income transaction
		income.GET("/", controllers.GetIncome)		// for fetch all the income transaction
	}

	expense := r.Group("/expense")
	{
		expense.POST("/", controllers.AddExpense)
		
	}
	return r
}