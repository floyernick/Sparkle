package posts

import (
	"Sparkle/entities"
	"Sparkle/gateways/common/database"
	"Sparkle/tools/logger"
	"database/sql"
)

func (gateway PostsGateway) getByBuilder(builder database.QueryBuilder) (entities.Post, error) {

	query := "SELECT id, user_id, text, location_code, created_at, (SELECT COUNT(id) FROM likes WHERE post_id = posts.id) AS likes_number FROM posts"
	query = builder.FormatQuery(query)

	var post entities.Post

	err := gateway.db.Performer().QueryRow(query, builder.Params...).Scan(&post.Id, &post.UserId, &post.Text, &post.LocationCode, &post.CreatedAt, &post.LikesNumber)

	if err != nil && err != sql.ErrNoRows {
		if gateway.db.Tx != nil {
			gateway.db.Tx.Rollback()
		}
		logger.Warning(err)
		return post, err
	}

	return post, nil
}

func (gateway PostsGateway) GetById(id int) (entities.Post, error) {
	builder := gateway.db.GetBuilder().And().Equals("id", id)
	return gateway.getByBuilder(builder)
}
