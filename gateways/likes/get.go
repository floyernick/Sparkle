package likes

import (
	"Sparkle/entities"
	"Sparkle/gateways/common/database"
	"Sparkle/tools/logger"
	"database/sql"
)

func (gateway LikesGateway) getByBuilder(builder database.QueryBuilder) (entities.Like, error) {

	query := "SELECT id, user_id, post_id FROM likes"
	query = builder.FormatQuery(query)

	var like entities.Like

	err := gateway.db.Performer().QueryRow(query, builder.Params...).Scan(&like.Id, &like.UserId, &like.PostId)

	if err != nil && err != sql.ErrNoRows {
		if gateway.db.Tx != nil {
			gateway.db.Tx.Rollback()
		}
		logger.Warning(err)
		return like, err
	}

	return like, nil
}

func (gateway LikesGateway) GetByUserIdAndPostId(userId int, postId int) (entities.Like, error) {
	builder := gateway.db.GetBuilder().And().Equals("post_id", postId).And().Equals("user_id", userId)
	return gateway.getByBuilder(builder)
}
