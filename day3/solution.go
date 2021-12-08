package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadInput(file_name string) []string {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal("Could not open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var binary_strs []string
	for scanner.Scan() {
		cur_line := scanner.Text()
		binary_strs = append(binary_strs, cur_line)
	}
	return binary_strs
}

func Part1(binary_strs []string) int64 {
	cur_one_count := 0
	cur_zero_count := 0
	var gamma_rate string
	var epsilon_rate string

	for i := 0; i < len(binary_strs[0]); i++ {
		for j := 0; j < len(binary_strs); j++ {
			if string(binary_strs[j][i]) == "0" {
				cur_zero_count += 1
			} else {
				cur_one_count += 1
			}
		}
		if cur_zero_count > cur_one_count {
			gamma_rate += "0"
			epsilon_rate += "1"
		} else {
			gamma_rate += "1"
			epsilon_rate += "0"
		}
		cur_zero_count = 0
		cur_one_count = 0
	}
	gamma_rate_dec, err := strconv.ParseInt(gamma_rate, 2, 64)
	if err != nil {
		log.Fatalf("Could not convert string: %s to decimal", gamma_rate)
	}
	epsilon_rate_dc, err := strconv.ParseInt(epsilon_rate, 2, 64)
	if err != nil {
		log.Fatalf("Could not convert string: %s to decimal", epsilon_rate)
	}
	return gamma_rate_dec * epsilon_rate_dc
}

func Part2(binary_strs []string) int64 {
	var oxygen_rating string
	cur_one_count := 0
	cur_zero_count := 0
	oxygen_rating_binary_strs := binary_strs
	for j := 0; len(oxygen_rating_binary_strs) > 1 && j < len(oxygen_rating_binary_strs[0]); j++ {
		for i := 0; i < len(oxygen_rating_binary_strs); i++ {
			if string(oxygen_rating_binary_strs[i][j]) == "1" {
				cur_one_count++
			} else {
				cur_zero_count++
			}
		}
		var num_to_remove string
		if cur_one_count >= cur_zero_count {
			num_to_remove = "0"
		} else {
			num_to_remove = "1"
		}
		oxygen_rating_binary_strs = removeFromList(oxygen_rating_binary_strs, j, num_to_remove)
		cur_one_count = 0
		cur_zero_count = 0
	}
	oxygen_rating = oxygen_rating_binary_strs[0]
	co2_rating_binary_strs := binary_strs
	var co2_rating string
	for j := 0; len(co2_rating_binary_strs) > 1 && j < len(co2_rating_binary_strs[0]); j++ {
		for i := 0; i < len(co2_rating_binary_strs); i++ {
			if string(co2_rating_binary_strs[i][j]) == "1" {
				cur_one_count++
			} else {
				cur_zero_count++
			}
		}
		var num_to_remove string
		if cur_zero_count <= cur_one_count {
			num_to_remove = "1"
		} else {
			num_to_remove = "0"
		}
		co2_rating_binary_strs = removeFromList(co2_rating_binary_strs, j, num_to_remove)
		cur_one_count = 0
		cur_zero_count = 0
	}
	co2_rating = co2_rating_binary_strs[0]

	co2_rating_dec, err := strconv.ParseInt(co2_rating, 2, 64)
	if err != nil {
		log.Fatalf("Could not convert string: %s to decimal", co2_rating)
	}
	oxygen_rating_dec, err := strconv.ParseInt(oxygen_rating, 2, 64)
	if err != nil {
		log.Fatalf("Could not convert string: %s to decimal", oxygen_rating)
	}
	return co2_rating_dec * oxygen_rating_dec
}

func removeFromList(binary_strs []string, idx int, num_string string) []string {
	var modified_binary_strs []string
	for _, binary_str := range binary_strs {
		if string(binary_str[idx]) != num_string {
			modified_binary_strs = append(modified_binary_strs, binary_str)
		}
	}
	return modified_binary_strs
}

func main() {
	binary_strs := ReadInput("input.txt")
	fmt.Println(Part1(binary_strs))
	fmt.Println(Part2(binary_strs))
}
