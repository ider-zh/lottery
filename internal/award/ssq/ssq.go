/*
 * @Version: 0.0.1
 * @Author: ider
 * @Date: 2020-07-26 00:34:16
 * @LastEditors: ider
 * @LastEditTime: 2020-08-24 11:00:32
 * @Description:
 */
package ssq

import (
	"log"
	"math/rand"
	"sort"

	"lottery/utils"
)

type AwardCheck struct {
	Qihao    string
	Redboll  []int
	Blueboll int
}

type AwardResult struct {
	Q string
	A int
	B int
	C int
	D int
	E int
	F int
}

type SsqBall struct {
	Redboll  []int `form:"redboll" binding:"required"`
	Blueboll []int `form:"blueboll" binding:"required"`
}

func NewAwardCheck(qihao string, redboll []int, blueboll int) *AwardCheck {
	return &AwardCheck{
		Qihao:    qihao,
		Redboll:  redboll,
		Blueboll: blueboll,
	}
}

// 复式兑奖
func (c *AwardCheck) CheckMulitAward(ssqball *SsqBall) *AwardResult {
	var (
		rc, bc   int
		retaward AwardResult
	)
	// 蓝球比较
	for _, blue := range ssqball.Blueboll {
		if c.Blueboll == blue {
			bc = 1
			break
		}
	}

	retred := utils.IntersectionInt(&c.Redboll, &ssqball.Redboll)
	rc = len(*retred)
	rcn := len(ssqball.Redboll) - rc
	bcn := len(ssqball.Blueboll) - bc
	retaward.Q = c.Qihao
	retaward.A = utils.Combine(rc, 6) * utils.Combine(bc, 1)
	retaward.B = utils.Combine(rc, 6) * utils.Combine(bcn, 1)
	retaward.C = utils.Combine(rc, 5) * utils.Combine(rcn, 1) * utils.Combine(bc, 1)
	retaward.D = utils.Combine(rc, 5)*utils.Combine(rcn, 1)*utils.Combine(bcn, 1) + utils.Combine(rc, 4)*utils.Combine(rcn, 2)*utils.Combine(bc, 1)
	retaward.E = utils.Combine(rc, 4)*utils.Combine(rcn, 2)*utils.Combine(bcn, 1) + utils.Combine(rc, 3)*utils.Combine(rcn, 3)*utils.Combine(bc, 1)
	retaward.F = utils.Combine(rc, 2)*utils.Combine(rcn, 4)*utils.Combine(bc, 1) + utils.Combine(rc, 1)*utils.Combine(rcn, 5)*utils.Combine(bc, 1) + utils.Combine(rcn, 6)*utils.Combine(bc, 1)
	return &retaward
}

type RandomWorksr struct {
	generator *rand.Rand
}

func NewRandomWorksr(seed int64) *RandomWorksr {
	return &RandomWorksr{
		generator: rand.New(rand.NewSource(seed)),
	}
}

func (c *RandomWorksr) GetTicket() *SsqBall {
	set := make(map[int]bool)
	var redboll []int

	for {
		if len(redboll) == 6 {
			break
		}

		ret := int(c.generator.Intn(31))

		if _, ok := set[ret]; ok {
			continue
		}
		redboll = append(redboll, ret)

		set[ret] = true
	}
	sort.Ints(redboll)
	blueboll := []int{c.generator.Intn(16)}
	return &SsqBall{Redboll: redboll, Blueboll: blueboll}
}

func work(seed int64) {
	// rw := NewRandomWorksr(time.Now().UnixNano())
	rw := NewRandomWorksr(seed)

	ac := NewAwardCheck("202098", []int{9, 15, 18, 21, 23, 26}, 8)
	// startT := time.Now()
	for i := 0; i < 1000000; i++ {
		ssqboll := rw.GetTicket()
		if ret := ac.CheckMulitAward(ssqboll); ret.A > 0 || ret.B > 0 {
			log.Println(seed, i, ret, ssqboll)
		}
	}
	// tc := time.Since(startT) //计算耗时
	// fmt.Printf("time cost = %v\n", tc)
}

func main() {
	for i := 0; i < 10000000; i++ {
		work(int64(i))
	}
}
