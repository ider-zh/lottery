package test

import (
	"fmt"
	"lottery/database"
	"lottery/models"
	"testing"
)

func Test_Lab(t *testing.T) {
	var adBall []models.UserDoubleBall
	database.LUCKDB.Where("is_open = ?", true).Order("qihao desc").Limit(5000).Find(&adBall)
	na := 0
	for _, obj := range adBall {
		if obj.A == 0 && obj.B == 0 && obj.C == 0 && obj.D == 0 && obj.E == 0 && obj.F == 0 {
			na++
		} else {
			break
		}
	}
	subject := fmt.Sprintf("，连续 %d 注未中奖！！！", na)
	fmt.Println(subject)
}
