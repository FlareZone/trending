package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

// Area area
type Area struct {
	AREA string
}

func getGoogleTrends(area string) *http.Response {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	proxyURL := os.Getenv("PROXY_URL")
	var transport *http.Transport

	if proxyURL != "" {
		proxyURL, err := url.Parse(proxyURL)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		transport = &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	} else {
		transport = &http.Transport{}
	}

	client := &http.Client{
		Transport: transport,
	}

	res, err := client.Get(`https://trends.google.com/trends/trendingsearches/daily/rss?geo=` + area)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return res
}

// ReadGoogleTrends get google trending news
func ReadGoogleTrends(area string) []byte {
	results := getGoogleTrends(area)
	data, err := ioutil.ReadAll(results.Body)
	if err != nil {
		return nil
	}
	return data
}
