package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type NewsAPIResponse struct {
	Status       string `json:"status"`
	TotalResults int    `json:"totalResults"`
	Articles     []struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		URL         string `json:"url"`
	} `json:"articles"`
}

func main() {
	apiKey := "89d19cedb5c84dd391f2bb00f2099ad6"
	endpoint := "https://newsapi.org/v2/everything?q=tesla&from=2023-07-16&sortBy=publishedAt&apiKey=" + apiKey

	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("API request error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Response body read error:", err)
		return
	}

	var apiResponse NewsAPIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println("JSON decoding error:", err)
		return
	}

	for _, article := range apiResponse.Articles {
		fmt.Println("Title:", article.Title)
		fmt.Println("Description:", article.Description)
		fmt.Println("URL:", article.URL)
		fmt.Println()
	}
}
