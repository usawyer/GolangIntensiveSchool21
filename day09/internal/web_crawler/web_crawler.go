package main

import (
	"context"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
)

func main() {
	urlQueue := make(chan string)
	foundCh := make(chan *string)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\nCrawling interrupted. Stopping gracefully...")
		cancel()
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		urlQueue <- "https://www.example.com"
	}()

	wg.Add(1)
	go crawlWorker(ctx, urlQueue, foundCh, &wg)

	wg.Wait()
	close(foundCh)
	close(urlQueue)
	close(sigChan)
}

func crawlWorker(ctx context.Context, urlQueue chan string, foundCh chan *string, wg *sync.WaitGroup) {
	defer wg.Done()
	processedURLs := make(map[string]bool)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task cancelled")
			return
		default:
			url, ok := <-urlQueue
			log.Println("Сrawling: ", url)
			if !ok {
				return
			}

			if processedURLs[url] {
				continue
			}

			body, err := fetch(url)
			if err != nil {
				continue
			}

			go func() {
				foundCh <- &body
			}()

			links := extractLinks(body)
			for _, link := range links {
				wg.Add(1)
				go func(link string) {
					defer wg.Done()
					urlQueue <- link
				}(link)
			}

			processedURLs[url] = true
		}
	}
}

func fetch(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP request failed with status: %v", response.Status)
	}

	body, err := io.ReadAll(response.Body)

	return string(body), nil
}

func extractLinks(body string) []string {
	links := make([]string, 0)
	tokenizer := html.NewTokenizer(strings.NewReader(body))
	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			return links
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" && len(attr.Val) != 0 && attr.Val[0] == 'h' {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}
}
