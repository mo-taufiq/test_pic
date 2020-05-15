package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPalindrom(number int) bool {
	strNumber := strconv.Itoa(number)
	// split string to array
	strArr := strings.Split(strNumber, "")
	// get character lenght of str
	lengthOfStr := len(strArr)
	// get middle of str
	middleOfStr := lengthOfStr / 2
	for i := 0; i < middleOfStr; i++ {
		// from front to middle
		frontToMiddle := strArr[i]
		// from back to middle
		backToMiddle := strArr[(lengthOfStr-1)-i]
		if frontToMiddle != backToMiddle {
			return false
		}
	}
	return true
}

func getTotalPalindromeBasedOnRangeProvided(start int, end int) int {
	count := 0
	for i := start; i <= end; i++ {
		if isPalindrom(i) {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(getTotalPalindromeBasedOnRangeProvided(1, 10))
	fmt.Println(getTotalPalindromeBasedOnRangeProvided(99, 100))
	fmt.Println(getTotalPalindromeBasedOnRangeProvided(21, 31))
}
