/*
 * @Version: 0.0.1
 * @Author: ider
 * @Date: 2020-07-26 00:15:18
 * @LastEditors: ider
 * @LastEditTime: 2020-08-23 09:41:22
 * @Description:
 */

package main

import (
	"flag"
	"log"

	"lottery/config"
	"lottery/crawler/ssq"
	"lottery/database"

	"gorm.io/gorm/clause"
)

var Mode = flag.Int("mode", 0, "Input mode 1,500 彩票网 双色球 all update")

func init() {
	flag.Parse()
}
func main() {
	cfg := config.GetConfig()
	database.NewLuckDBConn(cfg.PgConn)

	switch *Mode {
	case 1:
		log.Println("更新双色球全部")
		ret := ssq.SsqSchedule()
		for _, obj := range *ret {
			database.LUCKDB.Clauses(clause.OnConflict{DoNothing: true}).Create(&obj)
		}
	}
}
