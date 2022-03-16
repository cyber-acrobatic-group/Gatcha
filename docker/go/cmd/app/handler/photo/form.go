package photo

import "github.com/cyber-acrobatic-group/Gatcha/docker/go/cmd/app/model"

type CreateForm struct {
	// テキスト
	text string `json:"text"`
	// レアリティ
	rarity uint `json:"rarity"`
	// 画像パス
	image_path string `json:"image_path"`
}

func (form CreateForm) convertToModel() model.Photo {
	return model.Photo{
		Text:       form.text,
		Rarity:     form.rarity,
		Image_path: form.image_path,
	}
}
