package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

func getGoogleTrends() *http.Response {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	proxyURL, err := url.Parse(os.Getenv("PROXY_URL"))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	client := &http.Client{
		Transport: transport,
	}

	res, err := client.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=US")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return res
}

func ReadGoogleTrends() []byte {
	results := getGoogleTrends()
	data, err := ioutil.ReadAll(results.Body)
	if err != nil {
		return nil
	}
	return data
}
