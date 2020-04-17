package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func mainPageView(ctx *gin.Context) {
	var sonHaber []SonHaber
	err := getAllNews(&sonHaber)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"news": sonHaber,
		})
	}
}

func newsPageView(ctx *gin.Context) {
	var sonHaber SonHaber
	haberId := ctx.Param("haber_id")
	err := getNews(&sonHaber, haberId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.HTML(http.StatusOK, "news.html", gin.H{
			"news": sonHaber,
		})
	}
}
