/*
 * @Version: 0.0.1
 * @Author: ider
 * @Date: 2020-07-23 21:09:54
 * @LastEditors: ider
 * @LastEditTime: 2020-07-23 21:49:19
 * @Description:
 */

package models

import "github.com/jinzhu/gorm"

type DoubleBall struct {
	gorm.Model
	Qihao        string `gorm:"UNIQUE_INDEX;not null;" json:"qihao"`
	OpenDate     string `gorm:"UNIQUE_INDEX;not null;" json:"opendate"`
	RedBall      string `json:"redball"`
	RedBallOrder string `json:"redballorder"`
	BlueBall     string `json:"blueball"`
	TotalSales   int64  `json:"totalsales"`
	OneCount     int64  `json:"onecount"`
	OneAward     int64  `json:"oneaward"`
	TwoCount     int64  `json:"twocount"`
	TwoAward     int64  `json:"twoaward"`
	ThreeCount   int64  `json:"threecount"`
	ThreeAward   int64  `json:"threeaward"`
	Jackpot      int64  `json:"jackpot"`
}
