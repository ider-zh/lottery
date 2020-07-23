package database

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type CacheSummary struct {
	gorm.Model
	Title  string `gorm:"UNIQUE_INDEX;not null"`
	Data   postgres.Jsonb
	IsExit bool
}
