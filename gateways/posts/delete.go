package posts

import (
	"Sparkle/entities"
	"Sparkle/tools/logger"
)

func (gateway PostsGateway) Delete(post entities.Post) error {

	query := "DELETE FROM posts WHERE id = $1"

	_, err := gateway.db.Performer().Exec(query, post.Id)

	if err != nil {
		if gateway.db.Tx != nil {
			gateway.db.Tx.Rollback()
		}
		logger.Error(err)
		return err
	}

	return nil
}
