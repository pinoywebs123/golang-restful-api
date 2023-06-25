package controllers

import (
	"example.com/crud/initializers"
	"example.com/crud/models"
	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {

	var posts []models.Post
	initializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	id := c.Param("id")
	var post models.Post

	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {

	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	id := c.Param("id")

	var post models.Post

	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Body, Body: body.Title})

	c.JSON(200, gin.H{
		"upodate": body,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	initializers.DB.Delete(&post, id)

	c.JSON(200, gin.H{
		"delete": "Record Deleted successfully",
	})

}
