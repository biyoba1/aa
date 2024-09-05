package main

import (
	"github.com/biyoba1/redisProject/controllers"
	"github.com/biyoba1/redisProject/initializer"
	"github.com/biyoba1/redisProject/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
	initializer.SyncDatabase()
}

func main() {
	rout := gin.Default()

	rout.POST("/signup", controllers.SignUp)
	rout.POST("/login", controllers.Login)
	rout.GET("/validate", middleware.RequireAuth, controllers.Validate)

	rout.Run()
}
