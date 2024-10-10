package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	sum := 0
	primes := []int{}

	for i := 1; i <= 10; i++ {
		if isPrime(i) {
			primes = append(primes, i)
			sum += i
		}
	}

	fmt.Printf("Prime numbers between 1 and 10: %v\n", primes)
	fmt.Printf("Sum of prime numbers: %d\n", sum)
}
