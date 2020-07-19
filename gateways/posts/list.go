package posts

import (
	"Sparkle/entities"
	"Sparkle/gateways/common/database"
	"Sparkle/tools/logger"
	"database/sql"
)

type ListFilter struct {
	UserIdEquals           int
	LocationCodeStartsWith string
	CreatedAfter           string
	OrderByCreatedAtDesc   bool
	Offset                 int
	Limit                  int
}

func (gateway PostsGateway) listByBuilder(builder database.QueryBuilder) ([]entities.Post, error) {

	query := "SELECT id, user_id, text, location_code, created_at, (SELECT COUNT(id) FROM likes WHERE post_id = posts.id) AS likes_number FROM posts"
	query = builder.FormatQuery(query)

	var posts []entities.Post

	rows, err := gateway.db.Performer().Query(query, builder.Params...)

	if err != nil && err != sql.ErrNoRows {
		if gateway.db.Tx != nil {
			gateway.db.Tx.Rollback()
		}
		logger.Warning(err)
		return posts, err
	}

	defer rows.Close()

	for rows.Next() {
		var post entities.Post
		err := rows.Scan(&post.Id, &post.UserId, &post.Text, &post.LocationCode, &post.CreatedAt, &post.LikesNumber)
		if err != nil {
			if gateway.db.Tx != nil {
				gateway.db.Tx.Rollback()
			}
			logger.Warning(err)
			return posts, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (gateway PostsGateway) ListByFilter(filter ListFilter) ([]entities.Post, error) {
	builder := gateway.db.GetBuilder()
	if filter.UserIdEquals != 0 {
		builder = builder.And().Equals("user_id", filter.UserIdEquals)
	}
	if filter.LocationCodeStartsWith != "" {
		builder = builder.And().StartsWith("location_code", filter.LocationCodeStartsWith)
	}
	if filter.CreatedAfter != "" {
		builder = builder.And().GreaterOrEquals("created_at", filter.CreatedAfter)
	}
	if filter.OrderByCreatedAtDesc {
		builder = builder.OrderBy("created_at", "DESC")
	}
	if filter.Offset != 0 {
		builder = builder.Offset(filter.Offset)
	}
	if filter.Limit != 0 {
		builder = builder.Limit(filter.Limit)
	}
	return gateway.listByBuilder(builder)
}
