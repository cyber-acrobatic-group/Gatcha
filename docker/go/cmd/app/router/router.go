package router

import (
	"github.com/cyber-acrobatic-group/Gatcha/docker/go/cmd/app/handler"
	"github.com/gin-gonic/gin"
)

// URL設定
func Setting() {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		handler.Hello(ctx)
	})

	router.Run()
}
