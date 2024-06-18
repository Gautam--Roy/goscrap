package main

import (
	"encoding/xml"
	_fl "flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	var flag struct {
		Url        string
		Sitemap    string
		ReturnType string
		ApiKey     string
	}

	_fl.StringVar(&flag.ApiKey, "apiKey", "", "JinaAi reader API key")
	_fl.StringVar(&flag.Sitemap, "sitemap", "", "Url the sitemap")
	_fl.StringVar(&flag.ReturnType, "return", "html", "Return type. Eg: mdx, text")
	_fl.StringVar(&flag.Url, "url", "", "Url to scrap")

	_fl.Parse()

	if flag.ReturnType == "mdx" || flag.ReturnType == "text" {
		if flag.ApiKey == "" {
			log.Fatal("Please enter the apikey")
		}
	}

	if flag.Url != "" {
		scrapUrl(flag.Url, flag.ApiKey, flag.ReturnType)
		os.Exit(0)
	}

	if flag.Sitemap != "" {
		siteMapUrl := flag.Sitemap

		if isSiteMap(siteMapUrl) {
			fmt.Println("Sitemap confirmed")
		} else {
			log.Fatal("Invalid sitemap")
			os.Exit(0)
		}

		res, err := http.Get(siteMapUrl)
		if err != nil {
			panic(err)
		}

		var urlSet struct {
			URLs []struct {
				Loc string `xml:"loc"`
			} `xml:"url"`
		}

		if err = xml.NewDecoder(res.Body).Decode(&urlSet); err != nil {
			log.Fatal("There was an error while Decoding xml file at ScrapToFile: ", err.Error())
		}

		fmt.Println(urlSet.URLs[0])
		// TODO
		// Run scrap urls for each url inside goroutines
		os.Exit(0)
	}
}

func isSiteMap(url string) bool {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	return res.Header.Get("Content-Type") == "application/xml"
}

func scrapUrl(url string, api string, returnType string) {

	var header string
	var _url string

	switch returnType {
	case "mdx":
		header = "markdown"
		_url = "https://r.jina.ai/" + url
	case "text":
		header = "text"
		_url = "https://r.jina.ai/" + url
	default:
		header = "html"
		api = ""
		_url = url
	}

	req, err := http.NewRequest("GET", _url, nil)
	req.Header.Set("X-Return-Format", header)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", api))
	if err != nil {
		log.Fatal("Invalid url")
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error while processing request:", err.Error())
	}

	if res.StatusCode != 200 {
		log.Fatal("Non-success status code received: ", res.StatusCode)
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err.Error())
	}

	log.Print(string(bytes))

	res.Body.Close()
}
