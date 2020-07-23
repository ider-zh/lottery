package ssq

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

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
	retData.RedBallOrder = strings.Trim(txt, " ")

	txt = doc.Find(".cfont2 strong").Text()
	retData.Qihao = strings.Trim(txt, " ")

	txt = doc.Find(".span_right").Text()
	fmt.Println(txt)
	retData.OpenDate = strings.Trim(txt, " ")

	txt = doc.Find("span.cfont1:nth-child(1)").Text()
	txt = strings.Replace(txt, ",", "", -1)
	txt = strings.Replace(txt, "Ԫ", "", -1)
	number, _ := strconv.Atoi(strings.Trim(txt, " "))
	retData.TotalSales = int64(number)

	txt = doc.Find("span.cfont1:nth-child(2)").Text()
	txt = strings.Replace(txt, ",", "", -1)
	txt = strings.Replace(txt, "Ԫ", "", -1)
	number, _ = strconv.Atoi(strings.Trim(txt, " "))
	retData.Jackpot = int64(number)

	return &retData
}
