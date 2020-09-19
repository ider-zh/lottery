package schedule

import (
	"fmt"
	"log"
	"lottery/database"
	"lottery/internal/award/ssq"
	"lottery/models"
	"lottery/utils/mail"
	"strings"

	"github.com/bamzi/jobrunner"
	"github.com/emirpasic/gods/sets/treeset"
)

// 定时任务,一小时更新一次
func Schedule() {
	jobrunner.Start() // optional: jobrunner.Start(pool int, concurrent int) (10, 1)
	jobrunner.Schedule("15 1,2,12,22,23 * * *", AwardCheckerJob{})
	//	jobrunner.Schedule("* * * * *", AwardCheckerJob{})
}

// Job Specific Functions
type AwardCheckerJob struct {
	// filtered
}

// AwardCheckerJob.Run() will get triggered automatically.
func (e AwardCheckerJob) Run() {
	log.Println("start crawler")
	// 更新数据
	ssq.NewDoubleBollAll()

	udbs := ssq.UpdateSsqAward()
	// 没有未开奖的
	if len(udbs) == 0 {
		return
	}
	// 未开奖的主动开奖
	var (
		subject  = "这次没有中奖"
		awdCount int
		textA    []string
	)
	// 输出开奖号码
	set := treeset.NewWithStringComparator()
	for _, obj := range udbs {
		if !set.Contains(obj.Qihao) {
			var adBall models.DoubleBall
			database.LUCKDB.Where("qihao = ?", obj.Qihao).First(&adBall)
			textA = append(textA, "开奖号码：期号："+obj.Qihao+", 红球："+adBall.RedBall+"，蓝球："+adBall.BlueBall)
			set.Add(obj.Qihao)
		}
	}

	// 统计中奖情况
	for _, obj := range udbs {
		ret := obj.ToString()
		if len(ret) > 0 {
			awdCount += 1
			textA = append(textA, "期号："+obj.Qihao+" ，红："+obj.RedBall+" ，蓝："+obj.BlueBall+"。"+strings.Join(ret, " , "))
		} else {
			textA = append(textA, "期号："+obj.Qihao+" ，红："+obj.RedBall+" ，蓝："+obj.BlueBall+"。未中奖！")
		}
	}

	// 输出中奖情况或者连续未中次数
	if awdCount > 0 {
		subject = fmt.Sprintf("恭喜，有 %d 注幸运中奖！！！", awdCount)
	} else {
		// 统计连续未中奖次数
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
		subject = subject + fmt.Sprintf("，连续 %d 注未中奖！！！", na)
	}

	fmt.Println(subject, strings.Join(textA, "/n"))
	mail.NewSimpleTextMail(subject, strings.Join(textA, "\r\n"))
}
