package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gocolly/colly/v2"
	"github.com/hnucamendi/ws-colly_lambda/scrape"
)

func HandleRequest() string {
	c := colly.NewCollector(colly.Async(true), colly.UserAgent("ws-colly"))
	s := scrape.NewScrape()

	json, err := s.ScrapeTopHeadLines(c, scrape.CNNConfig())
	if err != nil {
		fmt.Println(err)
	}

	return json

}

func main() {
	lambda.Start(HandleRequest)
}
