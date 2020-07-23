package database

import (
	"log"

	"github.com/jinzhu/gorm"
)

var (
	LUCKDB *gorm.DB
)

func NewLuckDBConn(PostgresConn string) {
	var err error
	LUCKDB, err = gorm.Open("postgres", PostgresConn)
	if err != nil {
		log.Fatal("ping 失败", err)
	}
	// dont do this
	// WikiDB.DropTableIfExists(&Revision{}, &Article{})
	// create table
	LUCKDB.AutoMigrate(&PageViewGather{})

}
