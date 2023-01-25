package main

import (
	"apitest/pkg/controllers"
	"apitest/pkg/middlewares"
	"apitest/pkg/models"

	_ "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func main() {
	DB := models.ConnectDataBase()
	r := gin.Default()
	public := r.Group("/api")
	// new `GET /users` route associated with our `getUsers` function

	public.POST("/register", controllers.Register(DB))
	public.GET("/login", controllers.Login(DB))
	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser(DB))
	r.Run("localhost:8080")
}
