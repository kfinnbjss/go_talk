package main

import (
	"fmt"
	"math/big"
	"time"
)

const (
	NWorkers = 5
	NPrimes  = 25
)

func findPrime(input chan int, output chan int) {
	for {
		n := <-input
		time.Sleep(100 * time.Millisecond) //simulate some work
		if big.NewInt(int64(n)).ProbablyPrime(0) {
			output <- n
		}
	}
}

func generateInts() chan int {
	numberChan := make(chan int)
	go func() {
		i := 1
		for {
			numberChan <- i
			i++
		}
	}()
	return numberChan
}

func main() {
	inputChan := generateInts()
	outputChan := make(chan int)
	for i := 0; i < NWorkers; i++ {
		go findPrime(inputChan, outputChan)
	}

	for i := 0; i < NPrimes; i++ {
		fmt.Println(<-outputChan)
	}
}
