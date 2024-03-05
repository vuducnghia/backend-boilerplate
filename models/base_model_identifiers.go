package models

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

var _ bun.BeforeInsertHook = (*BaseModelUUID)(nil)

type BaseModel struct {
	Id int32 `json:"id" bun:"pk"`
}

type BaseModelUUID struct {
	Id string `json:"id" bun:"id,pk" swaggerignore:"true"`
}

func (b *BaseModelUUID) BeforeInsert(ctx context.Context, query *bun.InsertQuery) error {
	id := fmt.Sprintf("'%s_%s'", query.GetTableName()[:4], uuid.New().String())
	query.Value("id", id)
	return nil
}
