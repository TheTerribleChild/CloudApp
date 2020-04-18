package postgresql

import (
	"fmt"
	"theterriblechild/CloudApp/tools/utils/database/databaseconfig"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetPostgreSQLDB(config databaseconfig.DatabaseConfig) (db *sqlx.DB, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.Database)

	db, err = sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		return
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
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}
	if len(config.Schema) > 0 {
		db.Exec(fmt.Sprintf("set search_path=%s", config.Schema))
	}
	return
}
