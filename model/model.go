package model

import (
	"capstone/constant"
	"capstone/data_type"
	"capstone/util"
	"fmt"
	"strings"

	"github.com/Masterminds/squirrel"
)

// for base repository model
type BaseModel interface {
	TableName() string
	TableIds() []string
	ToMap() map[string]interface{}
	GetCreatedAt() data_type.NullDateTime
	GetUpdatedAt() data_type.NullDateTime
	SetCreatedAt(dateTime data_type.NullDateTime)
	SetUpdatedAt(dateTime data_type.NullDateTime)
}

// created_at updated_at implementions for every table
type Timestamp struct {
	CreatedAt data_type.NullDateTime `db:"created_at"`
	UpdatedAt data_type.NullDateTime `db:"updated_at"`
}

func (m Timestamp) GetCreatedAt() data_type.NullDateTime {
	return m.CreatedAt
}

func (m Timestamp) GetUpdatedAt() data_type.NullDateTime {
	return m.UpdatedAt
}

func (m *Timestamp) SetCreatedAt(dateTime data_type.NullDateTime) {
	m.CreatedAt = dateTime
}

func (m *Timestamp) SetUpdatedAt(dateTime data_type.NullDateTime) {
	m.UpdatedAt = dateTime
}

// for pagination
type QueryOption struct {
	SelectOption
	PaginationOption
	IsCount bool
}

func (o QueryOption) Prepare(stmt squirrel.SelectBuilder) squirrel.SelectBuilder {
	if o.IsCount {
		o.Fields = []string{"COUNT(*) as count"}

		stmt = stmt.Columns(o.Fields...)

		return stmt
	}

	o.SelectOption.Prepare(stmt)
	o.PaginationOption.Prepare(stmt)

	return stmt
}

func (o *QueryOption) SetDefault() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"*"}
	}

	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{}
	}

	if o.Limit == nil {
		o.Limit = util.IntP(constant.PaginationDefaultLimit)
	}

	if o.Page == nil {
		o.Page = util.IntP(constant.PaginationDefaultPage)
	}
}

type SelectOption struct {
	Fields []string
	Sorts
}

func (o SelectOption) Prepare(stmt squirrel.SelectBuilder) squirrel.SelectBuilder {
	stmt = stmt.Columns(o.Fields...)

	o.Sorts.Prepare(stmt)

	return stmt
}

type Sorts []struct {
	Field     string
	Direction string
}

func (o Sorts) Prepare(stmt squirrel.SelectBuilder) squirrel.SelectBuilder {
	if len(o) > 0 {
		for _, sort := range o {
			direction := strings.ToUpper(sort.Direction)
			if direction == "ASC" || direction == "DESC" {
				stmt = stmt.OrderBy(fmt.Sprintf("%s %s", sort.Field, direction))
			}
		}
	}

	return stmt
}

type PaginationOption struct {
	Limit *int
	Page  *int
}

func (o PaginationOption) Prepare(stmt squirrel.SelectBuilder) squirrel.SelectBuilder {
	if o.Page != nil && o.Limit != nil && *o.Page > 1 && *o.Limit > 0 {
		offset := (*o.Page - 1) * *o.Limit
		stmt = stmt.Limit(uint64(*o.Limit))
		stmt = stmt.Offset(uint64(offset))
	}
	return stmt
}
