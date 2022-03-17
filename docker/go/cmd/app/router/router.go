package router

import (
	"net/http"

	"github.com/cyber-acrobatic-group/Gatcha/docker/go/cmd/app/handler/photo"
	"github.com/gin-gonic/gin"
)

func Setting() {
	router := gin.Default()
	router.POST("/photo", func(context *gin.Context) {
		photo := photo.CreateHandler(context)
		context.JSON(http.StatusOK, photo)
	})
	router.Run(":3000")
}
