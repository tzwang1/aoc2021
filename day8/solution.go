package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Display struct {
	patterns []string
	output   []string
}

func ReadInputPart1(file_name string) []string {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatalf("Could not open file: %s", file_name)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var nums []string
	for scanner.Scan() {
		cur_line := scanner.Text()
		cur_line_split := strings.Split(cur_line, "|")
		for _, num_str := range strings.Split(cur_line_split[1], " ") {
			if err != nil {
				log.Fatalf("Could not convert string: %s to int", num_str)
			}
			nums = append(nums, num_str)
		}
	}
	return nums
}

func ReadInputPart2(file_name string) []Display {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatalf("Could not open file: %s", file_name)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var display []Display
	for scanner.Scan() {
		cur_line := scanner.Text()
		var cur_display Display
		cur_line_split := strings.Split(cur_line, " | ")
		cur_display.patterns = strings.Split(cur_line_split[0], " ")
		cur_display.output = strings.Split(cur_line_split[1], " ")
		display = append(display, cur_display)
	}
	return display
}

func Part1(nums []string) int {
	num_digits := 0
	for _, num := range nums {
		if len(num) == 2 || len(num) == 4 || len(num) == 7 || len(num) == 3 {
			num_digits++
		}
	}
	return num_digits
}

func StringAContainsB(a, b string) bool {
	for _, str := range b {
		if !strings.Contains(a, string(str)) {
			return false
		}
	}
	return true
}

func NumLetterDiff(a, b string) int {
	var longer string
	var shorter string
	if len(a) > len(b) {
		longer = a
		shorter = b
	} else {
		longer = b
		shorter = a
	}

	num_diff := 0
	for _, str := range longer {
		if !strings.Contains(shorter, string(str)) {
			num_diff++
		}
	}
	return num_diff
}

func DecodeDisplay(display Display) int {
	patterns := display.patterns

	num_to_pattern := make(map[int]string)
	pattern_to_num := make(map[string]int)

	// First iteration numbers
	for _, pattern := range patterns {
		if len(pattern) == 2 {
			num_to_pattern[1] = pattern
			pattern_to_num[pattern] = 1
		} else if len(pattern) == 3 {
			num_to_pattern[7] = pattern
			pattern_to_num[pattern] = 7
		} else if len(pattern) == 4 {
			num_to_pattern[4] = pattern
			pattern_to_num[pattern] = 4
		} else if len(pattern) == 7 {
			num_to_pattern[8] = pattern
			pattern_to_num[pattern] = 8
		}
	}

	// Second iteration numbers
	for _, pattern := range patterns {
		if len(pattern) == 5 && StringAContainsB(pattern, num_to_pattern[1]) {
			num_to_pattern[3] = pattern
			pattern_to_num[pattern] = 3
		} else if len(pattern) == 6 && !StringAContainsB(pattern, num_to_pattern[1]) {
			num_to_pattern[6] = pattern
			pattern_to_num[pattern] = 6
		}
	}

	// Third iteration number
	for _, pattern := range patterns {
		if len(pattern) == 6 && NumLetterDiff(num_to_pattern[3], pattern) == 1 {
			num_to_pattern[9] = pattern
			pattern_to_num[pattern] = 9
		}
	}

	// Fourth iteration numbers
	for _, pattern := range patterns {
		if len(pattern) == 5 && NumLetterDiff(pattern, num_to_pattern[6]) == 1 {
			num_to_pattern[5] = pattern
			pattern_to_num[pattern] = 5
		} else if len(pattern) == 6 && pattern != num_to_pattern[9] && pattern != num_to_pattern[6] {
			num_to_pattern[0] = pattern
			pattern_to_num[pattern] = 0
		}
	}

	// Fifth iteration number
	for _, pattern := range patterns {
		if len(pattern) == 5 && pattern != num_to_pattern[5] && pattern != num_to_pattern[3] {
			num_to_pattern[2] = pattern
			pattern_to_num[pattern] = 2
		}
	}
	// fmt.Println(num_to_pattern)
	// fmt.Println(pattern_to_num)

	outputs := display.output
	num := 0
	for _, output := range outputs {
		for k, v := range pattern_to_num {
			if NumLetterDiff(output, k) == 0 {
				num = (num * 10) + v
			}
		}
	}
	return num
}

func Part2(displays []Display) int {
	nums := 0
	for _, display := range displays {
		nums += DecodeDisplay(display)
	}
	return nums
}

func main() {
	nums := ReadInputPart1("input.txt")
	fmt.Println(Part1(nums))

	display := ReadInputPart2("input.txt")
	// fmt.Println(display)
	fmt.Println(Part2(display))
}

/*
0 -> pattern with len 6 that is not 6 or 9 (4)
1 -> pattern with len 2 (1)
2 -> pattern with len 5 that is not 5 and not 3 (5)
3 -> pattern with len 5 with both letters from 1 (2)
4 -> pattern with len 4 (1)
5 -> pattern with len 5 with 1 letter different from 6 (4)
6 -> pattern with len 6 that does not have one letter from 1 (2)
7 -> pattern with len 3 (1)
8 -> pattern with len 7 (1)
9 -> pattern with len 6 with 1 letter different from 3 (3)

*/
