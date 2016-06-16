/*
Problem 2: Find the sum of even fibonacci numbers up to 4 million
*/

package main

import (
	"flag"
	"fmt"
)

func getFibSequence(currNum int, prevNum int, threshold int) []int {
	/*
	   Function getFibSequence
	     Get a fibonacci sequence array up to a specified threshold
	   args:
	     currNum (int) - The current integer
	*/
	nextNum := currNum + prevNum
	output := []int{}
	if nextNum <= threshold {
		output = append(output, nextNum)
		output = append(output, getFibSequence(nextNum, currNum, threshold)...)
	}
	return output
}

func sumEvenIntArray(arr []int) int {
	/*
	  Function sumEvenIntArray
	    Sum up all the even values in an integer array.
	  args:
	    arr ([]int) - An array of integers.
	  returns:
	    sum (int) - The integer sum of all the even elements in the array.
	*/
	sum := 0
	for _, val := range arr {
		if val%2 == 0 {
			sum += val
		}
	}
	return sum
}

func main() {
	var maxNumber int
	flag.IntVar(&maxNumber, "maxNumber", 100, "Max number to get even fibonacci numbers up to.")
	flag.Parse()

	fibSequence := getFibSequence(1, 0, maxNumber)
	evenSum := sumEvenIntArray(fibSequence)
	fmt.Println(fibSequence)
	fmt.Printf("Sum of even fibonacci numbers up to %d:\n%d\n", maxNumber, evenSum)
}
