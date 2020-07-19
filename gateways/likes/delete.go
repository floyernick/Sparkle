package likes

import (
	"Sparkle/entities"
	"Sparkle/tools/logger"
)

func (gateway LikesGateway) Delete(like entities.Like) error {

	query := "DELETE FROM likes WHERE id = $1"

	_, err := gateway.db.Performer().Exec(query, like.Id)

	if err != nil {
		if gateway.db.Tx != nil {
			gateway.db.Tx.Rollback()
		}
		logger.Warning(err)
		return err
	}

	return nil
}
