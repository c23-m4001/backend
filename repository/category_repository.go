package repository

import (
	"capstone/infrastructure"
	"capstone/model"
	"context"

	"github.com/Masterminds/squirrel"
)

type CategoryRepository interface {
	// create
	Insert(ctx context.Context, category *model.Category) error
	InsertMany(ctx context.Context, categories []model.Category) error

	// read
	Count(ctx context.Context, options ...model.CategoryQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.CategoryQueryOption) ([]model.Category, error)
	Get(ctx context.Context, id string) (*model.Category, error)
	IsExist(ctx context.Context, id string) (bool, error)

	// update
	Update(ctx context.Context, category *model.Category) error

	// delete
	Delete(ctx context.Context, category *model.Category) error
	Truncate(ctx context.Context) error
}

type categoryRepository struct {
	db infrastructure.DBTX
}

func NewCategoryRepository(
	db infrastructure.DBTX,
) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) fetch(stmt squirrel.Sqlizer) ([]model.Category, error) {
	categories := []model.Category{}
	if err := fetch(r.db, &categories, stmt); err != nil {
		return nil, err
	}

	return categories, nil
}

func (r *categoryRepository) get(stmt squirrel.Sqlizer) (*model.Category, error) {
	category := model.Category{}
	if err := get(r.db, &category, stmt); err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepository) prepareQuery(option model.CategoryQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(model.CategoryTableName)

	orStatements := squirrel.Or{}
	andStatements := squirrel.And{}

	if option.UserId != nil {
		andStatements = append(andStatements, squirrel.Eq{"user_id": option.UserId})
	}

	if option.Phrase != nil {
		phrase := "%" + *option.Phrase + "%"
		andStatements = append(andStatements, squirrel.ILike{"name": phrase})
	}

	if option.IncludeGlobal != nil {
		orStatements = squirrel.Or{
			squirrel.Eq{"is_global": option.IncludeGlobal},
		}
	}

	orStatements = append(orStatements, andStatements)

	stmt = stmt.Where(orStatements)

	stmt = option.Prepare(stmt)

	return stmt
}

func (r *categoryRepository) Insert(ctx context.Context, category *model.Category) error {
	return defaultInsert(r.db, ctx, category, "*")
}

func (r *categoryRepository) InsertMany(ctx context.Context, categories []model.Category) error {
	arr := []model.BaseModel{}
	for i := range categories {
		arr = append(arr, &categories[i])
	}
	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *categoryRepository) Count(ctx context.Context, options ...model.CategoryQueryOption) (int, error) {
	option := model.CategoryQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.SetDefault()
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, stmt)
}

func (r *categoryRepository) Fetch(ctx context.Context, options ...model.CategoryQueryOption) ([]model.Category, error) {
	option := model.CategoryQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.SetDefault()

	stmt := r.prepareQuery(option)

	return r.fetch(stmt)
}

func (r *categoryRepository) Get(ctx context.Context, id string) (*model.Category, error) {
	stmt := stmtBuilder.Select("*").
		From(model.CategoryTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(stmt)
}

func (r *categoryRepository) IsExist(ctx context.Context, id string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("*").
			From(model.CategoryTableName).
			Where(squirrel.Eq{"id": id}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, stmt)
}

func (r *categoryRepository) Update(ctx context.Context, category *model.Category) error {
	excludedColumns := []string{
		"user_id",
	}
	columns := extractColumnsFromBaseModel(category, excludedColumns)
	return defaultUpdate(r.db, ctx, category, columns, nil)
}

func (r *categoryRepository) Delete(ctx context.Context, category *model.Category) error {
	return defaultDestroy(r.db, ctx, category, nil)
}

func (r *categoryRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, model.CategoryTableName)
}
