package handler

import (
	"net/http"

	Database "project/DataBase"
	userdto "project/contract/UserDto"
	"project/models"

	"github.com/gin-gonic/gin"
)

// @title           Example API
// @version         1.0
// @description     Minimal Gin + Swagger example
// @host            localhost:8000
// @BasePath        /api/v1
// @schemes         http

// TestHandler godoc
// @Summary         adding new user
// @Description     add user
// @Tags            User
// @Accept          json
// @Produce         json
// @Param           user body userdto.UserCommandDto true "user"
// @Success         200  {object}  models.User
// @Router          /user/add [post]

func UserCreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var usercommand userdto.UserCommandDto
		if err := ctx.ShouldBindJSON(&usercommand); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		userobj := models.User{
			Name:   usercommand.Name,
			Family: usercommand.Family,
		}

		if err := Database.DB.Create(&userobj).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, userobj)
	}
}
