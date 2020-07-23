package ssq

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/encoding/simplifiedchinese"

	"github.com/PuerkitoBio/goquery"

	"github.com/ider-zh/lottery/models"
)

func Ssq() {
	client := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := client.Get("https://kaijiang.500.com/shtml/ssq/03001.shtml")
	if err != nil {
		fmt.Println("http get error", err)
		return
	}
	retData := extractHtml(resp.Body)
	fmt.Printf("%+v", retData)

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("read error", err)
	// 	return
	// }
	// fmt.Println(string(body))

	// fmt.Println(string(body))
}

func extractHtml(body io.ReadCloser) *models.DoubleBall {

	var retData models.DoubleBall
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(body)

	if err != nil {
		log.Fatal(err)
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

	r, _ := regexp.Compile(`(\w{4})年(\w{1,2})月(\w{1,2})`)
	matched := r.FindStringSubmatch(txt)
	r, _ = regexp.Compile(`年|月`)
	retbyte := r.ReplaceAll([]byte(matched[0]), []byte("-"))
	fmt.Println(string(retbyte)) // true
	retData.OpenDate = string(retbyte)

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

	return &retData
}

func formatHtmlStr(str string) string {
	var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes([]byte(str))
	return strings.Trim(string(decodeBytes), " \n\t")
}
