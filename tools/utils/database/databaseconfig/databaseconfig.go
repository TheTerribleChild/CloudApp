package databaseconfig

import(
	"time"
)

type DatabaseConfig struct {
	Host            string
	Port            int
	User            string
	Password        string
	Database        string
	Schema          string
	MaxConns        int
	MaxIdleConns    int
	MaxConnLifetime time.Duration
}