package schedule

import (
	"fmt"
	"log"
	"lottery/database"
	"lottery/internal/award/ssq"
	"lottery/utils/mail"
	"strings"

	"github.com/bamzi/jobrunner"
)

// 定时任务,一小时更新一次
func Schedule() {
	jobrunner.Start() // optional: jobrunner.Start(pool int, concurrent int) (10, 1)
	jobrunner.Schedule("14 * * * *", AwardCheckerJob{})
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

	udbs := database.UpdateSsqAward()
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
	for _, obj := range udbs {
		ret := obj.ToString()
		if len(ret) > 0 {
			awdCount += 1
			textA = append(textA, "期号："+obj.Qihao+" "+strings.Join(ret, ","))
		} else {
			textA = append(textA, "期号："+obj.Qihao+" 未中奖")
		}
	}
	if awdCount > 0 {
		subject = fmt.Sprintf("恭喜，有 %d 注幸运中奖", awdCount)
	}
	fmt.Println(subject, strings.Join(textA, "/n"))
	mail.NewSimpleTextMail(subject, strings.Join(textA, "。 "))
}
