package api

import (
	"lottery/database"
	"lottery/models"
	"lottery/utils"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"

	"lottery/internal/award/ssq"
)

// // AwardCheckPost - 检测 award
// func AwardCheckPost(c *gin.Context) {
// 	var form []ssq.SsqBall
// 	if c.ShouldBind(&form) == nil {

// 		// for _,boll := range form{
// 		// 	ssq.
// 		// }

// 	}

// }

// UserDoubaleBollPut - 用户新增投注
func UserDoubaleBollPut(c *gin.Context) {
	var form []models.UserDoubleBall
	if c.ShouldBind(&form) == nil {
		for _, obj := range form {
			database.LUCKDB.Clauses(clause.OnConflict{DoNothing: true}).Create(&obj)
		}
		c.String(200, "success")
	} else {
		c.String(403, "参数错误")
	}
}

// UserDoubaleBollDelete - 用户删除
func UserDoubaleBollDelete(c *gin.Context) {
	var form models.UserDoubleBall
	if c.ShouldBind(&form) == nil {
		database.LUCKDB.Delete(&models.UserDoubleBall{}, form.ID)

		c.String(200, "success")
	} else {
		c.String(403, "参数错误")
	}
}

// UserDoubaleBollGet - 获取用户投注情况
func UserDoubaleBollGet(c *gin.Context) {
	var (
		ret_un_open []*models.UserDoubleBall
		ret_open    []*models.UserDoubleBall
	)
	database.UpdateSsqAward()
	database.LUCKDB.Model(&models.UserDoubleBall{}).Where("is_open = ?", false).Order("qihao desc").Find(&ret_un_open)
	database.LUCKDB.Model(&models.UserDoubleBall{}).Where("is_open = ?", true).Order("qihao desc").Find(&ret_open)

	// todo is_open 为 false 的更新

	// 历史兑奖
	for _, obj := range ret_un_open {
		ssqball := ssq.SsqBall{Redboll: utils.BollStrToNum(obj.RedBall), Blueboll: utils.BollStrToNum(obj.BlueBall)}
		ret := ssq.DBAll.AwardCheck(&ssqball)
		obj.History = ret
	}

	// 历史兑奖
	for _, obj := range ret_open {
		ssqball := ssq.SsqBall{Redboll: utils.BollStrToNum(obj.RedBall), Blueboll: utils.BollStrToNum(obj.BlueBall)}
		ret := ssq.DBAll.AwardCheck(&ssqball)
		obj.History = ret
	}

	c.JSON(200, gin.H{"open": ret_open, "unopen": ret_un_open})
}

type APQihao struct {
	Qihao    string    `gorm:"primary_key;" json:"q"`
	OpenDate time.Time `json:"d"`
}

// 获取期号
func QiHaoGet(c *gin.Context) {
	var dbl []APQihao
	database.LUCKDB.Model(&models.DoubleBall{}).Select("qihao", "open_date").Order("open_date desc").Limit(100).Find(&dbl)
	c.JSON(200, dbl)
}
