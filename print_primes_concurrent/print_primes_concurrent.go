package main

import (
	"fmt"
	"math/big"
	"time"
)

func printIfPrime(n int) {
	time.Sleep(100 * time.Millisecond) //simulate some work
	if big.NewInt(int64(n)).ProbablyPrime(0) {
		fmt.Println(n)
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		go printIfPrime(i)
	}
	time.Sleep(200 * time.Millisecond) //simulate some work
}
