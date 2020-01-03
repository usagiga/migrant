package main

import (
	"flag"
	"fmt"
	"github.com/upamune/go-esa/esa"
	"github.com/usagiga/migrant/entity"
	"github.com/usagiga/migrant/model"
	"log"
	"time"
)

func main() {
	var teamName string
	var apiKey string
	var maxPostsPerWindowSize int
	var postsWindowSize int

	// Read flag
	flag.StringVar(&teamName,"team", "", "Team Name of {your-team}.esa.io")
	flag.StringVar(&apiKey,"key", "", "API Key generated from {your-team}.esa.io")
	flag.IntVar(&maxPostsPerWindowSize, "max-posts", 70, "The number of posts within window size")
	flag.IntVar(&postsWindowSize, "window-size", 15, "The window size of {your-team}.esa.io's post rate limits. Wrote by minutes only.")
	flag.Parse()

	if teamName == "" {
		log.Fatalln("You must set \"team\" to esa.io team name.")
		return
	}

	if apiKey == "" {
		log.Fatalln("You must set \"key\" to esa.io API Key.")
		return
	}

	// Post to esa.io
	client := esa.NewClient(apiKey)
	crawler := model.NewCrawler(entity.CrawlerType_EsaOffline, nil)

	loopCount := 0
	for post := range crawler.CrawlIter() {
		// Rate limit
		if loopCount != 0 && loopCount%maxPostsPerWindowSize == 0 {
			time.Sleep(time.Duration(postsWindowSize) * time.Minute)
		}

		// Post
		res, err := client.Post.Create(teamName, *post)
		fmt.Println(" - Case ", loopCount+1, "\nres: ", res, "\nerr: ", err)

		loopCount++
	}
}
