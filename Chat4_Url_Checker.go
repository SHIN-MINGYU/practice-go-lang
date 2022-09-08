package main

import (
	"errors"
	"fmt"
	"net/http"
)

type result struct {
	url string
	status string
}

func Url_Checker() {
	results := make(map[string]string)// same map[string]string{}
	// uninitailized map, cant input some value
	c := make(chan result)
	urls := []string{
		"https://www.google.co.kr/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls{
		go hitUrls(url, c)
	}

	for range urls{
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results{
		fmt.Println(url, status)
	}

}


var errReqFailed =errors.New("request failed")

func hitUrls(url string, c chan result) {
	// it is better what write explicitly the channels direction 
	fmt.Println("Checking : ", url)
	res, err := http.Get(url)
	status := "OK"
	if(err != nil || res.StatusCode >= 400){
		status = "FAILED"
	}
	c <- result{url : url, status: status}
}