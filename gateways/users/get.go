package users

import (
	"Sparkle/entities"
	"Sparkle/gateways/common/database"
	"Sparkle/tools/logger"
	"database/sql"
)

func (gateway UsersGateway) getByBuilder(builder database.QueryBuilder) (entities.User, error) {

	query := "SELECT id, username, password, access_token FROM users"
	query = builder.FormatQuery(query)

	var user entities.User

	err := gateway.db.Performer().QueryRow(query, builder.Params...).Scan(&user.Id, &user.Username, &user.Password, &user.AccessToken)

	if err != nil && err != sql.ErrNoRows {
		if gateway.db.Tx != nil {
			gateway.db.Tx.Rollback()
		}
		logger.Error(err)
		return user, err
	}

	return user, nil
}

func (gateway UsersGateway) GetByAccessToken(token string) (entities.User, error) {
	builder := gateway.db.GetBuilder().And().Equals("access_token", token)
	return gateway.getByBuilder(builder)
}

func (gateway UsersGateway) GetById(id int) (entities.User, error) {
	builder := gateway.db.GetBuilder().And().Equals("id", id)
	return gateway.getByBuilder(builder)
}

func (gateway UsersGateway) GetByUsername(username string) (entities.User, error) {
	builder := gateway.db.GetBuilder().And().Equals("username", username)
	return gateway.getByBuilder(builder)
}
