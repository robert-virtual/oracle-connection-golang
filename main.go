package main

import "github.com/gin-gonic/gin"

func main() {
	connectDB()
	router := gin.Default()
	users := router.Group("/users")
	{
		users.GET("", getUsers)
		users.POST("", postUser)
	}

	router.Run("localhost:8080")

	// defer closeDB()
}
