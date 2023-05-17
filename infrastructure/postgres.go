package infrastructure

import (
	"capstone/config"
	"capstone/internal/pgx/logger"
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

type DBTX interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Rebind(query string) string
}

func NewPostgreSqlDB(dbConfig config.DatabaseConfig) *sqlx.DB {
	conf, err := pgx.ParseConfig("")
	if err != nil {
		panic(err)
	}

	// add pgx custom logger
	if config.IsDebug() {
		pgxLoggerConfig := logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			Colorful:      true,
			LogLevel:      pgx.LogLevelInfo,
		}
		pgxLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), pgxLoggerConfig)
		conf.Logger = pgxLogger
		conf.LogLevel = pgxLoggerConfig.LogLevel
	}

	pgConnConf := &conf.Config
	pgConnConf.Host = dbConfig.Host
	pgConnConf.Port = dbConfig.Port
	pgConnConf.Database = dbConfig.DatabaseName
	pgConnConf.User = dbConfig.Username
	pgConnConf.Password = dbConfig.Password
	pgConnConf.RuntimeParams["timezone"] = "UTC"

	pgxDB := stdlib.OpenDB(*conf)
	if err = pgxDB.Ping(); err != nil {
		pgxDB.Close()
		// panic(err)
	}

	db := sqlx.NewDb(pgxDB, "pgx")
	return db
}
