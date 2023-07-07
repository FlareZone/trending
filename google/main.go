package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/FlareZone/trending/google/controller"
	"github.com/FlareZone/trending/google/model"
)

func main() {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "pong",
		})
	})

	router.GET("/trending", func(c *gin.Context) {
		area := c.Query("area")
		data := controller.ReadGoogleTrends(area)
		var rssFeed model.Rss
		err := xml.Unmarshal(data, &rssFeed)
		if err != nil {
			fmt.Println(err.Error())
			c.JSON(500, gin.H{
				"error": "Failed to parse XML data",
			})
			return
		}

		var trends []gin.H
		for i := range rssFeed.Channel.Items {
			trend := gin.H{
				"title":       rssFeed.Channel.Items[i].Title,
				"traffic":     rssFeed.Channel.Items[i].Traffic,
				"url":         rssFeed.Channel.Items[i].Link,
				"date":        rssFeed.Channel.Items[i].PublishedDate,
				"picture_url": rssFeed.Channel.Items[i].PictureURL,
				"news_items":  rssFeed.Channel.Items[i].NewsItems,
			}
			trends = append(trends, trend)
		}

		c.JSON(200, gin.H{
			"trends": trends,
		})
	})

	router.Run(":8080")

	wMsg := "Google Trends Cli Application"
	fmt.Printf("%s\n", wMsg)
}

func logItem(title string, traffic string, url string, date string, picture string) {
	log.Printf("Trend Title: %s Trend Traffic :%s  Trend Url: %s Trend Date: %s Trend Picture:%s \n", title, traffic, url, date, picture)
}

func logNewsItem(newsItems []model.NewsItem) {
	for _, newsItem := range newsItems {
		log.Println(newsItem)
	}
}
