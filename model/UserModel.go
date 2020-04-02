package model


//UserModel 是用来进行用户基本信息的实体类
type UserModel struct {
	Username   string  `json:"username"`
	Password   string  `json:"password"`
	LoginName string `json:"loginname"`
	Age        int  `json:"age"`
	Sex       bool  `json:"sex"`
	CardNumber string `json:"cardNumber"`
}

//CheckUser 校验用户的用户名密码是否一直
func (um *UserModel)CheckUser(loginname,password string)bool{
	if loginname == um.Loginname && password == um.Password {
		return true
	}
	return false
}
