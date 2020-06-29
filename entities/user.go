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

func UsersListToMap(users []User) map[int]User {
	usersMap := make(map[int]User)
	for _, user := range users {
		usersMap[user.Id] = user
	}
	return usersMap
}
