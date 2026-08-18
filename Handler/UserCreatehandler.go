package handler

import (
	"net/http"

	Database "project/DataBase"
	userdto "project/contract/UserDto"
	"project/models"

	"github.com/gin-gonic/gin"
)

// UserCreateHandler godoc
// @Summary Add new user
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param user body userdto.UserCommandDto true "User information"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/add [post]

func UserCreateHandler(ctx *gin.Context) {
	var userCommand userdto.UserCommandDto

	if err := ctx.ShouldBindJSON(&userCommand); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userObj := models.User{
		Name:   userCommand.Name,
		Family: userCommand.Family,
	}

	if err := Database.DB.Create(&userObj).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, userObj)
}
