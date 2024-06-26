package application

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"time"
)

var DB *bun.DB

func ConnectDatabase(config *DatabaseConfig) {
	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(config.Host+":"+config.Port),
		pgdriver.WithTLSConfig(nil), // disable sslmode
		pgdriver.WithUser(config.Username),
		pgdriver.WithPassword(config.Password),
		pgdriver.WithDatabase(config.Database),
		pgdriver.WithTimeout(5*time.Second),
		pgdriver.WithDialTimeout(5*time.Second),
		pgdriver.WithReadTimeout(5*time.Second),
		pgdriver.WithWriteTimeout(5*time.Second),
	)

	sqldb := sql.OpenDB(pgconn)
	DB = bun.NewDB(sqldb, pgdialect.New())
}

func (c *DatabaseConfig) EnableDebug() {
	DB.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(c.IsDebug)))
}
