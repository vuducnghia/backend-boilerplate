package models

import (
	"context"
	"strings"
)

type UserPassword struct {
	Password string `json:"password" binding:"required"`
}
type UserCredentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type User struct {
	BaseModelUUID
	Username    string `json:"username" binding:"required"`
	Password    string `json:"-"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
	BaseModelAudit
	BaseModelSoftDelete

	//	relations
	Auth *Auth `json:"auth" bun:"rel:has-one,join:id=user_id" swaggerignore:"true"`
}
type Users []User

func (u *User) Create(c context.Context) error {
	if _, err := db.NewInsert().Model(u).Exec(c); err != nil {
		return err
	}
	return nil
}

func (u *Users) GetAll(c context.Context) error {
	if err := db.NewSelect().Model(u).Scan(c); err != nil {
		return err
	}
	return nil
}

func (u *User) GetById(c context.Context) error {
	return db.NewSelect().Model(u).Scan(c)
}

func (u *User) Update(c context.Context) error {
	if _, err := db.NewUpdate().Model(u).WherePK().Exec(c); err != nil {
		return err
	}
	return nil
}

func (u *User) Delete(c context.Context) error {
	if _, err := db.NewDelete().Model(u).WherePK().Exec(c); err != nil {
		return err
	}
	return nil
}

func (u *User) GetByUsername(c context.Context) error {
	return db.NewSelect().
		Model(u).
		Relation("Auth").
		Where("username = ?", strings.ToLower(u.Username)).
		Scan(c)
}
