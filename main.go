package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rickyguzm/go-jwt/controllers"
	"github.com/rickyguzm/go-jwt/initializers"
	"github.com/rickyguzm/go-jwt/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDataBase()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/logout", controllers.Logout)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()
}
