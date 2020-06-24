package db

import (
	"context"
	"fmt"
	"time"

	log "github.com/oatsaysai/simple-go-http-server/log"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type DB struct {
	logger log.Logger
	Config *Config
	DB     *pgxpool.Pool
}

func New(config *Config, logger log.Logger) (db *DB, err error) {
	db = &DB{
		logger: logger.WithFields(log.Fields{
			"module": "db",
		}),
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost,
		config.DBPort,
		config.DBUsername,
		config.DBPassword,
		config.DBName,
	)
	db.logger.Infof("DB conection string: %s", connStr)

	//db, err = pgx.Connect(context.Background(), connStr)
	var connectConf, _ = pgxpool.ParseConfig(connStr)
	connectConf.MaxConns = config.MaxOpenConns
	connectConf.MaxConnLifetime = time.Duration(config.MaxConnLifetime) * time.Second // use defaults until we have benchmarked this further

	//connectConf.HealthCheckPeriod = 300 * time.Second
	//connectConf.ConnConfig.PreferSimpleProtocol = true // don't wrap queries into transactions
	connectConf.ConnConfig.Logger = NewDatabaseLogger(&db.logger)
	connectConf.ConnConfig.LogLevel = pgx.LogLevelWarn
	db.DB, err = pgxpool.ConnectConfig(context.Background(), connectConf)
	if err != nil {
		db.logger.Errorf("Error connecting to postgres: %+v")
		return nil, err
	}

	return db, nil
}

func (db *DB) Close() error {
	db.DB.Close()
	return nil
}
