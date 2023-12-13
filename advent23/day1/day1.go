package main

import (
	"advent23/util/assert"
	"advent23/util/input"
	"fmt"
	"strconv"
)

func main() {
	file, lines := input.LineByLine("day1/input.txt")
	defer file.Close()

	fmt.Println("part_1", part1(lines))
	fmt.Println("part_2", part2(lines))
}

func part1(lines []string) int {
	result := 0
	extractFn := func(r rune, _ string) (string, bool) {
		return string(r), isDigit(r)
	}
	for _, str := range lines {
		result += extractDigits(str, extractFn)
	}
	return result
}

func part2(lines []string) int {
	result := 0
	words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	extractFn := func(r rune, token string) (string, bool) {
		if isDigit(r) {
			return string(r), true
		}
		match := false
		index := 0
		for i, word := range words {
			if len(word) > len(token) {
				continue
			}
			if token[len(token)-len(word):] == word {
				match = true
				index = i
				break
			}
		}
		return strconv.Itoa(index + 1), match
	}
	for _, str := range lines {
		result += extractDigits(str, extractFn)
	}
	return result
}

func isDigit(r rune) bool {
	return r >= 48 && r <= 57
}

type extractDigitFn func(rune, string) (string, bool)

func extractDigits(str string, fn extractDigitFn) int {
	var digits []string
	token := ""
	for _, r := range str {
		token += string(r)
		digit, ok := fn(r, token)
		if !ok {
			continue
		}
		digits = append(digits, digit)
	}
	if len(digits) == 0 {
		return 0
	}
	digit, err := strconv.Atoi(digits[0] + digits[len(digits)-1])
	assert.Empty(err)
	return digit
}
