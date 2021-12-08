package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadInput(txt string) []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Could not open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var nums []int
	for scanner.Scan() {
		cur_line := scanner.Text()
		num, err := strconv.Atoi(cur_line)
		if err != nil {
			log.Fatalf("Could not convert input line: %s to int", cur_line)
		}
		nums = append(nums, num)
	}
	return nums
}

// Find number of times where nums[i+1] > nums[i]
func Part1(nums []int) int {
	num_increase := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			num_increase++
		}
	}
	return num_increase
}

// Find number of times where sum of nums[i] to nums[i+2] is larger than nums[i-3] to nums[i-1]
func Part2(nums []int) int {
	start := 0
	end := 3
	cur_sum := nums[0] + nums[1] + nums[2]
	num_increase := 0
	for end < len(nums) {
		next_sum := cur_sum - nums[start] + nums[end]
		if next_sum > cur_sum {
			num_increase++
		}
		cur_sum = next_sum
		end++
		start++
	}
	return num_increase
}

func main() {
	nums := ReadInput("input.txt")
	fmt.Println(Part1(nums))
	fmt.Println(Part2(nums))
}
