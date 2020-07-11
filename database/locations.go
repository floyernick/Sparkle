package database

import (
	"Sparkle/entities"
	"Sparkle/tools/logger"
	"database/sql"
	"fmt"
)

type LocationsFilter struct {
	ChildCodeLengthEquals int
	ParentCodeStartsWith  string
	CreatedAfter          string
	Limit                 int
}

func (db DB) GetLocationsByFilter(filter LocationsFilter) ([]entities.Location, error) {

	query := fmt.Sprintf("SELECT SUBSTRING(location_code FROM 1 FOR %d) AS code, COUNT(id) AS number FROM posts", filter.ChildCodeLengthEquals)

	builder := db.getBuilder()
	builder = builder.And().StartsWith("location_code", filter.ParentCodeStartsWith).
		And().GreaterOrEquals("created_at", filter.CreatedAfter).GroupBy("code").Limit(filter.Limit)

	query = builder.formatQuery(query)

	var locations []entities.Location

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

	return locations, nil
}
