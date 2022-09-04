/*
This is exercise "Counting Duplicates" from Codewars (6 kyu level)

Instructions:
Count the number of Duplicates
Write a function that will return the count of distinct case-insensitive alphabetic characters and numeric digits that occur more than once in the input string. The input string can be assumed to contain only alphabets (both uppercase and lowercase) and numeric digits.

Example
"abcde" -> 0 # no characters repeats more than once
"aabbcde" -> 2 # 'a' and 'b'
"aabBcde" -> 2 # 'a' occurs twice and 'b' twice (`b` and `B`)
"indivisibility" -> 1 # 'i' occurs six times
"Indivisibilities" -> 2 # 'i' occurs seven times and 's' occurs twice
"aA11" -> 2 # 'a' and '1'
"ABBA" -> 2 # 'A' and 'B' each occur twice
*/
package katas

import (
	"strings"
)

func Duplicate_count(s1 string) int {

	strArr := strings.Split(s1, "")
	strMap := make(map[string]int)

	for _, value := range strArr {
		strMap[strings.ToLower(value)] = 0
	}

	for i := 0; i < len(strArr); i++ {
		counter := 0
		for j := 0; j < len(strArr); j++ {
			if strings.EqualFold(strArr[i], strArr[j]) {
				counter++
				strMap[strings.ToLower(strArr[i])] = counter
			}
		}
	}
	result := 0
	for _, value := range strMap {
		if value >= 2 {
			result++
		}
	}

	return result
}
