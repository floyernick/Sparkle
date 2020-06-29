package database

import (
	"Sparkle/entities"
	"Sparkle/tools/logger"
	"database/sql"
)

func (db DB) CreateLike(like entities.Like) (int, error) {

	query := "INSERT INTO likes(user_id, post_id) VALUES($1, $2) RETURNING id"

	var id int

	err := db.performer().QueryRow(query, like.UserId, like.PostId).Scan(&id)

	if err != nil {
		if db.tx != nil {
			db.tx.Rollback()
		}
		logger.Warning(err)
		return id, err
	}

	return id, nil
}

func (db DB) DeleteLike(like entities.Like) error {

	query := "DELETE FROM likes WHERE id = $1"

	_, err := db.performer().Exec(query, like.Id)

	if err != nil {
		if db.tx != nil {
			db.tx.Rollback()
		}
		logger.Warning(err)
		return err
	}

	return nil
}

func (db DB) GetLikeByBuilder(builder queryBuilder) (entities.Like, error) {

	query := "SELECT id, user_id, post_id FROM likes"
	query = builder.formatQuery(query)

	var like entities.Like

	err := db.performer().QueryRow(query, builder.params...).Scan(&like.Id, &like.UserId, &like.PostId)

	if err != nil && err != sql.ErrNoRows {
		if db.tx != nil {
			db.tx.Rollback()
		}
		logger.Warning(err)
		return like, err
	}

	return like, nil
}

func (db DB) GetLikesByBuilder(builder queryBuilder) ([]entities.Like, error) {

	query := "SELECT id, user_id, post_id FROM likes"
	query = builder.formatQuery(query)

	var likes []entities.Like

	rows, err := db.performer().Query(query, builder.params...)

	if err != nil && err != sql.ErrNoRows {
		if db.tx != nil {
			db.tx.Rollback()
		}
		logger.Warning(err)
		return likes, err
	}

	defer rows.Close()

	for rows.Next() {
		var like entities.Like
		err := rows.Scan(&like.Id, &like.UserId, &like.PostId)
		if err != nil {
			if db.tx != nil {
				db.tx.Rollback()
			}
			logger.Warning(err)
			return likes, err
		}
		likes = append(likes, like)
	}

	return likes, nil
}

func (db DB) GetLikeByUserIdAndPostId(userId int, postId int) (entities.Like, error) {
	builder := db.GetBuilder().Equals("post_id", postId).And().Equals("user_id", userId)
	return db.GetLikeByBuilder(builder)
}
