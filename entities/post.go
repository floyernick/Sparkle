package entities

type Post struct {
	Id           int
	UserId       int
	Text         string
	LocationCode string
	CreatedAt    string
	LikesNumber  int
}

func (post Post) Exists() bool {
	return post.Id != 0
}

func (post Post) CreatedBy(user User) bool {
	return post.UserId == user.Id
}

func GetUserIdsFromPosts(posts []Post) []int {
	userIds := make([]int, 0, len(posts))
	for _, post := range posts {
		userIds = append(userIds, post.UserId)
	}
	return userIds
}
