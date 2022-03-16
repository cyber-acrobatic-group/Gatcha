package mysql

import "github.com/cyber-acrobatic-group/Gatcha/docker/go/cmd/app/model"

func CreatePhoto(photo model.Photo) (model.Photo, error) {
	db := GormConnect()
	if err := db.Error; err != nil {
		return model.Photo{}, err
	}
	defer db.Close()
	if err := db.Create(&photo).Error; err != nil {
		return model.Photo{}, err
	}
	return photo, nil
}
