package handler

import (
	"log"
	"net/http"
	"userapi_GORM/database"

	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {
	var userinfo []database.Users
	database.Db.Find(&userinfo)
	context.JSON(http.StatusOK, userinfo)
}

func AddUser(context *gin.Context) {
	var user database.Users
	if err := context.BindJSON(&user); err != nil {
		log.Fatal(err)
	}
	database.Db.Create(&user)
	context.JSON(http.StatusOK, user)
}

func GetUserByID(context *gin.Context) {
	var user database.Users
	id := context.Param("id")
	database.Db.First(&user, id)
	context.JSON(http.StatusOK, user)
}

func DeleteUserById(context *gin.Context) {
	var user database.Users
	id := context.Param("id")
	database.Db.Delete(&user, id)
	context.JSON(http.StatusOK, id)
}
