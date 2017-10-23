package main

import (
	"fmt"
	"math"
)

var testNumber = 100

func main() {

	fmt.Printf("Is Prime ? %d ,  %t ", testNumber, isPrime(testNumber))
}

func isPrime(num int) bool {

	for index := 2; index < int(math.Sqrt(float64(num))); index++ {

		if num%index == 0 {
			return false
		}
	}
	return true
}
