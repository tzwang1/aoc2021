package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func ReadInput(file_name string) []string {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatalf("Could not open file: %s", file_name)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var patterns []string
	for scanner.Scan() {
		cur_line := scanner.Text()
		patterns = append(patterns, cur_line)
	}
	return patterns
}

func CalculateErrorScore(pattern string) int {
	stack := []string{}
	for _, c := range pattern {
		if string(c) == "{" || string(c) == "(" || string(c) == "[" || string(c) == "<" {
			stack = append(stack, string(c))
		} else if string(c) == ")" {
			if stack[len(stack)-1] != "(" {
				return 3
			} else {
				stack = stack[:len(stack)-1]
			}
		} else if string(c) == "]" {
			if stack[len(stack)-1] != "[" {
				return 57
			} else {
				stack = stack[:len(stack)-1]
			}
		} else if string(c) == "}" {
			if stack[len(stack)-1] != "{" {
				return 1197
			} else {
				stack = stack[:len(stack)-1]
			}
		} else if string(c) == ">" {
			if stack[len(stack)-1] != "<" {
				return 25137
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return 0
}

func Part1(patterns []string) int {
	total_error_score := 0
	for _, pattern := range patterns {
		total_error_score += CalculateErrorScore(pattern)
	}
	return total_error_score
}

func IsOpenBracket(str string) bool {
	return str == "{" || str == "[" || str == "(" || str == "<"
}

func IsCloseBracket(str string) bool {
	return str == "}" || str == "]" || str == ")" || str == ">"
}

func CalculateCompletionScore(pattern string) int {
	score := 0
	simplified_pattern := []string{}
	for _, c := range pattern {
		if IsOpenBracket(string(c)) {
			simplified_pattern = append(simplified_pattern, string(c))
		} else if IsCloseBracket(string(c)) {
			simplified_pattern = simplified_pattern[:len(simplified_pattern)-1]
		}
	}
	for i := len(simplified_pattern) - 1; i >= 0; i-- {
		cur_char := string(simplified_pattern[i])
		if cur_char == "(" {
			score = score*5 + 1
		} else if cur_char == "[" {
			score = score*5 + 2
		} else if cur_char == "{" {
			score = score*5 + 3
		} else {
			score = score*5 + 4
		}
	}
	return score
}

func Part2(patterns []string) int {
	all_scores := []int{}
	for _, pattern := range patterns {
		if CalculateErrorScore(pattern) == 0 {
			all_scores = append(all_scores, CalculateCompletionScore(pattern))
		}
	}
	sort.Ints(all_scores)
	return all_scores[len(all_scores)/2]
}

func main() {
	patterns := ReadInput("input.txt")
	fmt.Println(Part1(patterns))
	fmt.Println(Part2(patterns))
}
