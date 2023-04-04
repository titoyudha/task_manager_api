package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/titoyudha/task_manager_api/internal/config"
)

func Run() {
	router := gin.Default()

	config.ConnectDB()

	router.GET("/home", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Task Manager API",
		})
	})

	router.Run(":8000")
}
