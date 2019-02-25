package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"theterriblechild/CloudApp/applications/adminapp/internal/dal"
)

type PostgreDB struct {
	Host string
	Port int
	User string
	Password string
	Database string
	db *sql.DB
}

func (instance *PostgreDB) InitializeDatabase(config dal.DatabaseConfig) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	config.Host, config.Port, config.User, config.Password, config.Database)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}
	if config.MaxConnLifetime > 0 {
		db.SetConnMaxLifetime(config.MaxConnLifetime)
	}
	if config.MaxConns > 0 {
		db.SetMaxOpenConns(config.MaxConns)
	}
	if config.MaxIdleConns > 0 {
		db.SetMaxIdleConns(config.MaxIdleConns)
	}
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	instance.db = db
	return nil
}