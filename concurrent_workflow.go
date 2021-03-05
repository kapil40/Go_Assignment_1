package main

import "fmt"

func square(c chan int) {
	fmt.Println("[square] reading")
	num := <-c
	c <- num * num
}

func cube(c chan int) {
	fmt.Println("[cube] reading")
	num := <-c
	c <- num * num * num
}

func main() {
	fmt.Println("[main] main started")

	squareChan := make(chan int)
	cubeChan := make(chan int)

	go square(squareChan)
	go cube(cubeChan)

	n := 5
	fmt.Println("[main] sent n to squareChan")

	squareChan <- n

	fmt.Println("[main] resuming")
	fmt.Println("[main] sent n to cubeChan")

	cubeChan <- n

	fmt.Println("[main] resuming")
	fmt.Println("[main] reading from channels")

	squareVal, cubeVal := <-squareChan, <-cubeChan
	sum := squareVal + cubeVal

	fmt.Println("[main] sum of square and cube is: ", sum)
	fmt.Println("[main] main() stopped")
}
