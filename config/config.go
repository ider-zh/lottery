/*
 * @Version: 0.0.1
 * @Author: ider
 * @Date: 2020-07-26 00:38:40
 * @LastEditors: ider
 * @LastEditTime: 2020-08-24 12:41:43
 * @Description:
 */
package config

import (
	"log"

	"github.com/caarlos0/env"
)

type Config struct {
	PgConn  string `env:"PGCONN" envDefault:"host=127.0.0.1 port=5432 user=postgres dbname=ssq password=ider001 sslmode=disable"`
	GinPort string `env:"GINPORT" envDefault:"18080"`
}

func GetConfig() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Println("cfg error ", err)
	}
	return cfg
}
