package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"name": "Eka",
			"bio":  "Software Developer",
		})
	})

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"content":  "judul",
			"subtitle": "subjudul",
		})
	})

	router.Run(":8888")
}
