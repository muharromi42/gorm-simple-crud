package main

import (
	"github.com/gin-gonic/gin"
	"github.com/muharromi42/gorm-simple-crud.git/controllers"
	"github.com/muharromi42/gorm-simple-crud.git/models"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.POST("/posts", controllers.CreatePost)

	router.Run(":3050")
}
