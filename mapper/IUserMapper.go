package mapper

type UserMapper interface {
	QueryUser(username string) UserModel
}
