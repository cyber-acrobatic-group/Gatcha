package model

import "github.com/jinzhu/gorm"

/*
 gormで構造体Modelが以下の形で定義されているのでそれを利用し、追加したいフィールドを自作する
 参考URL: https://gorm.io/ja_JP/docs/models.html
 type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}
*/
type Photo struct {
	gorm.Model
	// テキスト
	Text string `json:"text"`
	// レアリティ
	Rarity uint `json:"rarity"`
	// 画像パス
	Image_path string `json:"image_path"`
}
