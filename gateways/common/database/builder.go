package database

import (
	"fmt"
)

func (db DB) GetBuilder() QueryBuilder {
	return QueryBuilder{}
}

type QueryBuilder struct {
	clause     string
	Params     []interface{}
	offset     *int
	limit      *int
	groupBy    *string
	orderItems [][2]string
}

func (builder QueryBuilder) Equals(field string, value interface{}) QueryBuilder {
	builder.Params = append(builder.Params, value)
	builder.clause += fmt.Sprintf("%s = $%d ", field, len(builder.Params))
	return builder
}

func (builder QueryBuilder) NotEquals(field string, value interface{}) QueryBuilder {
	builder.Params = append(builder.Params, value)
	builder.clause += fmt.Sprintf("%s != $%d ", field, len(builder.Params))
	return builder
}

func (builder QueryBuilder) Greater(field string, value interface{}) QueryBuilder {
	builder.Params = append(builder.Params, value)
	builder.clause += fmt.Sprintf("%s > $%d ", field, len(builder.Params))
	return builder
}

func (builder QueryBuilder) GreaterOrEquals(field string, value interface{}) QueryBuilder {
	builder.Params = append(builder.Params, value)
	builder.clause += fmt.Sprintf("%s >= $%d ", field, len(builder.Params))
	return builder
}

func (builder QueryBuilder) Less(field string, value interface{}) QueryBuilder {
	builder.Params = append(builder.Params, value)
	builder.clause += fmt.Sprintf("%s < $%d ", field, len(builder.Params))
	return builder
}

func (builder QueryBuilder) LessOrEquals(field string, value interface{}) QueryBuilder {
	builder.Params = append(builder.Params, value)
	builder.clause += fmt.Sprintf("%s <= $%d ", field, len(builder.Params))
	return builder
}

func (builder QueryBuilder) Contains(field string, value interface{}) QueryBuilder {
	builder.Params = append(builder.Params, value)
	builder.clause += fmt.Sprintf("$%d = ANY(%s) ", len(builder.Params), field)
	return builder
}

func (builder QueryBuilder) Like(field string, value interface{}) QueryBuilder {
	builder.Params = append(builder.Params, value)
	builder.clause += fmt.Sprintf("%s ILIKE '%%' || $%d || '%%' ", field, len(builder.Params))
	return builder
}

func (builder QueryBuilder) StartsWith(field string, value interface{}) QueryBuilder {
	builder.Params = append(builder.Params, value)
	builder.clause += fmt.Sprintf("%s ILIKE $%d || '%%' ", field, len(builder.Params))
	return builder
}

func (builder QueryBuilder) In(field string, value interface{}) QueryBuilder {
	builder.Params = append(builder.Params, value)
	builder.clause += fmt.Sprintf("%s = ANY($%d) ", field, len(builder.Params))
	return builder
}

func (builder QueryBuilder) NotIn(field string, value interface{}) QueryBuilder {
	builder.Params = append(builder.Params, value)
	builder.clause += fmt.Sprintf("%s = ANY($%d) ", field, len(builder.Params))
	return builder
}

func (builder QueryBuilder) And() QueryBuilder {
	builder.clause += "AND "
	return builder
}

func (builder QueryBuilder) Or() QueryBuilder {
	builder.clause += "OR "
	return builder
}

func (builder QueryBuilder) Group() QueryBuilder {
	builder.clause += "("
	return builder
}

func (builder QueryBuilder) EndGroup() QueryBuilder {
	builder.clause += ") "
	return builder
}

func (builder QueryBuilder) GroupBy(field string) QueryBuilder {
	builder.groupBy = &field
	return builder
}

func (builder QueryBuilder) Offset(offset int) QueryBuilder {
	builder.offset = &offset
	return builder
}

func (builder QueryBuilder) Limit(limit int) QueryBuilder {
	builder.limit = &limit
	return builder
}

func (builder QueryBuilder) OrderBy(field string, direction string) QueryBuilder {
	builder.orderItems = append(builder.orderItems, [2]string{field, direction})
	return builder
}

func (builder QueryBuilder) FormatQuery(query string) string {
	if len(builder.Params) != 0 {
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
