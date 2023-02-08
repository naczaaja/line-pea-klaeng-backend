package models

import "gorm.io/gorm"

type UserDB struct {
	gorm.Model
	LineId      string `json:"lineId" xml:"lineId" form:"lineId" query:"lineId"`
	ImageAvatar string `json:"imageAvatar" xml:"imageAvatar" form:"imageAvatar" query:"imageAvatar"`
	IDcard      int    `json:"idCard" xml:"idCard" form:"idCard" query:"idCard"`
}