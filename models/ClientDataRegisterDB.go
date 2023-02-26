package models

import "gorm.io/gorm"

type ClientDataRegisterDB struct {
	gorm.Model
	LineId      string `json:"lineId"`
	ImageAvatar string `json:"image_avatar"`
	IDcard      int    `json:"id_card"`
}
