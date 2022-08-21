/*
This is exercise "Get the Middle Character" from Codewars (7 kyu level)

Instructions:
You are going to be given a word. Your job is to return the middle character of the word. If the word's length is odd, return the middle character. If the word's length is even, return the middle 2 characters.

#Examples:

Kata.getMiddle("test") should return "es"

Kata.getMiddle("testing") should return "t"

Kata.getMiddle("middle") should return "dd"

Kata.getMiddle("A") should return "A"
#Input

A word (string) of length 0 < str < 1000 (In javascript you may get slightly more than 1000 in some test cases due to an error in the test cases). You do not need to test for this. This is only here to tell you that you do not need to worry about your solution timing out.

#Output

The middle character(s) of the word represented as a string.
*/

// Solution 1
package kata


func GetMiddle(s string) string {
  strLen := len(s)
  if (strLen % 2 != 0) { return string(s[strLen / 2]) }
  return s[(strLen / 2) - 1 : strLen / 2 + 1]
}

// Solution 2
package kata

import "strings"

func GetMiddle(s string) string {
  strArr := strings.Split(s, "")
  strLen := len(strArr)

  if (strLen % 2 != 0) {
    return strArr[int(strLen / 2)]
  } else {
    return strArr[(strLen / 2) - 1] + strArr[strLen / 2]
  }
}
