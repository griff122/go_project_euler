/*
Problem 3: Find largest prime factor
TODO: TOO SLOW. Update the script to work backwards from num/2
      Once a number is ID'd as a prime, and it is a factor,
      then that is our answer
*/

package main

import (
	"flag"
	"fmt"
	"math"
	"time"
)

func sieveOfEratosthenes(n int) []int {
	/*
		Function sieveOfEratosthenes
			Get the list of primes through sieve of eratosthenes algorithm.
		args:
			n (int) - The max number to find primes for
		returns:
			([]int) - A list of primes
	*/
	fmt.Println("Launching sieve function...")
	s1 := time.Now()
	fmt.Println("Starting initial array creation.")
	totalNum := int(math.Ceil(math.Sqrt(float64(n)))) + 1
	vals := make([]int, totalNum/2)
	primes := []int{2}
	counter := 0
	for i := 3; i <= totalNum; i += 2 {
		vals[counter] = i
		counter++
	}
	fmt.Printf(" -- Array initialization took: %s\n", time.Since(s1))
	// now need to parse through the values and remove all elements that are
	fmt.Println("Starting sieve.")
	s2 := time.Now()
	for len(vals) > 1 {
		currPrime := vals[0]
		primes = append(primes, currPrime)
		vals = append(vals[:0], vals[1:]...)
		for i := len(vals) - 1; i >= 0; i-- {
			val := vals[i]
			if val%currPrime == 0 {
				vals = append(vals[:i], vals[i+1:]...)
			}
		}
	}
	fmt.Printf(" -- Sieve took: %s\n", time.Since(s2))
	fmt.Printf("Total time for sieve function: %s\n", time.Since(s1))
	return primes
}

func getLargestPrimeFactor(num int) int {
	/*
	   Function getLargestPrimeFactor
	     Return the largest prime factor for a number
	   args:
	     num (int) - Number to find the largest prime factor of
	   returns:
	     (int) - The largest prime factor
	*/
	largestPrime := -1

	primes := sieveOfEratosthenes(num)

	for i := len(primes) - 1; i >= 0; i-- {
		if num%primes[i] == 0 {
			largestPrime = primes[i]
			break
		}
	}
	return largestPrime
}

func main() {
	var num int
	flag.IntVar(&num, "n", 81, "Number to get largest prime factor for.")
	flag.Parse()

	largestPrimeFactor := getLargestPrimeFactor(num)

	fmt.Printf("The largest prime factor for %d:\n%d\n", num, largestPrimeFactor)
}
