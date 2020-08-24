/*
 * @Version: 0.0.1
 * @Author: ider
 * @Date: 2020-07-27 21:16:07
 * @LastEditors: ider
 * @LastEditTime: 2020-08-24 11:43:27
 * @Description:
 */

package main

import (
	"log"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//

	"lottery/config"
	"lottery/database"
	sw "lottery/internal/api"
	"lottery/internal/award/ssq"
	"lottery/utils/schedule"
)

func main() {

	cfg := config.GetConfig()
	database.NewLuckDBConn(cfg.PgConn)
	log.Printf("Server started")
	ssq.NewDoubleBollAll()
	schedule.Schedule()
	router := sw.NewRouter()

	// router.Use(gin.Recovery())
	log.Fatal(router.Run(":18080"))
}
