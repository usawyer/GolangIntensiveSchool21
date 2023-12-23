package fanin

import "sync"

func Multiplex(input ...<-chan interface{}) <-chan interface{} {
	fanIn := make(chan interface{})
	var wg sync.WaitGroup

	send := func(c <-chan interface{}) {
		for n := range c {
			fanIn <- n
		}
		wg.Done()
	}

	wg.Add(len(input))
	for _, c := range input {
		go send(c)
	}

	go func() {
		wg.Wait()
		close(fanIn)
	}()

	return fanIn
}
