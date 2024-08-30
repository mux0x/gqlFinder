package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func fetchURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func processURL(url string, re *regexp.Regexp) {
	content, err := fetchURL(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to fetch %s: %v\n", url, err)
		return
	}

	matches := re.FindAllString(content, -1)
	for _, match := range matches {
		fmt.Println(match)
	}
}

func main() {
	re := regexp.MustCompile(`(mutation\s|query\s)[A-Za-z0-9_]+[^\(]\([^\)]+\)[^` + "`" + `'"]+`)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		url := scanner.Text()
		processURL(url, re)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}

