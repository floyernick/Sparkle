package database

import (
	"Sparkle/entities"
	"Sparkle/tools/logger"
	"database/sql"
)

type PostsFilter struct {
	UserIdEquals           int
	LocationCodeStartsWith string
	CreatedAfter           string
	OrderByIdDesc          bool
	Offset                 int
	Limit                  int
}

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

func (db DB) getPostByBuilder(builder queryBuilder) (entities.Post, error) {

	query := "SELECT id, user_id, text, location_code, created_at, (SELECT COUNT(id) FROM likes WHERE post_id = posts.id) AS likes_number FROM posts"
	query = builder.formatQuery(query)

	var post entities.Post

	err := db.performer().QueryRow(query, builder.params...).Scan(&post.Id, &post.UserId, &post.Text, &post.LocationCode, &post.CreatedAt, &post.LikesNumber)

	if err != nil && err != sql.ErrNoRows {
		if db.tx != nil {
			db.tx.Rollback()
		}
		logger.Warning(err)
		return post, err
	}

	return post, nil
}

func (db DB) getPostsByBuilder(builder queryBuilder) ([]entities.Post, error) {

	query := "SELECT id, user_id, text, location_code, created_at, (SELECT COUNT(id) FROM likes WHERE post_id = posts.id) AS likes_number FROM posts"
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
		err := rows.Scan(&post.Id, &post.UserId, &post.Text, &post.LocationCode, &post.CreatedAt, &post.LikesNumber)
		if err != nil {
			if db.tx != nil {
				db.tx.Rollback()
			}
			logger.Warning(err)
			return posts, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (db DB) GetPostById(id int) (entities.Post, error) {
	builder := db.getBuilder().And().Equals("id", id)
	return db.getPostByBuilder(builder)
}

func (db DB) GetPostsByFilter(filter PostsFilter) ([]entities.Post, error) {
	builder := db.getBuilder()
	if filter.UserIdEquals != 0 {
		builder = builder.And().Equals("user_id", filter.UserIdEquals)
	}
	if filter.LocationCodeStartsWith != "" {
		builder = builder.And().StartsWith("location_code", filter.LocationCodeStartsWith)
	}
	if filter.CreatedAfter != "" {
		builder = builder.And().GreaterOrEquals("created_at", filter.CreatedAfter)
	}
	if filter.OrderByIdDesc {
		builder = builder.OrderBy("id", "DESC")
	}
	if filter.Offset != 0 {
		builder = builder.Offset(filter.Offset)
	}
	if filter.Limit != 0 {
		builder = builder.Limit(filter.Limit)
	}
	return db.getPostsByBuilder(builder)
}
