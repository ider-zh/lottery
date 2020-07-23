package crawler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
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
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	txt := doc.Find(".ball_blue").Text()
	log.Println(txt)
	// Find the review items
	doc.Find(".ball_red").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Text()
		log.Println(band)
	})

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("read error", err)
	// 	return
	// }
	// fmt.Println(string(body))
}
