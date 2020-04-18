package dal

import (
	"database/sql"
)

type Account struct {
	ID          string         `db:"id"`
	Name        sql.NullString `db:"name"`
	CreatedDate sql.NullInt64  `db:"created_date"`
}

func (instance Account) GetKeyString() string {
	return instance.ID
}

type User struct {
	ID           string         `db:"id"`
	AccountID    string         `db:"account_id"`
	Email        string         `db:"email"`
	PasswordHash string         `db:"password_hash"`
	FirstName    sql.NullString `db:"first_name"`
	LastName     sql.NullString `db:"last_name"`
	CreatedDate  sql.NullInt64  `db:"created_date"`
}

func (instance User) GetKeyString() string {
	return instance.ID
}
