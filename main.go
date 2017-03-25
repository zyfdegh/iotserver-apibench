package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	url = "http://104.199.219.138:8080/"
)

var count int64

func main() {
	n := flag.Int("n", 100, "loop count, default 100")
	flag.Parse()

	t1 := time.Now()

	errCh := make(chan *error)

	var wg sync.WaitGroup

	if *n <= 0 {
		*n = 100
	}

	for i := 0; i < *n; i++ {
		wg.Add(1)

		go getRoot(&wg, errCh)
	}

	go func() {
		for {
			select {
			case e := <-errCh:
				log.Printf("get root error: %v\n", *e)
			}
		}
	}()

	wg.Wait()

	t2 := time.Now()
	fmt.Printf("Time: %fs\n", t2.Sub(t1).Seconds())
	fmt.Printf("%d of %d success\n", count, *n)
}

func getRoot(wg *sync.WaitGroup, errCh chan *error) {
	defer wg.Done()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		errCh <- &err
		return
	}

	req.Close = true

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		errCh <- &err
		return

	}
	defer res.Body.Close()

	count++
	return
}
