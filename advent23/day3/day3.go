package main

import (
	"advent23/util/input"
	"fmt"
	"strconv"
)

func main() {
	file, lines := input.LineByLine("day3/input.txt")
	defer file.Close()

	fmt.Println("part_1", part1(lines))
	fmt.Println("part_2", part2(lines))
}

func part1(lines []string) int {
	sum := 0
	for i, line := range lines {
		numSeq := extractSeqOfNums(line)
		for seqIndex, num := range numSeq {
			seqLength := len(strconv.Itoa(num))
			above := ""
			below := ""
			left := seqIndex - 1
			right := seqIndex + seqLength + 1
			if left < 0 {
				left = 0
			}
			if right > len(line) {
				right = len(line)
			}
			if i > 0 {
				above = lines[i-1][left:right]
			}
			if i < len(lines)-1 {
				below = lines[i+1][left:right]
			}
			prev := line[left : left+1]
			next := line[right-1 : right]

			if containsSymbol(above) || containsSymbol(below) || containsSymbol(prev) || containsSymbol(next) {
				fmt.Printf("seq: %d; above: %s; below: %s; prev: %s; next: %s\n", num, above, below, prev, next)
				sum += num
			}
		}
	}
	return sum
}

func extractSeqOfNums(line string) map[int]int {
	seq := map[int]int{}
	s := ""
	for i, r := range line {
		if isDigit(r) {
			s += string(r)
		} else if len(s) > 0 {
			add(i-len(s), &s, seq)
			n, _ := strconv.Atoi(s)
			seq[i-len(s)] = n
			s = ""
		}
	}
	if len(s) > 0 {
		add(len(line)-len(s), &s, seq)
	}
	return seq
}

func add(key int, value *string, seq map[int]int) {
	n, _ := strconv.Atoi(*value)
	seq[key] = n
	value = nil
}

func part2(lines []string) int {
	sum := 0
	gearCandidates := map[int][]int{}
	for i, line := range lines {
		numSeq := extractSeqOfNums(line)
		for seqIndex, num := range numSeq {
			seqLength := len(strconv.Itoa(num))
			above := ""
			below := ""
			left := seqIndex - 1
			right := seqIndex + seqLength + 1
			if left < 0 {
				left = 0
			}
			lineLen := len(line)
			if right > lineLen {
				right = lineLen
			}
			if i > 0 {
				above = lines[i-1][left:right]
			}
			if i < len(lines)-1 {
				below = lines[i+1][left:right]
			}
			prev := line[left : left+1]
			next := line[right-1 : right]
			candidates := map[int]string{
				(i-1)*lineLen + left:    above,
				(i+1)*lineLen + left:    below,
				(i)*lineLen + left:      prev,
				(i)*lineLen + right - 1: next,
			}
			for index, str := range candidates {
				if relIndex, ok := containsGear(str); ok {
					index = index + *relIndex
					_, ok := gearCandidates[index]
					if ok {
						gearCandidates[index] = append(gearCandidates[index], num)
					} else {
						gearCandidates[index] = []int{num}
					}
				}
			}
		}
	}
	for _, parts := range gearCandidates {
		if len(parts) != 2 {
			continue
		}
		sum += parts[0] * parts[1]
	}
	return sum
}

func containsGear(str string) (*int, bool) {
	if len(str) == 0 {
		return nil, false
	}
	for i, r := range str {
		if r == '*' {
			return &i, true
		}
	}
	return nil, false
}

func isDigit(r rune) bool {
	return r >= 48 && r <= 57
}

func containsSymbol(str string) bool {
	if len(str) == 0 {
		return false
	}
	for _, r := range str {
		if !isDigit(r) && r != '.' {
			return true
		}
	}
	return false
}
