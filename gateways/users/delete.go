package users

import (
	"Sparkle/entities"
	"Sparkle/tools/logger"
)

func (gateway UsersGateway) Delete(user entities.User) error {

	query := "DELETE FROM users WHERE id = $1"

	_, err := gateway.db.Performer().Exec(query, user.Id)

	if err != nil {
		if gateway.db.Tx != nil {
			gateway.db.Tx.Rollback()
		}
		logger.Warning(err)
		return err
	}

	return nil
}
