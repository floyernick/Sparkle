package locations

import (
	"Sparkle/entities"
	"Sparkle/tools/logger"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type ListFilter struct {
	ChildCodeLengthEquals int
	ParentCodeStartsWith  string
	CreatedAfter          string
	Limit                 int
}

func (gateway LocationsGateway) ListByFilter(filter ListFilter) ([]entities.Location, error) {

	var locations []entities.Location

	cachedValue, err := gateway.cache.Client.Get(gateway.cache.Ctx, filter.ParentCodeStartsWith).Result()

	if err == nil {
		err = json.Unmarshal([]byte(cachedValue), &locations)
		if err == nil {
			return locations, nil
		} else {
			logger.Error(err)
		}
	} else if err != redis.Nil {
		logger.Error(err)
	}

	query := fmt.Sprintf("SELECT SUBSTRING(location_code FROM 1 FOR %d) AS code, COUNT(id) AS number FROM posts", filter.ChildCodeLengthEquals)

	builder := gateway.db.GetBuilder()
	builder = builder.And().StartsWith("location_code", filter.ParentCodeStartsWith).
		And().GreaterOrEquals("created_at", filter.CreatedAfter).GroupBy("code").OrderBy("number", "DESC").Limit(filter.Limit)

	query = builder.FormatQuery(query)

	rows, err := gateway.db.Performer().Query(query, builder.Params...)

	if err != nil && err != sql.ErrNoRows {
		if gateway.db.Tx != nil {
			gateway.db.Tx.Rollback()
		}
		logger.Error(err)
		return locations, err
	}

	defer rows.Close()

	for rows.Next() {
		var location entities.Location
		err := rows.Scan(&location.Code, &location.PostsNumber)
		if err != nil {
			if gateway.db.Tx != nil {
				gateway.db.Tx.Rollback()
			}
			logger.Error(err)
			return locations, err
		}
		locations = append(locations, location)
	}

	cachingValue, err := json.Marshal(locations)

	if err != nil {
		logger.Error(err)
		return locations, nil
	}

	err = gateway.cache.Client.SetNX(gateway.cache.Ctx, filter.ParentCodeStartsWith, cachingValue, 10*time.Minute).Err()

	if err != nil {
		logger.Error(err)
	}

	return locations, nil
}
