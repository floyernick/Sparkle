package database

import (
	"Sparkle/entities"
	"Sparkle/tools/logger"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type LocationsFilter struct {
	ChildCodeLengthEquals int
	ParentCodeStartsWith  string
	CreatedAfter          string
	Limit                 int
}

func (db DB) GetLocationsByFilter(filter LocationsFilter) ([]entities.Location, error) {

	var locations []entities.Location

	cachedValue, err := db.cache.Get(db.ctx, filter.ParentCodeStartsWith).Result()
	if err == nil || err == redis.Nil {
		err = json.Unmarshal([]byte(cachedValue), &locations)
		if err == nil {
			return locations, nil
		} else {
			logger.Warning(err)
		}
	} else {
		logger.Warning(err)
	}

	query := fmt.Sprintf("SELECT SUBSTRING(location_code FROM 1 FOR %d) AS code, COUNT(id) AS number FROM posts", filter.ChildCodeLengthEquals)

	builder := db.getBuilder()
	builder = builder.And().StartsWith("location_code", filter.ParentCodeStartsWith).
		And().GreaterOrEquals("created_at", filter.CreatedAfter).GroupBy("code").Limit(filter.Limit)

	query = builder.formatQuery(query)

	rows, err := db.performer().Query(query, builder.params...)

	if err != nil && err != sql.ErrNoRows {
		if db.tx != nil {
			db.tx.Rollback()
		}
		logger.Warning(err)
		return locations, err
	}

	defer rows.Close()

	for rows.Next() {
		var location entities.Location
		err := rows.Scan(&location.Code, &location.PostsNumber)
		if err != nil {
			if db.tx != nil {
				db.tx.Rollback()
			}
			logger.Warning(err)
			return locations, err
		}
		locations = append(locations, location)
	}

	cachingValue, err := json.Marshal(locations)

	if err != nil {
		logger.Warning(err)
		return locations, err
	}

	err = db.cache.SetNX(db.ctx, filter.ParentCodeStartsWith, cachingValue, 10*time.Minute).Err()

	if err != nil {
		logger.Warning(err)
	}

	return locations, nil
}
