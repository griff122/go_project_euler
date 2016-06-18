/*
Problem 5: Find the smallest positive number that is evenly divisible by all
  of the number from 1 to 20
*/

package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var maxNum int
	flag.IntVar(&maxNum, "n", 10, "Max number to get multiples of 3 and 5 for.")
	flag.Parse()
	currNum := maxNum
	multipleFound := false
	fmt.Printf("Finding smallest number with even multiples up to: %d\n", maxNum)
	s1 := time.Now()
	for multipleFound == false {
		multipleFound = true
		for i := 1; i <= maxNum; i++ {
			if currNum%i != 0 {
				multipleFound = false
				break
			}
		}
		currNum++
	}
	fmt.Println("Finished...")
	fmt.Printf("Smallest Multiple: %d\n", currNum-1)
	fmt.Printf("Total Time: %s\n", time.Since(s1))
}
