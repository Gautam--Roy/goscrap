package main

import (
	_fl "flag"
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
		header := ""
		url := flag.Url
		switch flag.ReturnType {
		case "mdx":
			header = "markdown"
			url = "https://r.jina.ai/" + flag.Url
		case "text":
			header = "text"
			url = "https://r.jina.ai/" + flag.Url
		default:
			header = "html"
		}

		req, err := http.NewRequest("GET", url, nil)
		req.Header.Set("X-Return-Format", header)
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
		os.Exit(0)
	}

}
