package main

import (
	"fmt"
	"math/rand"
)

func CalculateRandomValue(results chan int) {
	result := rand.Intn(1000)
	fmt.Println("Random Value: {}", result)
	results <- result
}

func main() {

	results := make(chan int)
	defer close(results)

	go CalculateRandomValue(results)

	result := <-results
	fmt.Println(result)
}
