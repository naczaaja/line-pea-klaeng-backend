package models

type ClientDataRegister struct {
	LineId      string `json:"lineId" xml:"lineId" form:"lineId" query:"lineId"`
	ImageAvatar string `json:"imageAvatar" xml:"imageAvatar" form:"imageAvatar" query:"imageAvatar"`
	IDcard      int    `json:"idCard" xml:"idCard" form:"idCard" query:"idCard"`
}