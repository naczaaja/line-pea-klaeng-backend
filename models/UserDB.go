package models

import "gorm.io/gorm"

type UserDB struct {
	gorm.Model
	LineId     string `json:"lineId"`
	PrefixName string `json:"prefix_name" gorm:"text;not null;default:null"`
	Firstname  string `json:"firstname" gorm:"text;not null;default:null"`
	Lastname   string `json:"lastname" gorm:"text;not null;default:null"`
	IDcard     int    `json:"id_card"`
}
