package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

const (
	url = "http://localhost:3000/user/signup"
)

func main() {
	t1 := time.Now()

	usernameCh := make(chan string)
	errCh := make(chan *error)

	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		username := fmt.Sprintf("user%03d", i)
		password := fmt.Sprintf("secret%03d", i)
		email := fmt.Sprintf("%s@email.com", username)

		go register(username, password, email, &wg, usernameCh, errCh)
	}

	go func() {
		for {
			select {
			case r := <-usernameCh:
				log.Printf("%s OK\n", r)
			case e := <-errCh:
				log.Printf("register error: %v\n", e)
			}
		}
	}()

	wg.Wait()

	t2 := time.Now()
	fmt.Printf("Time: %fs\n", t2.Sub(t1).Seconds())
}

func register(username, password, email string, wg *sync.WaitGroup, usernameCh chan string, errCh chan *error) {
	defer wg.Done()

	params := fmt.Sprintf("username=%s&email=%s&password=%s", username, email, password)
	payload := strings.NewReader(params)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		errCh <- &err
		return
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		errCh <- &err
		return
	}

	defer res.Body.Close()

	usernameCh <- username
	return
}
