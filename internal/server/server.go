package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.GET("/home", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Task Manager API",
		})
	})

	router.Run(":8000")
}
