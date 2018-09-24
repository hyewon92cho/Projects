package main

import (
	"fmt"
	"math"
)

func main() {
	var num int
	print("Enter a number: ")
	fmt.Scanln(&num)

	if num <= 0 {
		println("Invalid input!")
	} else {
		fmt.Println(primeFactor(num))
	}
}

// Validate prime
func primeCheck(num int) bool {
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// Find prime numbers smaller than or equal to given number.
func primeFinder(num int) []int {
	// false means i is a prime number, true means i is not a prime number
	primeBoolArr := make([]bool, num+1)
	primeBoolArr[0] = true
	primeBoolArr[1] = true
	numRoot := int(math.Floor(math.Sqrt(float64(num))))

	for i := 2; i <= numRoot; i++ {
		// if i is a prime number, then its multiples are not prime numbers.
		multiple := 2
		if primeCheck(i) == true {
			for i*multiple <= num {
				primeBoolArr[i*multiple] = true
				multiple++
			}
		}
	}
	var primeNumbers []int
	for index, primeBool := range primeBoolArr {
		if primeBool == false {
			primeNumbers = append(primeNumbers, index)
		}
	}
	return primeNumbers
}

func primeFactor(num int) []int {
	primes := primeFinder(num)

	var factors []int
	for i := 0; i < len(primes); i++ {
		square := primes[i]
		for num%square == 0 {
			square *= primes[i]
			factors = append(factors, primes[i])
		}
	}
	return factors
}
