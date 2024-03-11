package models

import (
	"github.com/gin-gonic/gin"
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

func (u *User) Create(c *gin.Context) error {
	if _, err := db.NewInsert().Model(u).Exec(c); err != nil {
		return err
	}
	return nil
}

func (u *Users) GetAll(c *gin.Context) (int, error) {
	q := db.NewSelect().Model(u)
	if c.Value("search_query") != "" {
		q.Where("username ILIKE ?", c.Value("search_query"))
	}
	return ApplyPagination(q, c).ScanAndCount(c.Request.Context())
}

func (u *User) GetById(c *gin.Context) error {
	return db.NewSelect().Model(u).Scan(c)
}

func (u *User) Update(c *gin.Context) error {
	if _, err := db.NewUpdate().Model(u).WherePK().Exec(c); err != nil {
		return err
	}
	return nil
}

func (u *User) Delete(c *gin.Context) error {
	if _, err := db.NewDelete().Model(u).WherePK().Exec(c); err != nil {
		return err
	}
	return nil
}

func (u *User) GetByUsername(c *gin.Context) error {
	return db.NewSelect().
		Model(u).
		Relation("Auth").
		Where("username = ?", strings.ToLower(u.Username)).
		Scan(c)
}
