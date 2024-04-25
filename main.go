package main

import (
	"rest-go-with-gin/controllers"
	"rest-go-with-gin/initialisers"

	"github.com/gin-gonic/gin"
)

func init() {
	initialisers.LoadEnvVariables()
	initialisers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/posts", controllers.PostCreate)        //done
	r.PUT("/posts/:id", controllers.PostsUpdate)    //done
	r.GET("/posts", controllers.PostsIndex)         //done
	r.GET("/posts/:id", controllers.PostsShow)      //done
	r.DELETE("/posts/:id", controllers.PostsDelete) //done

	r.Run() // listen and serve on 0.0.0.0:8080
}
