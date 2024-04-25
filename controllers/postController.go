package controllers

import (
	"fmt"
	"rest-go-with-gin/initialisers"
	"rest-go-with-gin/models"

	"github.com/gin-gonic/gin"
)

func PostCreate(c *gin.Context) {
	// get data from request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// crate a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initialisers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return

	} // returns error

	c.JSON(200, gin.H{
		"message": "Created",
		"post":    post,
	})
}

func PostsUpdate(c *gin.Context) {

	id := c.Param("id")

	// reading the data pof request body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	fmt.Println(body)
	fmt.Println(id)
	// find ht epost we are updating
	var post models.Post
	initialisers.DB.First(&post, id)

	//update it
	initialisers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// respond

	c.JSON(200, gin.H{
		"message": "Created",
		"post":    post,
	})
}

func PostsShow(c *gin.Context) {
	// get id from url
	id := c.Param("id")
	var post models.Post
	initialisers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"message": "retrived",
		"post":    post,
	})
}

func PostsIndex(c *gin.Context) {
	// get the posts
	var posts []models.Post
	initialisers.DB.Find(&posts)
	fmt.Println(posts)

	// response with them

	c.JSON(200, gin.H{
		"message": "All posts response",
		"post":    posts,
	})
}
func PostsDelete(c *gin.Context) {

	id := c.Param("id")

	initialisers.DB.Delete(&models.Post{}, id)

	c.JSON(200, gin.H{
		"message": "Deleted successfully",
	})
}
