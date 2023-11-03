package controllers

import (
	"fmt"
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	fmt.Println(body)

	post := models.Post{Title: body.Body, Body: body.Title}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {

	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostsShow(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	initializers.DB.First(&post, id)
	c.JSON(200, gin.H{"post": post})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Body})
	c.JSON(200, gin.H{
		"post": post,
	})

}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	var post models.Post

	initializers.DB.Delete(&post, id)

	c.JSON(204, gin.H{})
}
