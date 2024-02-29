package models

import "gorm.io/gorm"

var db *gorm.DB

func SetDatabase(ndb *gorm.DB) {
	db = ndb
}
