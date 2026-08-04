package main

import (
	"log"
	"net/http"

	_ "project/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /user/getlist [get]
func TestHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "I am test handler")
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /user/create [post]
func UserCreateHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "I am test handler")
}

func main() {

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	api := r.Group("/api/v1")
	{
		user := api.Group("/user")
		{
			user.GET("/getlist", TestHandler)
			user.POST("/create", TestHandler)

		}

	}
	log.Println("Starting server")
	r.Run(":8000")

}
