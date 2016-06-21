/*
Problem 6: Find the difference between the sum of squares for the first
  100 natural numbers and the square of the sum
*/

package main

import (
	"flag"
	"fmt"
	"math"
	"time"
)

func main() {
	var maxNumber int
	flag.IntVar(&maxNumber, "maxNumber", 100, "Max number to get multiples of 3 and 5 for.")
	flag.Parse()
	fmt.Printf("Finding the sum of squares for the first %d natural numbers\n", maxNumber)

	// Start a timer
	s1 := time.Now()

	sumOfSquares := 0.
	squareOfSums := 0.

	// Iterate over the range of natural numbers
	for i := 1; i <= maxNumber; i++ {
		sumOfSquares += math.Pow(float64(i), 2.)
		squareOfSums += float64(i)
	}
	// Square the final sum
	squareOfSums = math.Pow(squareOfSums, 2)

	fmt.Printf("Finished. Time taken: %s\n", time.Since(s1))
	fmt.Printf("sumOfSquares: %f\n", sumOfSquares)
	fmt.Printf("squareOfSums: %f\n", squareOfSums)
	fmt.Printf("Difference: %f\n", squareOfSums-sumOfSquares)

}
