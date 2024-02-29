package models

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	Id uint `json:"id"` // field named `ID` will be used as a primary field by default
}

type BaseModelUUID struct {
	Id string `json:"id" gorm:"primaryKey"`
}

func (m *BaseModelUUID) BeforeCreate(tx *gorm.DB) (err error) {
	table := tx.Statement.Table
	m.Id = fmt.Sprintf("%s_%s", table[:4], uuid.New().String())
	return nil
}
