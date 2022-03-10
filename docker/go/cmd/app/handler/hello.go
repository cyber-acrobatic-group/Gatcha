package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/cyber-acrobatic-group/Gatcha/docker/go/cmd/app/service"

// 処理を呼び出してレスポンスを返す関数を書く（処理のロジックはserviceパッケージ内で書く）
func Hello(ctx *gin.Context) {
	service.GetMessage(ctx)
}
