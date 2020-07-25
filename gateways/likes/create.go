package likes

import (
	"Sparkle/entities"
	"Sparkle/tools/logger"
)

func (gateway LikesGateway) Create(like entities.Like) (int, error) {

	query := "INSERT INTO likes(user_id, post_id) VALUES($1, $2) RETURNING id"

	var id int

	err := gateway.db.Performer().QueryRow(query, like.UserId, like.PostId).Scan(&id)

	if err != nil {
		if gateway.db.Tx != nil {
			gateway.db.Tx.Rollback()
		}
		logger.Error(err)
		return id, err
	}

	return id, nil
}
