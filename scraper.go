package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"topicalert/ResultStruct"
)

func scrape(topic string) {
	topic = strings.ReplaceAll(topic, " ", "%20")
	var apiKey string = os.Getenv("newsdata_key")
	var scraperApiUrl string = "https://newsdata.io/api/1/news?apikey=" + apiKey + "&q=" + topic + "&language=en&timeframe=24"
	//fmt.Println(scraperApiUrl)
	resp, err := http.Get(scraperApiUrl)
	if err != nil {
		log.Fatal(err)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	scraperResult := ResultStruct.ScraperResult{}
	error := json.Unmarshal(b, &scraperResult)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(scraperResult)

}
