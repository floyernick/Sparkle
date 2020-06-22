package entities

type User struct {
	Id          int
	Username    string
	Password    string
	AccessToken string
}

func (user User) Exists() bool {
	return user.Id != 0
}
