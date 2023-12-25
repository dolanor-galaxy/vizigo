package main

import (
	"fmt"
	"log"
	"strings"
	"unicode"
)

func PadStringToCenter(s string, width int) string {
	if len(s) >= width {
		return s
	}
	leftPadding := (width - len(s)) / 2
	rightPadding := width - len(s) - leftPadding
	return strings.Repeat(" ", leftPadding) + s + strings.Repeat(" ", rightPadding)
}

func SplitStringAt(s string, i int) (string, string, string, error) {
	if i < 0 || i > len(s) {
		return "", "", "", fmt.Errorf("can't split string at %d", i)
	}
	return s[:i], s[i : i+1], s[i+1:], nil
}

func UnderlineChar(s string, i int) string {

	if i < 0 {
		return s
	}

	start, underline, end, error := SplitStringAt(s, i)
	if error != nil {
		log.Fatal(error)
	}

	return fmt.Sprintf("%s\033[4m%s\033[0m%s", start, underline, end)
	
}

func ColumnToLetters(n int) string {
	var result string
	for n > 0 {
		remainder := (n - 1) % 26
		result = string('A'+remainder) + result
		n = (n - 1) / 26
	}
	return result
}

func LettersToColumn(s string) int {
	var result int
	for _, c := range s {
		result *= 26
		result += int(c) - 'A' + 1
	}
	return result
}

func SplitAlphaNumeric(s string) (alphaPart string, numericPart string) {
	splitIndex := -1
	for i, char := range s {
		if unicode.IsDigit(char) {
			splitIndex = i
			break
		}
	}
	if splitIndex != -1 {
		return s[:splitIndex], s[splitIndex:]
	}
	return s, ""
}
