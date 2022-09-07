package main

import (
	"fmt"
	"time"
)

func Goroutines_Channel() {
	go sexyCount("mingyu") 
	// if wanna excute go-routine function, just add keyword "go"
	// this is work like generator in javascript
	go sexyCount("shin")
	// but the go-routine function live in main function's life cycle

	// Channel
	c := make(chan bool) // how to make channel in go
	// this will be used like data pipe
	people := [2]string{"hi","hihi"}
	for _, person := range people{
		go isSexy(person, c)// how to use channel in go routine
	}
	for range people{
		fmt.Println(<-c) // blocking operation
	}

}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

func isSexy(person string, c chan bool){
	time.Sleep(time.Second * 5)
	c <- true // how to send data to channel in go routine
}