package handler

import (
	"net/http"

	Database "project/DataBase"
	"project/models"

	"github.com/gin-gonic/gin"
)

// UserListHandler godoc
// @Summary Get user list
// @Description Get all users
// @Tags User
// @Produce json
// @Success 200 {object} map[string][]models.User
// @Failure 500 {object} map[string]string
// @Router /user/list [get]

func UserListHandler(c *gin.Context) {
	var users []models.User

	if err := Database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
