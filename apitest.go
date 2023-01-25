package main

import (
	"apitest/pkg/controllers"
	"apitest/pkg/graph"
	"apitest/pkg/middlewares"
	"apitest/pkg/models"

	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDataBase()
	r := gin.Default()
	public := r.Group("/api")
	// new `GET /users` route associated with our `getUsers` function

	public.POST("/register", controllers.Register)
	public.GET("/login", controllers.Login)
	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)
	graphqlGr := r.Group("/graph")
	graphqlGr.Use(middlewares.JwtAuthMiddleware())
	graphqlGr.POST("/query", graphqlHandler())
	r.Run("localhost:8080")
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	//resolveroot := new(graph.ResolverRoot)

	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: graph.Resolver{},
	}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
