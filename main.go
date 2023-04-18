package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// API returns the following - input is based off ?input= arg
// {"page":"words","input":"word1","words":["word1"]}
type Words struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func main() {
	// Get args from commandline
	// ./http-get <URL>
	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usage: ./http-get <url>\n")
		os.Exit(1)
	}

	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("URL format is invalid: %s\n ", err)
		os.Exit(1)
	}

	// do http GET of URL provided by args[1]
	response, err := http.Get(args[1])

	// Error if unable to GET URL
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	// Read all content of response stream
	body, err := io.ReadAll(response.Body)

	// Error if unable to get body of response
	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode != 200 {
		fmt.Printf("Invalid output (HTTP Code %d): %s\n", response.StatusCode, body)
		os.Exit(1)
	}

	var words Words

	err = json.Unmarshal(body, &words)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Page: %s\nWords %v\n", words.Page, strings.Join(words.Words, ", "))
}
