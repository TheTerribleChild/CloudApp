package postgres

import (
	sq "github.com/Masterminds/squirrel"
)

const (
	UserTable    string = "admin.user"
	AccountTable string = "admin.account"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
