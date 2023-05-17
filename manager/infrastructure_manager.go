package manager

import (
	"capstone/config"
	"capstone/database/migration"
	"capstone/infrastructure"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	migratePgx "github.com/golang-migrate/migrate/v4/database/pgx"
	"github.com/jmoiron/sqlx"
)

type InfrastructureManager interface {
	Close() error

	GetDB() *sqlx.DB
	MigrateDB(isRollback bool, steps int) error
	RefreshDB() error

	GetLoggerStack() infrastructure.LoggerStack
}

type infrastructureManager struct {
	loggerStack infrastructure.LoggerStack
	sqlDB       *sqlx.DB
}

func newInfrastructureManager() InfrastructureManager {
	return &infrastructureManager{
		sqlDB:       infrastructure.NewPostgreSqlDB(config.GetPostgresConfig()),
		loggerStack: infrastructure.NewLoggerStack(config.GetLogChannels()),
	}
}

func (i infrastructureManager) Close() error {
	if sqlDB := i.GetDB(); i.sqlDB != nil {
		return sqlDB.Close()
	}
	return nil
}

func (m infrastructureManager) GetDB() *sqlx.DB {
	return m.sqlDB
}

func (m infrastructureManager) MigrateDB(isRollback bool, steps int) error {
	dbDriver, err := migratePgx.WithInstance(m.GetDB().DB, &migratePgx.Config{})
	if err != nil {
		return err
	}

	migrator, err := migrate.NewWithInstance("", migration.SourceDriver(), "pgx", dbDriver)
	if err != nil {
		return err
	}

	if isRollback {
		_, _, err := migrator.Version()
		if err != nil {
			return err
		}

		if steps > 0 {
			err = migrator.Steps(-1 * int(steps))
		} else {
			err = migrator.Down()
		}

		if err != nil {
			return err
		}
	} else {
		var err error
		if steps > 0 {
			err = migrator.Steps(int(steps))
		} else {
			err = migrator.Up()
		}

		if err != nil && err != migrate.ErrNoChange {
			return err
		}
	}

	return nil
}

func (m infrastructureManager) RefreshDB() error {
	if err := m.dropDB(); err != nil {
		return err
	}

	if err := m.createDB(); err != nil {
		return err
	}

	if err := m.MigrateDB(false, 0); err != nil {
		return err
	}

	return nil
}

func (m *infrastructureManager) createDB() error {
	dbConfig := config.GetPostgresConfig()
	dbConfig.DatabaseName = ""
	sqlDB := infrastructure.NewPostgreSqlDB(dbConfig)

	if _, err := sqlDB.Exec(fmt.Sprintf("CREATE DATABASE \"%s\" WITH ENCODING='UTF8';", config.GetPostgresConfig().DatabaseName)); err != nil {
		return err
	}

	m.sqlDB = infrastructure.NewPostgreSqlDB(config.GetPostgresConfig())

	return sqlDB.Close()
}

func (m infrastructureManager) dropDB() error {
	dbConfig := config.GetPostgresConfig()
	dbConfig.DatabaseName = ""
	sqlDB := infrastructure.NewPostgreSqlDB(dbConfig)

	if _, err := sqlDB.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS \"%s\" WITH (FORCE);", config.GetPostgresConfig().DatabaseName)); err != nil {
		return err
	}

	return sqlDB.Close()
}

func (m infrastructureManager) GetLoggerStack() infrastructure.LoggerStack {
	return m.loggerStack
}
