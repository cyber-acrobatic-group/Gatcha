package photo

import (
	"github.com/cyber-acrobatic-group/Gatcha/docker/go/cmd/app/model"
	"github.com/cyber-acrobatic-group/Gatcha/docker/go/cmd/app/mysql"
)

func Create(photo model.Photo) model.Photo {
	//photo, err := photo.CreatePhoto()
	createdPhoto, err := mysql.CreatePhoto(photo)
	if err != nil {
		panic(err)
	}
	return createdPhoto
}
