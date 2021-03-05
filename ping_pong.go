package main

import (
	"fmt"
	"time"
)

func main() {
	pingChan := make(chan string, 1)
	pongChan := make(chan string, 1)

	go ping(pingChan, pongChan)
	go pong(pongChan, pingChan)

	pingChan <- "ping"

	select {}
}

func ping(pingChan <-chan string, pongChan chan<- string) {
	for {
		ball := <-pingChan

		fmt.Println(ball)
		time.Sleep(1 * time.Second)

		pongChan <- "pong"
	}
}

func pong(pongChan <-chan string, pingChan chan<- string) {
	for {
		ball := <-pongChan

		fmt.Println(ball)
		time.Sleep(1 * time.Second)

		pingChan <- "ping"
	}
}
