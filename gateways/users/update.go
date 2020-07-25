package users

import (
	"Sparkle/entities"
	"Sparkle/tools/logger"
)

func (gateway UsersGateway) Update(user entities.User) error {

	query := "UPDATE users SET username = $2, password = $3, access_token = $4 WHERE id = $1"

	_, err := gateway.db.Performer().Exec(query, user.Id, user.Username, user.Password, user.AccessToken)

	if err != nil {
		if gateway.db.Tx != nil {
			gateway.db.Tx.Rollback()
		}
		logger.Error(err)
		return err
	}

	return nil
}
