package models

import (
	"context"
	"github.com/uptrace/bun"
)

var db *bun.DB

func SetDatabase(ndb *bun.DB) error {
	db = ndb
	c := context.Background()
	return ping(&c)
}

func ping(c *context.Context) error {
	_, err := db.NewRaw("SELECT 1").Exec(*c)
	return err
}
