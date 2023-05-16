package repository

import (
	"capstone/constant"
	"capstone/infrastructure"
	"capstone/model"
	"capstone/util"
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

var (
	stmtBuilder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
)

func exec(db infrastructure.DBTX, stmt squirrel.Sqlizer) error {
	query, args, err := stmt.ToSql()
	if err != nil {
		return translateSqlError(err)
	}

	_, err = db.Exec(query, args...)
	return translateSqlError(err)
}

func insert(db infrastructure.DBTX, tableName string, arg map[string]interface{}) error {
	stmt := stmtBuilder.Insert(tableName).SetMap(arg)
	return exec(db, stmt)
}

func defaultInsert(db infrastructure.DBTX, ctx context.Context, m model.BaseModel, columns string) error {
	dateTime := util.CurrentNullDateTime()

	// set timestamp if nil
	if m.GetCreatedAt().DateTimeP() == nil {
		m.SetCreatedAt(dateTime)
	}
	if m.GetUpdatedAt().DateTimeP() == nil {
		m.SetUpdatedAt(dateTime)
	}

	arg := map[string]interface{}{}
	// use all columns if *
	if columns == "*" {
		arg = m.ToMap()
	} else {
		modelArg := m.ToMap()
		for _, column := range strings.Split(columns, ",") {
			column = strings.TrimSpace(column)
			arg[column] = modelArg[column]
		}
	}

	// set timestamp
	arg["created_at"] = m.GetCreatedAt()
	arg["updated_at"] = m.GetUpdatedAt()

	return insert(db, m.TableName(), arg)
}

func insertMany(db infrastructure.DBTX, tableName string, columns []string, arg interface{}) error {
	if v := reflect.ValueOf(arg); !(v.Kind() == reflect.Slice || v.Kind() == reflect.Array) || v.Len() == 0 {
		return nil
	}

	/*
		 	using .NamedExec that need ':' for every column
			example:
			INSERT INTO table id,name, VALUES(:id, :name);
	*/
	values := []interface{}{}
	for _, column := range columns {
		values = append(values, squirrel.Expr(fmt.Sprintf(":%s", column)))
	}

	stmt := stmtBuilder.Insert(tableName).Columns(columns...).Values(values...)
	query, _ := stmt.MustSql()
	_, err := db.NamedExec(query, arg)

	return translateSqlError(err)
}

func defaultInsertMany(db infrastructure.DBTX, ctx context.Context, arr []model.BaseModel, columns string) error {
	if len(arr) == 0 {
		return nil
	}

	// set timestamp if nil
	dateTime := util.CurrentNullDateTime()
	for _, m := range arr {
		if m.GetCreatedAt().DateTimeP() == nil {
			m.SetCreatedAt(dateTime)
		}
		if m.GetUpdatedAt().DateTimeP() == nil {
			m.SetUpdatedAt(dateTime)
		}
		// make created_at and updated_at as unique as possible
		dateTime = dateTime.Add(1 * time.Microsecond)
	}

	columnArr := []string{}
	if columns == "*" {
		for column := range arr[0].ToMap() {
			columnArr = append(columnArr, column)
		}
	} else {
		for _, column := range strings.Split(columns, ",") {
			columnArr = append(columnArr, strings.TrimSpace(column))
		}
	}

	// set timestamp
	if !util.StringInSlice("created_at", columnArr) {
		columnArr = append(columnArr, "created_at")
	}

	if !util.StringInSlice("updated_at", columnArr) {
		columnArr = append(columnArr, "updated_at")
	}

	if err := insertMany(db, arr[0].TableName(), columnArr, arr); err != nil {
		return err
	}

	return nil
}

func fetch(db infrastructure.DBTX, dest interface{}, stmt squirrel.Sqlizer) error {
	query, args, err := stmt.ToSql()
	if err != nil {
		return translateSqlError(err)
	}

	return translateSqlError(db.Select(dest, query, args...))
}

func get(db infrastructure.DBTX, dest interface{}, stmt squirrel.Sqlizer) error {
	query, args, err := stmt.ToSql()
	if err != nil {
		return translateSqlError(err)
	}

	return translateSqlError(db.Get(dest, query, args...))
}

func count(db infrastructure.DBTX, dest interface{}, stmt squirrel.Sqlizer) (int, error) {
	query, args, err := stmt.ToSql()
	if err != nil {
		return 0, translateSqlError(err)
	}

	count := 0
	if err := db.Get(&count, query, args...); err != nil {
		return 0, translateSqlError(err)
	}

	return count, nil
}

func isExist(db infrastructure.DBTX, stmt squirrel.Sqlizer) (bool, error) {
	query, args, err := stmt.ToSql()
	if err != nil {
		return false, translateSqlError(err)
	}

	isExist := false
	if err := db.Get(&isExist, query, args...); err != nil {
		return false, translateSqlError(err)
	}

	return isExist, nil
}

func update(db infrastructure.DBTX, tableName string, arg map[string]interface{}, whereStmt squirrel.Sqlizer) error {
	stmt := stmtBuilder.Update(tableName).SetMap(arg).Where(whereStmt)
	return exec(db, stmt)
}

func defaultUpdate(db infrastructure.DBTX, ctx context.Context, m model.BaseModel, columns string, whereStmt squirrel.Sqlizer) error {
	modelArg := m.ToMap()
	dateTime := util.CurrentNullDateTime()

	m.SetUpdatedAt(dateTime)

	arg := map[string]interface{}{}
	if columns == "*" {
		arg = m.ToMap()
	} else {
		for _, column := range strings.Split(columns, ",") {
			column = strings.TrimSpace(column)
			arg[column] = modelArg[column]
		}
	}

	// set timestamp
	arg["updated_at"] = dateTime

	// remove id and created at from arguments
	delete(arg, "created_at")
	for _, key := range m.TableIds() {
		delete(arg, key)
	}

	// auto use id when where statement is empty
	if whereStmt == nil {
		stmt := map[string]interface{}{}
		for _, key := range m.TableIds() {
			stmt[key] = modelArg[key]
		}
		whereStmt = squirrel.Eq(stmt)
	}

	if err := update(db, m.TableName(), arg, whereStmt); err != nil {
		return err
	}

	return nil
}

func destroy(db infrastructure.DBTX, tableName string, whereStmt squirrel.Sqlizer) error {
	stmt := stmtBuilder.Delete(tableName).Where(whereStmt)

	return exec(db, stmt)
}

func defaultDestroy(db infrastructure.DBTX, ctx context.Context, m model.BaseModel, whereStmt squirrel.Sqlizer) error {
	// auto use id when where statement is empty
	if whereStmt == nil {
		arg := m.ToMap()
		stmt := map[string]interface{}{}
		for _, key := range m.TableIds() {
			stmt[key] = arg[key]
		}
		whereStmt = squirrel.Eq(stmt)
	}

	if err := destroy(db, m.TableName(), whereStmt); err != nil {
		return err
	}

	return nil
}

func truncate(db infrastructure.DBTX, tableName string) error {
	if _, err := db.Exec(
		fmt.Sprintf(`TRUNCATE TABLE %s RESTART IDENTITY CASCADE;`, tableName),
	); err != nil {
		return translateSqlError(err)
	}

	return nil
}

func translateSqlError(err error) error {
	switch v := err.(type) {
	case *pgconn.PgError:
		switch v.Code {
		case pgerrcode.UniqueViolation:
			return constant.ErrDuplicateData

		case pgerrcode.ForeignKeyViolation:
			return constant.ErrForeignKeyViolation

		default:
			return err
		}

	default:
		switch v {
		case sql.ErrNoRows:
			return constant.ErrNoData

		default:
			return err
		}
	}
}

func extractColumnsFromBaseModel(m model.BaseModel, excludedColumns []string) string {
	columnArr := []string{}
	for column := range m.ToMap() {
		if !util.StringInSlice(column, excludedColumns) {
			columnArr = append(columnArr, column)
		}
	}

	return strings.Join(columnArr, ",")
}
