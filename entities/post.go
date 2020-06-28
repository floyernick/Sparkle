package entities

type Post struct {
	Id           int
	UserId       int
	Text         string
	LocationCode string
	CreatedAt    string
}

func (post Post) Exists() bool {
	return post.Id != 0
}

func (post Post) CreatedBy(user User) bool {
	return post.UserId == user.Id
}
