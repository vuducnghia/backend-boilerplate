package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModelAudit struct {
	CreatedAt  time.Time `json:"-"` // Automatically managed by GORM for creation time
	ModifiedAt time.Time `json:"-"` // Automatically managed by GORM for update time
}

type BaseModelSoftDelete struct {
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
