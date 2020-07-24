package ssq

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"

	"github.com/PuerkitoBio/goquery"

	"github.com/ider-zh/lottery/models"
)

var Client *http.Client

func init() {
	Client = &http.Client{
		Timeout: 30 * time.Second,
	}
}

func GetSsq(Qihao string) (*models.DoubleBall, error) {

	resp, err := Client.Get("https://kaijiang.500.com/shtml/ssq/" + Qihao + ".shtml")
	if err != nil {
		fmt.Println("http get error", err)
		return nil, err
	}
	retData, err := extractHtml(resp.Body)
	return retData, err

}

func extractHtml(body io.ReadCloser) (*models.DoubleBall, error) {

	var retData models.DoubleBall
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	txt := doc.Find(".ball_blue").Text()
	retData.BlueBall = txt

	// Find the review items
	doc.Find(".ball_red").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		ball := s.Text()
		retData.RedBall = retData.RedBall + " " + ball
	})
	retData.RedBall = strings.Trim(retData.RedBall, " ")

	txt = doc.Find("div.kjxq_box02  table:nth-child(1)  table tr:nth-child(2) > td:nth-child(2)").Text()
	retData.RedBallOrder = strings.Trim(txt, " \n\t")

	txt = doc.Find(".cfont2 strong").Text()
	retData.Qihao = strings.Trim(txt, " ")

	txt = doc.Find(".span_right").Text()
	txt = formatHtmlStr(txt)
	fmt.Println(txt)

	r, _ := regexp.Compile(`(\w{4}).*?(\w{1,2}).*?(\w{1,2})`)
	matched := r.FindStringSubmatch(txt)
	if len(matched) < 4 {
		return nil, errors.New("matched less 4")
	}
	daystr := fmt.Sprintf("%s-%s-%s", matched[1], matched[2], matched[3])
	fmt.Println(daystr) // true
	retData.OpenDate = daystr

	txt = doc.Find("span.cfont1:nth-child(1)").Text()
	txt = formatHtmlStr(txt)
	txt = strings.Replace(txt, ",", "", -1)
	txt = strings.Replace(txt, "元", "", -1)
	number, _ := strconv.Atoi(strings.Trim(txt, " "))
	retData.TotalSales = int64(number)

	txt = doc.Find("span.cfont1:nth-child(2)").Text()
	txt = formatHtmlStr(txt)
	txt = strings.Replace(txt, ",", "", -1)
	txt = strings.Replace(txt, "元", "", -1)
	number, _ = strconv.Atoi(strings.Trim(txt, " "))
	retData.Jackpot = int64(number)

	doc.Find(`.kj_tablelist02 tr[align="center"]`).Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		awardflag := formatHtmlStr(s.Find("td:nth-child(1)").Text())
		switch awardflag {
		case "一等奖":
			awardcount := formatHtmlStr(s.Find("td:nth-child(2)").Text())
			number, err := strconv.Atoi(strings.Trim(awardcount, " "))
			if err == nil {
				retData.OneCount = int64(number)
			}
			awardcount = formatHtmlStr(s.Find("td:nth-child(3)").Text())
			number, err = strconv.Atoi(strings.Replace(awardcount, ",", "", -1))
			if err == nil {
				retData.OneAward = int64(number)
			}
		case "二等奖":
			awardcount := formatHtmlStr(s.Find("td:nth-child(2)").Text())
			number, err := strconv.Atoi(strings.Trim(awardcount, " "))
			if err == nil {
				retData.TwoCount = int64(number)
			}
			awardcount = formatHtmlStr(s.Find("td:nth-child(3)").Text())
			number, err = strconv.Atoi(strings.Replace(awardcount, ",", "", -1))
			if err == nil {
				retData.TwoAward = int64(number)
			}
		case "三等奖":
			awardcount := formatHtmlStr(s.Find("td:nth-child(2)").Text())
			number, err := strconv.Atoi(strings.Trim(awardcount, " "))
			if err == nil {
				retData.ThreeCount = int64(number)
			}
			awardcount = formatHtmlStr(s.Find("td:nth-child(3)").Text())
			number, err = strconv.Atoi(strings.Replace(awardcount, ",", "", -1))
			if err == nil {
				retData.ThreeAward = int64(number)
			}
		}
	})

	return &retData, nil
}

// 找到双色球所有期号，返回 map
func extractSsqList(body io.ReadCloser) *[]string {

	var retData []string

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".iSelectList a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		Qihao := s.Text()
		retData = append(retData, formatHtmlStr(Qihao))
	})

	return &retData
}

func getSsqQihao() {
	resp, err := Client.Get("https://kaijiang.500.com/shtml/ssq/03001.shtml")
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	retData := extractSsqList(resp.Body)
	fmt.Printf("%+v", retData)
}

func formatHtmlStr(str string) string {
	var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes([]byte(str))
	return strings.Trim(string(decodeBytes), " \n\t")
}

func SsqSchedule() {
	var (
		retballsfinsh []models.DoubleBall
		qihaotodo     []string
		finshQihao    map[string]bool
	)
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(dir)

	filename := dir + "/file/ssq.json"
	f, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		f1, err1 := os.Create(filename) //创建文件

		log.Println(err1)
		f1.Close()
	} else {

		content, _ := ioutil.ReadAll(f)
		json.Unmarshal([]byte(content), &retballsfinsh)
		f.Close()
	}
	log.Println("已保存：", len(retballsfinsh))
	resp, err := Client.Get("https://kaijiang.500.com/shtml/ssq/03001.shtml")
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	retData := extractSsqList(resp.Body)
	// fmt.Printf("%+v", retData)

	// 获取已存储的期号
	finshQihao = make(map[string]bool)
	for _, obj := range retballsfinsh {
		finshQihao[obj.Qihao] = true
	}

	// 获取没有采集的期号
	for _, Qihao := range *retData {
		if _, ok := finshQihao[Qihao]; !ok {
			qihaotodo = append(qihaotodo, Qihao)
		}
	}

	for _, Qihao := range qihaotodo {
		log.Println(Qihao)
		dbball, err := GetSsq(Qihao)
		if err != nil {
			log.Println(err)
			continue
		}
		retballsfinsh = append(retballsfinsh, *dbball)

		jsons, _ := json.Marshal(retballsfinsh)
		err = ioutil.WriteFile(filename, jsons, 0666)
		if err != nil {
			log.Fatal(err)
		}

	}

}
