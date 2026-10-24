package main

import (
	"fmt"
	"net/http"
	"time"
)

type Status struct {
	url         string
	code        int
	elapsedTime time.Duration
}

func main() {
	urls := []string{"google.com", "microsoft.com", "reddit.com", "uscis.org"}
	var statusResponse []Status

	client := http.Client{Timeout: 5 * time.Second}
	for _, url := range urls {
		start := time.Now()
		status, err := client.Get(fmt.Sprintf("https://%s", url))
		if err != nil {
			fmt.Printf("error getting http status for %v\nError: %v\n", url, err)
			continue
		}

		defer status.Body.Close()
		elapsed := time.Since(start)
		statusResponse = append(statusResponse, Status{url: url, code: status.StatusCode, elapsedTime: elapsed})
	}

	for _, s := range statusResponse {
		fmt.Printf("URL: %v\nStatus Code: %v\nElapsed Time: %v\n\n", s.url, s.code, s.elapsedTime)
	}
}
