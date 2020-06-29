package entities

type Like struct {
	Id     int
	UserId int
	PostId int
}

func (like Like) Exists() bool {
	return like.Id != 0
}
