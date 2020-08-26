/*
 * @Version: 0.0.1
 * @Author: ider
 * @Date: 2020-07-27 22:37:43
 * @LastEditors: ider
 * @LastEditTime: 2020-08-26 21:36:31
 * @Description:
 */
package ssq

import (
	"strconv"
	"strings"
	"time"

	"lottery/crawler/ssq"
)

var DBAll DoubleBollAll

type DoubleBollAll struct {
	lastUpdate  time.Time
	AwardCheckS []*AwardCheck
}

func NewDoubleBollAll() {
	var acks []*AwardCheck
	ret := ssq.SsqSchedule()
	for _, obj := range *ret {
		// database.LUCKDB.Save(&obj)
		var redbolls []int
		for _, s := range strings.Split(obj.RedBall, " ") {
			ss, _ := strconv.Atoi(s)
			redbolls = append(redbolls, ss)
		}
		buleboll, _ := strconv.Atoi(obj.BlueBall)
		ac := NewAwardCheck(obj.Qihao, redbolls, buleboll)
		acks = append(acks, ac)
	}

	DBAll = DoubleBollAll{
		lastUpdate:  time.Now(),
		AwardCheckS: acks,
	}
}

// 每隔一小时更新一次数据
// func (c *DoubleBollAll) Update() {
// 	k := time.Now()
// 	d, _ := time.ParseDuration("+1h")
// 	k.Add(d)
// 	if c.lastUpdate.After(k) {
// 		log.Println("update DoubleBollAll")
// 		NewDoubleBollAll()
// 	}
// }

// 历史中奖判断,只返回5等奖以上的
func (c *DoubleBollAll) AwardCheck(ssqball *SsqBall) *[]AwardResult {
	var retAwardResult []AwardResult
	for _, ac := range c.AwardCheckS {
		ret := *ac.CheckMulitAward(ssqball)
		// if ret.A > 0 || ret.B > 0 || ret.C > 0 || ret.D > 0 || ret.E > 0 {
		if ret.A > 0 || ret.B > 0 || ret.C > 0 {
			retAwardResult = append(retAwardResult, ret)
		}
	}
	return &retAwardResult

}

// 判断指定期号是否中奖
func (c *DoubleBollAll) AwardCheckQiHao(ssqball *SsqBall, qihao string) (bool, *AwardResult) {
	for _, ac := range c.AwardCheckS {
		if ac.Qihao == qihao {
			ret := *ac.CheckMulitAward(ssqball)
			return true, &ret
		}
	}
	return false, nil

}
