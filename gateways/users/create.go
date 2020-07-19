package users

import (
	"Sparkle/entities"
	"Sparkle/tools/logger"
)

func (gateway UsersGateway) Create(user entities.User) (int, error) {

	query := "INSERT INTO users(username, password, access_token) VALUES($1, $2, $3) RETURNING id"

	var id int

	err := gateway.db.Performer().QueryRow(query, user.Username, user.Password, user.AccessToken).Scan(&id)

	if err != nil {
		if gateway.db.Tx != nil {
			gateway.db.Tx.Rollback()
		}
		logger.Warning(err)
		return id, err
	}

	return id, nil
}
