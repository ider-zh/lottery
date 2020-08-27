package database

import (
	"log"
	"lottery/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	LUCKDB *gorm.DB
)

func NewLuckDBConn(PostgresConn string) {
	var err error
	LUCKDB, err = gorm.Open(postgres.Open(PostgresConn), &gorm.Config{})
	if err != nil {
		log.Fatal("ping 失败", err)
	}
	// dont do this
	// WikiDB.DropTableIfExists(&Revision{}, &Article{})
	// create table
	LUCKDB.AutoMigrate(&models.DoubleBall{}, &models.UserDoubleBall{})
}
