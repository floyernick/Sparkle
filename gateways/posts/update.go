package posts

import (
	"Sparkle/entities"
	"Sparkle/tools/logger"
)

func (gateway PostsGateway) Update(post entities.Post) error {

	query := "UPDATE posts SET user_id = $2, text = $3, location_code = $4, created_at = $5 WHERE id = $1"

	_, err := gateway.db.Performer().Exec(query, post.Id, post.UserId, post.Text, post.LocationCode, post.CreatedAt)

	if err != nil {
		if gateway.db.Tx != nil {
			gateway.db.Tx.Rollback()
		}
		logger.Error(err)
		return err
	}

	return nil
}
