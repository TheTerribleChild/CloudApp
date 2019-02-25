package dal

import (
	"time"
)

type AdminDB interface {
	InitializeDatabase(DatabaseConfig) error
	Close()
	CreateAccount()
	CreateUser()
	CreateAgent()
}

type DatabaseConfig struct {
	Host string
	Port int
	User string
	Password string
	Database string
	MaxConns int 
	MaxIdleConns int
	MaxConnLifetime time.Duration
}