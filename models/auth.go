package models

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/uptrace/bun"
)

type RefreshToken struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}
type AccessToken struct {
	AccessToken string `json:"access_token"`
}
type AuthClaims struct {
	UserId string `json:"user_id"`
	*jwt.RegisteredClaims
}

type Auth struct {
	bun.BaseModel `bun:"table:auth"`
	BaseModelUUID
	AccessToken
	RefreshToken
	BaseModelAudit
	BaseModelSoftDelete

	UserId string `json:"user_id"`
}

func (a *Auth) Upsert(c context.Context) error {
	if _, err := db.NewInsert().Model(a).On("CONFLICT (user_id) DO UPDATE").Exec(c); err != nil {
		return err
	}
	return nil
}
