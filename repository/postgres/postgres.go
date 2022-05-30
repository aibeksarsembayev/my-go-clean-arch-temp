package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/quazar2000/my-go-clean-arch-temp/config"
)

// type Database struct {
// 	Conn *sql.DB
// }

func NewConnectPostgresDB(conf *config.Config, logger pgx.Logger, logLevel pgx.LogLevel) (*pgxpool.Pool, error) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", conf.Database.DBUser, conf.Database.DBPass, conf.Database.DBHost, conf.Database.DBPort, conf.Database.DBName)

	dbConf, err := pgxpool.ParseConfig(dbUrl)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", err)
	}

	dbConf.ConnConfig.Logger = logger
	// Set the log level for pgx, if set.
	if logLevel != 0 {
		dbConf.ConnConfig.LogLevel = logLevel
	}

	// Need to move config json !!!!!!!!!!!!!!!!!
	dbConf.MaxConns = 25
	dbConf.MaxConnLifetime = 5 * time.Minute

	pool, err := pgxpool.ConnectConfig(context.Background(), dbConf)
	if err != nil {
		return nil, fmt.Errorf("postgres: %w", err)
	}

	return pool, nil
}
