package main

import (
	"fmt"
	"math"
)

// init huge array for storing abundant numbers.

func main() {
	var abundantArray [7000]int
	maxNumber := 28123
	count := 0
	abundantCount := 0
	total := 0

	// calculate abundant numbers
	for i := 1; i < maxNumber; i++ {
		//fmt.Printf("\nIteration: %d \n", i)

		isAbundantNumber, numberValue := isAbundant(i)
		if isAbundantNumber {
			abundantArray[count] = numberValue
			//fmt.Printf("Abundant Number %d: %d\n", count, numberValue)
			count++
		} else {
			//fmt.Printf("Number[%d] is not abundant.\n", i)
		}

		if count > 0 {

			tryingTempNumber := i

			//var mokoko int = tryingTempNumber + 5
			//	_ //= tryingTempNumber
			//fmt.Printf("\nChecking number: %d \n ", tryingTempNumber)
			for tryingTempNumber > 0 {

				if abundantCount < count {
					//fmt.Printf("\n\n\n Abundant Count[%d] - ArrayCount[%d] \n\n\n ", abundantCount, count)

					tryingAbundantNumber := abundantArray[abundantCount]

					if tryingAbundantNumber < tryingTempNumber {
						tryingTempNumber = i - tryingAbundantNumber
						//fmt.Printf("Trying number: %d \n", tryingTempNumber)
						//if tryingTempNumber > 0 {
						isAbundant, _ := isAbundant(tryingTempNumber)
						if isAbundant {
							//fmt.Printf("#####Num is Abundantable: %d #########\n", nnumber)
							//fmt.Printf("----Sum of %d + %d \n", tryingTempNumber, tryingAbundantNumber)
							abundantCount = 0
							break
						} else {
							//fmt.Printf("Trying another number...\n")
							abundantCount++
							continue
						}
					} else {
						//fmt.Printf("<<<<<Num is NOT Abundantable: %d >>>>>\n ", nnumber)
						total += i
						abundantCount = 0
						break
					}
				} else {
					total += i
					abundantCount = 0
					break
				}

			}
		} else {
			total += i
		}

	}
	fmt.Printf("Calculation Finished RESULT\n %d ", total)

}

func isPrime(num int) bool {

	for index := 2; index < int(math.Sqrt(float64(num))); index++ {

		if num%index == 0 {
			return false
		}
	}
	return true
}

func isAbundant(number int) (bool, int) {

	if !isPrime(number) {
		if sumOfDividers(number) > number {
			return true, number
		}
	}
	return false, number
}

func greatestDivider(num int) int {

	for index := 2; index < num; index++ {
		if num%index == 0 {
			return num / index
		}
	}
	return -1
}

func sumOfDividers(num int) int {

	var loopLimit = greatestDivider(num)

	var sum int = 1

	for index := 2; index <= loopLimit; index++ {

		if num%index == 0 {
			sum += index
		}
	}

	return sum
}
