package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func scrapeSonHaberDetailPage(url string, haberId int) {
	res, err := http.Get(url)
	errorHandler(err)

	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	errorHandler(err)

	text := strings.TrimSpace(document.Find(".haber_metni").Text())
	title := document.Find(".haber_baslik").Text()
	image, imageExists := document.Find("div.drimg").Find("img").Attr("src")

	if !imageExists {
		image = ""
	}

	DB.Create(&SonHaber{
		HaberID: haberId,
		Title:   title,
		Content: text,
		Image:   image,
	})
}

func processSonHaberLinks(index int, element *goquery.Selection) {
	href, exists := element.Attr("href")
	if exists && strings.HasPrefix(href, "/haber/") {
		splitLink := strings.Split(href, "-")
		haberId, _ := strconv.Atoi(strings.Replace(splitLink[len(splitLink)-1], "/", "", -1))
		href = "https://www.sondakika.com" + href
		var sonHaber SonHaber
		if err := DB.Where("haber_id = ?", haberId).First(&sonHaber).Error; gorm.IsRecordNotFoundError(err) {
			scrapeSonHaberDetailPage(href, haberId)
		}
	}

}

func sonHaberComMainPageScraper() {
	res, err := http.Get("https://www.sondakika.com/")
	errorHandler(err)

	defer res.Body.Close()

	document, err := goquery.NewDocumentFromReader(res.Body)
	errorHandler(err)

	document.Find("a.resim").Each(processSonHaberLinks)
}

func scraperTask() {
	ticker := time.NewTicker(5 * time.Minute)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				sonHaberComMainPageScraper()
			}
		}
	}()
}
