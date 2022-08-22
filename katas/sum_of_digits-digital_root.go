/*
This is exercise "Sum of Digits / Digital Root" from Codewars (6 kyu level)

Instructions:
- Given n, take the sum of the digits of n. If that value has more than one digit, continue reducing in this way until a single-digit number is produced. The input will be a non-negative integer.

Examples
16  -->  1 + 6 = 7
942  -->  9 + 4 + 2 = 15  -->  1 + 5 = 6
132189  -->  1 + 3 + 2 + 1 + 8 + 9 = 24  -->  2 + 4 = 6
493193  -->  4 + 9 + 3 + 1 + 9 + 3 = 29  -->  2 + 9 = 11
*/
package katas

import (
	"strconv"
)

func DigitalRoot(n int) int {

	getSumOfDigits := func(number int) int {
		strNumber := strconv.Itoa(n)
		acc := 0
		for _, value := range strNumber {
			number, _ := strconv.Atoi(string(value))
			acc += number
		}
		return acc
	}

	for n >= 10 {
		n = getSumOfDigits(n)
	}
	return n
}
