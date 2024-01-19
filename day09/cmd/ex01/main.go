package main

import (
	"bytes"
	"context"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Found struct {
	Body []byte
	Url  string
}

const maxRoutines = 8

var (
	fetcher = 0
	extract = 0
	mtx     sync.Mutex
)

func main() {
	urlQueue := make(chan string, 1)
	foundCh := make(chan Found, 1)

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
		urlQueue <- "https://ya.ru"
		//urlQueue <- "http://localhost:8888"
	}()

	wg.Add(1)
	time.Sleep(time.Millisecond)
	go crawlWorker(ctx, urlQueue, foundCh, &wg)

	wg.Wait()
	close(foundCh)
	close(urlQueue)
	close(sigChan)
}

func crawlWorker(ctx context.Context, urlQueue chan string, foundCh chan Found, wg *sync.WaitGroup) {
	defer wg.Done()
	processedURLs := make(map[string]bool)
	for {
		time.Sleep(time.Millisecond * 100)
		select {
		case url, ok := <-urlQueue:
			if ok && !processedURLs[url] && fetcher < maxRoutines {
				processedURLs[url] = true
				fetcher++
				log.Println("Сrawling: ", url)
				go func() {
					fetch(url, foundCh)
				}()
			}
		case found, ok := <-foundCh:
			if ok && extract < maxRoutines {
				extract++
				go func() {
					extractLinks(urlQueue, found.Body, ctx)
				}()
			}
		case <-ctx.Done():
			fmt.Println("Task cancelled")
			return
		default:
			if fetcher == 0 && extract == 0 {
				log.Println("I'm done, ciao")
				return
			}
		}
	}
}

func fetch(url string, foundCh chan Found) {
	defer func() {
		fetcher--
	}()

	client := &http.Client{
		Timeout: time.Second * 5,
	}

	response, err := client.Get(url)

	if err != nil {
		log.Println(err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("%v: HTTP request failed with status: %v\n", url, response.Status)
		return
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return
	}

	foundCh <- Found{
		Body: body,
		Url:  url,
	}

}

func extractLinks(urlQueue chan string, body []byte, ctx context.Context) {
	defer func() {
		extract--
	}()

	tokenizer := html.NewTokenizer(bytes.NewReader(body))
	for {
		select {
		case <-ctx.Done():
			return
		default:
			tokenType := tokenizer.Next()

			switch tokenType {
			case html.ErrorToken:
				return
			case html.StartTagToken, html.SelfClosingTagToken:
				token := tokenizer.Token()
				if token.Data == "a" {
					for _, attr := range token.Attr {
						if attr.Key == "href" && len(attr.Val) != 0 && attr.Val[0] == 'h' {
							select {
							case <-ctx.Done():
								return
							default:
								ok := mtx.TryLock()
								if ok {
									urlQueue <- attr.Val
									mtx.Unlock()
								}
							}
						}
					}
				}
			}
		}
	}
}
