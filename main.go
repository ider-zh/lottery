package main

import (
	"fmt"
	"strconv"
	"strings"

	"lottery/config"
	"lottery/database"
	"lottery/internal/award/ssq"
	"time"

	"github.com/emirpasic/gods/sets/hashset"
)

func main() {
	// ssq.Ssq()

	// ssq.SsqSchedule()
	set := hashset.New()   // empty
	set.Add("1")           // 1
	set.Add(2, 2, 3, 4, 5) // 3, 1, 2, 4, 5 (random order, duplicates ignored)
	set.Remove(4)          // 5, 3, 2, 1 (random order)
	// set.Remove(2, 3)       // 1, 5 (random order)
	set.Contains(1)    // true
	set.Contains(1, 5) // true
	set.Contains(1, 6) // false
	// ret := set.Values() // []int{5,1} (random order)
	set.Clear() // empty
	set.Empty() // true
	set.Size()  // 0
	// fmt.Printf("%+v", set.Contains("1", 5))
	m := make(map[int]uint8)
	m[12] |= (1 << 1)
	// a |=
	// fmt.Printf("%+v", m)
	// a := []int{3, 4, 6}
	// b := []int{3, 4, 6, 1, 5}
	// ret := util.IntersectionInt(&a, &b)
	// fmt.Println(*ret)
	// 7 5
	// for i := range 9{

	// }
	// fmt.Println(util.Combine(1, 4))
	// ac := ssq.NewAwardCheck("202098", []int{9, 15, 18, 21, 23, 26}, 8)

	// ret := ac.CheckMulitAward()
	// fmt.Printf("%+v", ret)

	ss := "13 12 51 51 04"
	var retInt []int
	for _, s := range strings.Split(ss, " ") {
		ss, _ := strconv.Atoi(s)
		retInt = append(retInt, ss)
	}
	fmt.Println(retInt)

	ssqball := ssq.SsqBall{Redboll: []int{11, 12, 19, 27, 31, 32}, Blueboll: []int{15}}

	cfg := config.GetConfig()
	database.NewLuckDBConn(cfg.PgConn)

	ssq.NewDoubleBollAll()
	// ssq.DBAll.Update()
	ret := ssq.DBAll.AwardCheck(&ssqball)
	// fmt.Printf("%+v", ret)

	start := time.Now()
	ssqball = ssq.SsqBall{Redboll: []int{1, 3, 7, 21, 27, 32}, Blueboll: []int{3}}
	ret = ssq.DBAll.AwardCheck(&ssqball)
	// fmt.Printf("%+v", ret)
	fmt.Println("aa")
	for _, obj := range *ret {
		fmt.Println(obj)
	}
	fmt.Println("aa")
	// for _, obj := range ssq.DBAll.AwardCheckS {

	// 	ret = ssq.DBAll.AwardCheck(&ssq.SsqBall{obj.Redboll, []int{obj.Blueboll}})
	// 	for _, bb := range *ret {
	// 		if bb.Q != obj.Qihao {

	// 			// fmt.Printf("%+v", bb)
	// 			fmt.Println(obj.Qihao, bb)
	// 		}
	// 	}
	// }

	//输出执行时间，单位为毫秒。
	fmt.Println(time.Since(start))

	fmt.Println("test split data")
	sss := "08 17 24 26 27 31"

	var outs = []int{}
	for _, s := range strings.Split(sss, " ") {
		j, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		outs = append(outs, j)
	}

	fmt.Println(outs)
	// ssqball = ssq.SsqBall{Redboll: []int{1, 3, 7, 21, 27, 32}, Blueboll: []int{3}}

}
