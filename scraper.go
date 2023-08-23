package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func scrape(topic string) {
	topic = strings.ReplaceAll(topic, " ", "%20")
	res, err := http.Get("https://news.google.com/search?q=" + topic)
	if err != nil {
		return
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	fmt.Println(doc.Find(".lBwEZb.BL5WZb.GndZbb"))

}
