package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type Area struct {
	AREA string
}

func getGoogleTrends() *http.Response {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	// read json file to get tending data
	content, err := ioutil.ReadFile("./area.json")
	if err != nil {
		log.Fatal("Read area.json failed error:", err)
	}

	// Now let's unmarshall the data into `payload`
	var payload = make([]Area, 0)
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("json.Unmarshal() failed error:", err)
	}

	log.Printf("origin: %s\n", payload[1].AREA)

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

	for i := 0; i < len(payload); i++ {
		res, err := client.Get(`https://trends.google.com/trends/trendingsearches/daily/rss?geo=` + payload[i].AREA)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		return res
	}
	return nil
}

func ReadGoogleTrends() []byte {
	results := getGoogleTrends()
	data, err := ioutil.ReadAll(results.Body)
	if err != nil {
		return nil
	}
	return data
}
