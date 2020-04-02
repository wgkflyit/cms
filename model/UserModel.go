package model

type UserModel struct {
	Username   string  `json:"username"`
	Password   string  `json:"password"`
	Age        int  `json:"age"`
	Sex       bool  `json:"sex"`
	CardNumber string `json:"cardNumber"`
}
