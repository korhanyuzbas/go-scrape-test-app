package main

import (
	"github.com/jinzhu/gorm"
)

type SonHaber struct {
	gorm.Model
	HaberID int    `form:"haberid" binding:"required"`
	Title   string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
	Image   string `form:"image"`
}

func getAllNews(news *[]SonHaber) (err error) {
	if err = DB.Order("haber_id desc").Find(news).Error; err != nil {
		return err
	}
	return nil
}

func getNews(news *SonHaber, id string) (err error) {
	if err = DB.Where("haber_id = ?", id).First(news).Error; err != nil {
		return err
	}
	return nil
}
