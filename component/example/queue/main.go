package main

import (
	"log"
	"time"
)

func startCunsumer(queue chan int, name string) {
	for {
		time.Sleep(time.Second)
		log.Println(name, <-queue)
	}
}

func main() {
	n := 10000
	queue := make(chan int, n)
	for i := 1; i <= n; i++ {
		queue <- i
	}

	go startCunsumer(queue, "C1")
	go startCunsumer(queue, "C2")

	time.Sleep(10 * time.Second)
}
