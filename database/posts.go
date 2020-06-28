package database

import (
	"Sparkle/entities"
	"Sparkle/tools/logger"
	"database/sql"
)

func (db DB) CreatePost(post entities.Post) (int, error) {

	query := "INSERT INTO posts(user_id, text, location_code, created_at) VALUES($1, $2, $3, $4) RETURNING id"

	var id int

	err := db.performer().QueryRow(query, post.UserId, post.Text, post.LocationCode, post.CreatedAt).Scan(&id)

	if err != nil {
		if db.tx != nil {
			db.tx.Rollback()
		}
		logger.Warning(err)
		return id, err
	}

	return id, nil
}

func (db DB) UpdatePost(post entities.Post) error {

	query := "UPDATE posts SET user_id = $2, text = $3, location_code = $4, created_at = $5 WHERE id = $1"

	_, err := db.performer().Exec(query, post.Id, post.UserId, post.Text, post.LocationCode, post.CreatedAt)

	if err != nil {
		if db.tx != nil {
			db.tx.Rollback()
		}
		logger.Warning(err)
		return err
	}

	return nil
}

func (db DB) DeletePost(post entities.Post) error {

	query := "DELETE FROM posts WHERE id = $1"

	_, err := db.performer().Exec(query, post.Id)

	if err != nil {
		if db.tx != nil {
			db.tx.Rollback()
		}
		logger.Warning(err)
		return err
	}

	return nil
}

func (db DB) GetPostByBuilder(builder queryBuilder) (entities.Post, error) {

	query := "SELECT id, user_id, text, location_code, created_at FROM posts"
	query = builder.formatQuery(query)

	var post entities.Post

	err := db.performer().QueryRow(query, builder.params...).Scan(&post.Id, &post.UserId, &post.Text, &post.LocationCode, &post.CreatedAt)

	if err != nil && err != sql.ErrNoRows {
		if db.tx != nil {
			db.tx.Rollback()
		}
		logger.Warning(err)
		return post, err
	}

	return post, nil
}

func (db DB) GetPostsByBuilder(builder queryBuilder) ([]entities.Post, error) {

	query := "SELECT id, user_id, text, location_code, created_at FROM posts"
	query = builder.formatQuery(query)

	var posts []entities.Post

	rows, err := db.performer().Query(query, builder.params...)

	if err != nil && err != sql.ErrNoRows {
		if db.tx != nil {
			db.tx.Rollback()
		}
		logger.Warning(err)
		return posts, err
	}

	defer rows.Close()

	for rows.Next() {
		var post entities.Post
		err := rows.Scan(&post.Id, &post.UserId, &post.Text, &post.LocationCode, &post.CreatedAt)
		if err != nil {
			if db.tx != nil {
				db.tx.Rollback()
			}
			return posts, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (db DB) GetPostById(id int) (entities.Post, error) {
	builder := db.GetBuilder().Equals("id", id)
	return db.GetPostByBuilder(builder)
}
