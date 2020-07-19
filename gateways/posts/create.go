package posts

import (
	"Sparkle/entities"
	"Sparkle/tools/logger"
)

func (gateway PostsGateway) Create(post entities.Post) (int, error) {

	query := "INSERT INTO posts(user_id, text, location_code, created_at) VALUES($1, $2, $3, $4) RETURNING id"

	var id int

	err := gateway.db.Performer().QueryRow(query, post.UserId, post.Text, post.LocationCode, post.CreatedAt).Scan(&id)

	if err != nil {
		if gateway.db.Tx != nil {
			gateway.db.Tx.Rollback()
		}
		logger.Warning(err)
		return id, err
	}

	return id, nil
}
