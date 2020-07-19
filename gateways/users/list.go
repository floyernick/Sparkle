package users

import (
	"Sparkle/entities"
	"Sparkle/gateways/common/database"
	"Sparkle/tools/logger"
	"database/sql"

	"github.com/lib/pq"
)

func (gateway UsersGateway) listByBuilder(builder database.QueryBuilder) ([]entities.User, error) {

	query := "SELECT id, username, password, access_token FROM users"
	query = builder.FormatQuery(query)

	var users []entities.User

	rows, err := gateway.db.Performer().Query(query, builder.Params...)

	if err != nil && err != sql.ErrNoRows {
		if gateway.db.Tx != nil {
			gateway.db.Tx.Rollback()
		}
		logger.Warning(err)
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		var user entities.User
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.AccessToken)
		if err != nil {
			if gateway.db.Tx != nil {
				gateway.db.Tx.Rollback()
			}
			logger.Warning(err)
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (gateway UsersGateway) ListByIds(ids []int) ([]entities.User, error) {
	builder := gateway.db.GetBuilder().And().In("id", pq.Array(ids))
	return gateway.listByBuilder(builder)
}
