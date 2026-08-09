package main

import (
	"log"
	"net/http"

	_ "project/docs" // after running swag init
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

// @title           Example API
// @version         1.0
// @description     Minimal Gin + Swagger example
// @host            localhost:8000
// @BasePath        /api/v1
// @schemes         http

// TestHandler godoc
// @Summary         ping example
// @Description     do ping
// @Tags            example
// @Accept          json
// @Produce         json
// @Param           user  body  UserCommandDto  true  "send id"
// @Success         200  {string}  string
// @Router          /user/userdetail [post]
func TestHandler(c *gin.Context) {
	var user UserCommandDto
	c.ShouldBindJSON(&user)
	c.JSON(http.StatusOK, user)

}

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
// @Accept          multipart/form-data
// @Produce         json
// @Param           name formData string true  "enter the name"
// @Param           family formData string true  "enter the family"
// @Param file_data formData file true "image"
// @Success         200  {string}  string
// @Router          /user/add [post]
func CreateUserHandler(c *gin.Context) {
	name := c.PostForm("name")
	family := c.PostForm("family")
	fileRecieve, err := c.FormFile("file_data")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	path := "./upload/" + fileRecieve.Filename
	c.SaveUploadedFile(fileRecieve, path)
	c.JSON(http.StatusOK, fileRecieve.Filename+name+family)

}

func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	api := r.Group("/api/v1")
	{
		user := api.Group("/user")
		{
			user.POST("/userdetail", middle.VersionMiddleWare(), TestHandler) // consistent path, no trailing slash
			user.POST("/add", CreateUserHandler)
		}
	}

	log.Println("Starting server on :8000")
	if err := r.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}
