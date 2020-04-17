package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB
var err error

func main() {
	DB, err = gorm.Open("sqlite3", "db.sqlite3")

	errorHandler(err)

	defer DB.Close()

	DB.AutoMigrate(&SonHaber{})

	go sonHaberComMainPageScraper() // run first time
	go scraperTask()                // run cron scraper

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", mainPageView)
	router.GET("/news/:haber_id", newsPageView)

	errorHandler(router.Run(":8080"))

}
