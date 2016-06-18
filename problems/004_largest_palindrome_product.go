/*
Problem 4: Find the largest palindrome made from the product of two 3 digit numbers
*/
package main

import (
	"fmt"
	"strconv"
	"time"
)

func isPalindrome(val int) bool {
	/*
	   Function isPalindrome
	     Determines if an integer is the same number if the characters are
	     reversed.
	   args:
	     val (int) - The integer to test for palindromeness
	   returns:
	     (bool) - True/False if val is a palindrome or not.
	*/
	valStr := strconv.Itoa(val)
	reversedValStr := reverseString(valStr)
	return valStr == reversedValStr
}

func reverseString(val string) string {
	/*
	   Function reverseString
	     Returns the reverse of a string.
	   args:
	     val (string) - The string to reverse
	   returns:
	     (string) - The reversed string
	*/
	reversed := make([]byte, len(val))
	for i := 0; i < len(val); i++ {
		reversed[len(val)-i-1] = val[i]
	}
	return string(reversed)
}

func main() {
	fmt.Println("Starting program to get the largest palindrome product of two 3 digit numbers.")
	s1 := time.Now()
	minNum, maxNum := 100, 999
	largestPalindrome := 0
	currProd := 0
	for x := maxNum; x >= minNum; x-- {
		for y := maxNum; y >= minNum; y-- {
			currProd = x * y
			if isPalindrome(currProd) && currProd > largestPalindrome {
				largestPalindrome = currProd
			}
		}
	}
	e1 := time.Since(s1)
	fmt.Println("Finished...")
	fmt.Printf("Largest palindrome: %d\n", largestPalindrome)
	fmt.Printf(" - Took %s\n", e1)
}
