package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\nCrawling interrupted. Stopping gracefully...")
		cancel()
	}()

	wg.Add(1)
	go performTask(ctx, &wg)

	wg.Wait()

}

func performTask(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Task cancelled")
			return
		default:
			fmt.Println("Performing task...")
			time.Sleep(time.Second)
		}
	}
}
