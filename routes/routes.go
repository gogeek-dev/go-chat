package routes

import (
	"go-realtime-chat/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	go controllers.HandleMessages()
	r.LoadHTMLGlob("templates/*/*.html")
	r.Static("/assets", "./assets")
	r.GET("/", controllers.Index)
	r.GET("/ws", controllers.HandleConnections)

	return r
}
