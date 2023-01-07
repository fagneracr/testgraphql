package main

import (
	"apitest/pkg/controllers"
	"apitest/pkg/models"

	_ "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()
	r := gin.Default()
	public := r.Group("/api")
	// new `GET /users` route associated with our `getUsers` function

	public.POST("/register", controllers.Register)

	r.Run("localhost:8080")
}
