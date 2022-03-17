package photo

import (
	"github.com/cyber-acrobatic-group/Gatcha/docker/go/cmd/app/model"
	"github.com/cyber-acrobatic-group/Gatcha/docker/go/cmd/app/service/photo"
	"github.com/gin-gonic/gin"
)

func CreateHandler(context *gin.Context) model.Photo {
	var form CreateForm
	if err := context.Bind(form); err != nil {
		// TODO:エラーハンドリング実装
		panic(err)
	}
	photoModel := form.convertToModel()
	return photo.Create(photoModel)
}
