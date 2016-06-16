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
)

func isPrime(num int) bool {
	/*
	   Function isPrime
	     Return true or false if the number is prime
	   args:
	     num (int) - The number to test if it is prime
	   returns:
	     (boolean)
	*/
	output := true
	if num%2 == 0 {
		output = false
	}
	for i := 3; i < num/2; i += 2 {
		if num%i == 0 {
			output = false
			break
		}
	}
	return output
}

func getPrimeNumberSequence(maxNum int) []int {
	/*
	   Function getPrimeNumberSequence
	     Get an array of prime numbers up to a specified value.
	   args:
	     maxNum (int) - The maximum int to find primes up to.
	   returns:
	     output ([]int) - An array of prime numbers
	*/
	output := []int{}
	for i := 1; i <= maxNum; i += 2 {
		if isPrime(i) {
			output = append(output, i)
		}
	}
	return output
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
	for _, val := range getPrimeNumberSequence(num) {
		if num%val == 0 && val > largestPrime {
			largestPrime = val
		}
	}
	return largestPrime
}

func main() {
	var num int
	flag.IntVar(&num, "n", 81, "Number to get largest prime factor for.")
	flag.Parse()

	//fmt.Println(getPrimeNumberSequence(num))
	//fmt.Println(isPrime(9))
	largestPrimeFactor := getLargestPrimeFactor(num)
	fmt.Printf("The largest prime factor for %d:\n%d\n", num, largestPrimeFactor)
}
