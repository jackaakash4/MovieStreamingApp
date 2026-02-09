package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	controller "github.com/jackaakash4/MovieStreamingApp/Server/CineWorldMoviesServer/controllers"
)

func main() {
	fmt.Println("Hello world")

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	router.GET("/movies", controller.GetMovies())

	if err := router.Run(":8080"); err != nil {

		fmt.Println("Error in server: ", err)
	}
}
