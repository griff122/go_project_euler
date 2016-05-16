/*
Problem 1: Find the sum of all multiples of 3 and 5 up to 1000.
*/

package main

import (
	"flag"
	"fmt"
)

func getMultiplesOf(multiples []int, maxNum int) []int {
	/*
	  Function getMultiplesOf
	    Get all the multiples up to a specified number.
	  args:
	    multiples ([]int) - A list of integers to find multiples for.
	    maxNum (int) - An integer to iterate up to.
	  returns:
	    output ([]int) - A list of all the multiples up to max_num
	*/
	output := []int{}
	for i := 0; i <= maxNum; i++ {
		for _, mul := range multiples {
			if i%mul == 0 {
				output = append(output[:], i)
				break
			}
		}
	}
	return output
}

func sumIntArray(arr []int) int {
	/*
	  Function sumIntArray
	    Sum up all the values in an integer array.
	  args:
	    arr ([]int) - An array of integers.
	  returns:
	    sum (int) - The integer sum of all the elements in the array.
	*/
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func main() {

	var maxNumber int
	flag.IntVar(&maxNumber, "maxNumber", 100, "Max number to get multiples of 3 and 5 for.")
	flag.Parse()
	multiples := []int{3, 5}
	multipleArr := getMultiplesOf(multiples, maxNumber)
	multipleSum := sumIntArray(multipleArr)
	fmt.Printf("Total sum of multiples up to %d: %d\n", maxNumber, multipleSum)
}
