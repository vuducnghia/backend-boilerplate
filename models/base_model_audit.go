package models

import (
	"context"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/schema"
	"time"
)

var _ bun.BeforeAppendModelHook = (*BaseModelAudit)(nil)

type BaseModelAudit struct {
	CreatedAt  time.Time `json:"created_at" bun:",nullzero,notnull,default:current_timestamp" swaggerignore:"true"`
	ModifiedAt time.Time `json:"modified_at" bun:",nullzero,notnull,default:current_timestamp" swaggerignore:"true"`
}

func (b *BaseModelAudit) BeforeAppendModel(ctx context.Context, query schema.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		b.CreatedAt = time.Now()
		b.ModifiedAt = b.CreatedAt
	case *bun.UpdateQuery:
		b.ModifiedAt = time.Now()
	}
	return nil
}

type BaseModelSoftDelete struct {
	DeletedAt time.Time `json:"-" bun:",soft_delete,nullzero"`
}
