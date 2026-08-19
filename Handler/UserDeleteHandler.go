package handler

import (
	"net/http"
	"strconv"

	Database "project/DataBase"
	"project/models"

	"github.com/gin-gonic/gin"
)

// UserDeleteHandler godoc
// @Summary Delete user
// @Description Delete a user by ID
// @Tags User
// @Produce json
// @Param id path int true "ID of target user"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user/delete/{id} [delete]
func UserDeleteHandler(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

	var user models.User

	result := Database.DB.First(&user, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	result = Database.DB.Delete(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to delete user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
		"id":      id,
	})
}
