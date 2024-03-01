package models

import (
	"context"
	"github.com/uptrace/bun"
)

var db *bun.DB

func SetDatabase(ndb *bun.DB) error {
	db = ndb

	return PingDatabase()
}

func PingDatabase() error {
	c := context.Background()
	_, err := db.NewRaw("SELECT 1").Exec(c)
	return err
}
