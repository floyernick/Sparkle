package database

import (
	"fmt"
)

func (db DB) getBuilder() queryBuilder {
	return queryBuilder{}
}

type queryBuilder struct {
	clause     string
	params     []interface{}
	offset     *int
	limit      *int
	groupBy    *string
	orderItems [][2]string
}

func (builder queryBuilder) Equals(field string, value interface{}) queryBuilder {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s = $%d ", field, len(builder.params))
	return builder
}

func (builder queryBuilder) NotEquals(field string, value interface{}) queryBuilder {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s != $%d ", field, len(builder.params))
	return builder
}

func (builder queryBuilder) Greater(field string, value interface{}) queryBuilder {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s > $%d ", field, len(builder.params))
	return builder
}

func (builder queryBuilder) GreaterOrEquals(field string, value interface{}) queryBuilder {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s >= $%d ", field, len(builder.params))
	return builder
}

func (builder queryBuilder) Less(field string, value interface{}) queryBuilder {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s < $%d ", field, len(builder.params))
	return builder
}

func (builder queryBuilder) LessOrEquals(field string, value interface{}) queryBuilder {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s <= $%d ", field, len(builder.params))
	return builder
}

func (builder queryBuilder) Contains(field string, value interface{}) queryBuilder {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("$%d = ANY(%s) ", len(builder.params), field)
	return builder
}

func (builder queryBuilder) Like(field string, value interface{}) queryBuilder {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s ILIKE '%%' || $%d || '%%' ", field, len(builder.params))
	return builder
}

func (builder queryBuilder) StartsWith(field string, value interface{}) queryBuilder {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s ILIKE $%d || '%%' ", field, len(builder.params))
	return builder
}

func (builder queryBuilder) In(field string, value interface{}) queryBuilder {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s = ANY($%d) ", field, len(builder.params))
	return builder
}

func (builder queryBuilder) NotIn(field string, value interface{}) queryBuilder {
	builder.params = append(builder.params, value)
	builder.clause += fmt.Sprintf("%s = ANY($%d) ", field, len(builder.params))
	return builder
}

func (builder queryBuilder) And() queryBuilder {
	builder.clause += "AND "
	return builder
}

func (builder queryBuilder) Or() queryBuilder {
	builder.clause += "OR "
	return builder
}

func (builder queryBuilder) Group() queryBuilder {
	builder.clause += "("
	return builder
}

func (builder queryBuilder) EndGroup() queryBuilder {
	builder.clause += ") "
	return builder
}

func (builder queryBuilder) GroupBy(field string) queryBuilder {
	builder.groupBy = &field
	return builder
}

func (builder queryBuilder) Offset(offset int) queryBuilder {
	builder.offset = &offset
	return builder
}

func (builder queryBuilder) Limit(limit int) queryBuilder {
	builder.limit = &limit
	return builder
}

func (builder queryBuilder) OrderBy(field string, direction string) queryBuilder {
	builder.orderItems = append(builder.orderItems, [2]string{field, direction})
	return builder
}

func (builder queryBuilder) formatQuery(query string) string {
	if len(builder.params) != 0 {
		query += fmt.Sprintf(" WHERE 1=1 %s", builder.clause)
	}
	if builder.groupBy != nil {
		query += fmt.Sprintf(" GROUP BY %s", *builder.groupBy)
	}
	if len(builder.orderItems) != 0 {
		query += " ORDER BY "
		for i, orderItem := range builder.orderItems {
			if i != 0 {
				query += ", "
			}
			query += fmt.Sprintf("%s %s", orderItem[0], orderItem[1])
		}
	}
	if builder.offset != nil {
		query += fmt.Sprintf(" OFFSET %d", *builder.offset)
	}
	if builder.limit != nil {
		query += fmt.Sprintf(" LIMIT %d", *builder.limit)
	}
	return query
}
