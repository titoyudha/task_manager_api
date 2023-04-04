package utils

import "github.com/gin-gonic/gin"

func Response(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}
