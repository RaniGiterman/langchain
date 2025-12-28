package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type GetPage struct{}

func (c GetPage) Name() string {
	return "get_page_html"
}

func (c GetPage) Description() string {
	return "given URL input, returns page HTML code"
}

func (c GetPage) Call(ctx context.Context, input string) (string, error) {
	// custom logic!
	// Make the GET request
	resp, err := http.Get(input)
	if err != nil {
		log.Fatal("Error fetching URL: ", err)
	}
	// Ensure the response body is closed to prevent resource leaks
	defer resp.Body.Close()

	// Check if the request was successful (status code 200 OK)
	if resp.StatusCode != http.StatusOK {
		log.Fatal("Error: Status code is not OK: ", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
	}

	return string(body), nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
