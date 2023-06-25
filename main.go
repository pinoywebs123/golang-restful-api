package main

import (
	"example.com/crud/controllers"
	"example.com/crud/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()

}
func main() {
	r := gin.Default()

	r.GET("/", controllers.PostIndex)
	r.POST("/post", controllers.PostCreate)
	r.GET("/post/:id", controllers.PostShow)
	r.PUT("/post/:id", controllers.PostUpdate)
	r.DELETE("/post/:id", controllers.PostDelete)

	r.Run()
}
