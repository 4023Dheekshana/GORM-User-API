package main

import (
	"userapi_GORM/database"
	"userapi_GORM/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.Init()

	router.GET("/user", handler.GetUsers)
	router.POST("/user", handler.AddUser)
	router.GET("/userbyId/:id", handler.GetUserByID)
	router.DELETE("/userbyId/:id", handler.DeleteUserById)

	router.Run(":8080")
}
