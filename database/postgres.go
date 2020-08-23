package database

import (
	"log"
	"lottery/internal/award/ssq"
	"lottery/models"
	"lottery/utils"

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

// 双色当期中奖计算
func UpdateSsqAward() {
	var ret_un_open []*models.UserDoubleBall
	LUCKDB.Model(&models.UserDoubleBall{}).Where("is_open = ?", false).Order("qihao desc").Find(&ret_un_open)
	// 历史兑奖
	for _, obj := range ret_un_open {
		ssqball := ssq.SsqBall{Redboll: utils.BollStrToNum(obj.RedBall), Blueboll: utils.BollStrToNum(obj.BlueBall)}
		if status, ret := ssq.DBAll.AwardCheckQiHao(&ssqball, obj.Qihao); status == true {
			obj.A = ret.A
			obj.B = ret.B
			obj.C = ret.C
			obj.D = ret.D
			obj.E = ret.E
			obj.F = ret.F
			obj.IsOpen = true
			LUCKDB.Save(obj)
		}
	}
}
