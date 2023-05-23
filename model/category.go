package model

const CategoryTableName = "categories"

var category BaseModel = &Category{}

type Category struct {
	Id        string  `db:"id"`
	UserId    *string `db:"user_id"`
	Name      string  `db:"name"`
	IsGlobal  bool    `db:"is_global"`
	IsExpense bool    `db:"is_expense"`
	Timestamp
}

func (m Category) TableName() string {
	return CategoryTableName
}

func (m Category) TableIds() []string {
	return []string{"id"}
}

func (m Category) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         m.Id,
		"user_id":    m.UserId,
		"name":       m.Name,
		"is_global":  m.IsGlobal,
		"is_expense": m.IsExpense,
		"created_at": m.CreatedAt,
		"updated_at": m.UpdatedAt,
	}
}

type CategoryQueryOption struct {
	QueryOption
	IncludeGlobal *bool
	IsExpense     *bool
	IsGlobal      *bool
	UserId        *string
	Phrase        *string
}

func (o *CategoryQueryOption) SetDefault() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"*"}
	}

	o.QueryOption.SetDefault()

	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "name", Direction: "asc"},
		}
	}

	o.translateSorts(category, o.translateSort)
}
