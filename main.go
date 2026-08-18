package main

import (
	"log"
	"net/http"

	Database "project/DataBase"
	_ "project/docs"
	handler "project/handler"

	"project/middle"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type UserCommandDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Age      int    `json:"age" binding:"required"`
}

// @title Example API
// @version 1.0
// @description Minimal Gin + Swagger example
// @host localhost:8000
// @BasePath /api/v1
// @schemes http

// TestHandler godoc
// @Summary Test handler
// @Description Echo user
// @Tags User
// @Accept json
// @Produce json
// @Param user body UserCommandDto true "User information"
// @Success 200 {object} UserCommandDto
// @Router /user/userdetail [post]
func TestHandler(c *gin.Context) {
	var user UserCommandDto

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UserCreateWrapper godoc
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
func userCreateWrapper(c *gin.Context) {
	handler.UserCreateHandler(c)
}

// UserListWrapper godoc
// @Summary Get user list
// @Description Get all users
// @Tags User
// @Produce json
// @Success 200 {object} map[string][]models.User
// @Failure 500 {object} map[string]string
// @Router /user/list [get]
func userListWrapper(c *gin.Context) {
	handler.UserListHandler(c)
}

func main() {
	Database.ConDb()

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := r.Group("/api/v1")
	{
		user := api.Group("/user")
		{
			user.POST("/userdetail", middle.VersionMiddleWare(), TestHandler)
			user.POST("/add", userCreateWrapper)
			user.GET("/list", userListWrapper)
		}
	}

	log.Println("Starting server on :8000")

	if err := r.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
