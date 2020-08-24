/*
 * @Version: 0.0.1
 * @Author: ider
 * @Date: 2020-07-23 21:09:54
 * @LastEditors: ider
 * @LastEditTime: 2020-08-24 11:59:15
 * @Description:
 */

package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// 双色球采集后保存的结构
type DoubleBall struct {
	Qihao        string    `gorm:"primary_key;" json:"qihao"`
	OpenDate     time.Time `json:"opendate"`
	RedBall      string    `json:"redball"`
	RedBallOrder string    `json:"redballorder"`
	BlueBall     string    `json:"blueball"`
	TotalSales   int64     `json:"totalsales"`
	OneCount     int64     `json:"onecount"`
	OneAward     int64     `json:"oneaward"`
	TwoCount     int64     `json:"twocount"`
	TwoAward     int64     `json:"twoaward"`
	ThreeCount   int64     `json:"threecount"`
	ThreeAward   int64     `json:"threeaward"`
	Jackpot      int64     `json:"jackpot"`
}

// 用户投注数
type UserDoubleBall struct {
	gorm.Model
	Qihao    string      ` json:"qihao" form:"qihao"`
	RedBall  string      `json:"redball" form:"redball"`
	BlueBall string      `json:"blueball" form:"blueball"`
	Multiple int         `json:"multiple" form:"multiple"` //倍数
	IsOpen   bool        `gorm:"not null;default:'false'"` //是否计算开奖
	A        int         `json:"a"`
	B        int         `json:"b"`
	C        int         `json:"c"`
	D        int         `json:"d"`
	E        int         `json:"e"`
	F        int         `json:"f"`
	History  interface{} `gorm:"-" json:"history"` //历史中奖，实时计算，用来表达历史中奖
}

func (c *UserDoubleBall) ToString() []string {
	var retS []string
	if c.A > 0 {
		retS = append(retS, fmt.Sprintf("一等奖 %d 注", c.A))
	}
	if c.B > 0 {
		retS = append(retS, fmt.Sprintf("二等奖 %d 注", c.B))
	}
	if c.C > 0 {
		retS = append(retS, fmt.Sprintf("三等奖 %d 注", c.C))
	}
	if c.D > 0 {
		retS = append(retS, fmt.Sprintf("四等奖 %d 注", c.D))
	}
	if c.E > 0 {
		retS = append(retS, fmt.Sprintf("五等奖 %d 注", c.E))
	}
	if c.F > 0 {
		retS = append(retS, fmt.Sprintf("六等奖 %d 注", c.F))
	}
	return retS
}
