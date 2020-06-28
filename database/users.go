package database

import (
	"database/sql"

	"Sparkle/entities"
	"Sparkle/tools/logger"
)

func (db DB) CreateUser(user entities.User) (int, error) {

	query := "INSERT INTO users(username, password, access_token) VALUES($1, $2, $3) RETURNING id"

	var id int

	err := db.performer().QueryRow(query, user.Username, user.Password, user.AccessToken).Scan(&id)

	if err != nil {
		if db.tx != nil {
			db.tx.Rollback()
		}
		logger.Warning(err)
		return id, err
	}

	return id, nil
}

func (db DB) UpdateUser(user entities.User) error {

	query := "UPDATE users SET username = $2, password = $3, access_token = $4 WHERE id = $1"

	_, err := db.performer().Exec(query, user.Id, user.Username, user.Password, user.AccessToken)

	if err != nil {
		if db.tx != nil {
			db.tx.Rollback()
		}
		logger.Warning(err)
		return err
	}

	return nil
}

func (db DB) DeleteUser(user entities.User) error {

	query := "DELETE FROM users WHERE id = $1"

	_, err := db.performer().Exec(query, user.Id)

	if err != nil {
		if db.tx != nil {
			db.tx.Rollback()
		}
		logger.Warning(err)
		return err
	}

	return nil
}

func (db DB) GetUserByBuilder(builder queryBuilder) (entities.User, error) {

	query := "SELECT id, username, password, access_token FROM users"
	query = builder.formatQuery(query)

	var user entities.User

	err := db.performer().QueryRow(query, builder.params...).Scan(&user.Id, &user.Username, &user.Password, &user.AccessToken)

	if err != nil && err != sql.ErrNoRows {
		if db.tx != nil {
			db.tx.Rollback()
		}
		logger.Warning(err)
		return user, err
	}

	return user, nil
}

func (db DB) GetUsersByBuilder(builder queryBuilder) ([]entities.User, error) {

	query := "SELECT id, username, password, access_token FROM users"
	query = builder.formatQuery(query)

	var users []entities.User

	rows, err := db.performer().Query(query, builder.params...)

	if err != nil && err != sql.ErrNoRows {
		if db.tx != nil {
			db.tx.Rollback()
		}
		logger.Warning(err)
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		var user entities.User
		err := rows.Scan(&user.Id, &user.Username, &user.Password)
		if err != nil {
			if db.tx != nil {
				db.tx.Rollback()
			}
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (db DB) GetUserByAccessToken(token string) (entities.User, error) {
	builder := db.GetBuilder().Equals("access_token", token)
	return db.GetUserByBuilder(builder)
}

func (db DB) GetUserByUsername(username string) (entities.User, error) {
	builder := db.GetBuilder().Equals("username", username)
	return db.GetUserByBuilder(builder)
}
