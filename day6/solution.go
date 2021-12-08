package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadInput(file_name string) []int {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatalf("Could not open file: %s", file_name)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var nums []int
	for scanner.Scan() {
		cur_line := scanner.Text()
		cur_line_split := strings.Split(cur_line, ",")
		for _, num_str := range cur_line_split {
			num, err := strconv.Atoi(num_str)
			if err != nil {
				log.Fatalf("Could not convert string: %s to int", num_str)
			}
			nums = append(nums, num)
		}
	}
	return nums
}

func Part1(nums []int, num_days int) int {
	days := 0

	for days < num_days {
		var new_nums []int

		for _, num := range nums {
			if num == 0 {
				new_nums = append(new_nums, 8)
				new_nums = append(new_nums, 6)
			} else {
				new_nums = append(new_nums, num-1)
			}
		}
		days++
		nums = new_nums
	}
	return len(nums)
}

func Part2Attempt1(nums []int, num_days int) int {
	num_fish := len(nums)
	var cur_fish []int
	var new_fish []int

	for _, num := range nums {
		cur_fish = append(cur_fish, num_days-num)
		cur_num := num_days - num
		for cur_num >= 0 {
			cur_num -= 7
			if cur_num >= 0 {
				cur_fish = append(cur_fish, cur_num)
			}
		}
	}
	num_fish += len(cur_fish)

	for len(cur_fish) != 0 {
		new_fish = []int{}
		for _, fish := range cur_fish {
			new_born := true
			for fish >= 0 {
				fish -= 7
				if new_born {
					fish -= 2
					new_born = false
				}
				if fish >= 0 {
					new_fish = append(new_fish, fish)
				}
			}
		}
		cur_fish = new_fish
		num_fish += len(cur_fish)
	}
	return num_fish
}

func Part2Attempt2(nums []int, num_days int) int {
	num_fish := len(nums)
	var cur_fish []int

	for _, num := range nums {
		cur_fish = append(cur_fish, num_days-num)
		cur_num := num_days - num
		for cur_num >= 0 {
			cur_num -= 7
			if cur_num >= 0 {
				cur_fish = append(cur_fish, cur_num)
			}
		}
	}
	seen := make(map[int]int)
	for _, fish := range cur_fish {
		num_fish += Part2Attempt2Helper(fish, seen)
	}
	return num_fish
}

func Part2Attempt2Helper(num_days_left int, seen map[int]int) int {
	// fmt.Println(seen)
	if num_days_left <= 0 {
		return 1
	}
	if val, ok := seen[num_days_left]; ok {
		return val
	}
	num_fish := 0
	new_born := true
	orig_num_days_left := num_days_left
	for num_days_left >= 0 {
		num_days_left -= 7
		if new_born {
			num_days_left -= 2
			new_born = false
		}
		num_fish += Part2Attempt2Helper(num_days_left, seen)
	}
	seen[orig_num_days_left] = num_fish
	return num_fish
}

func UpdateMap(a map[int]int) {
	a[1] = 5
}

func main() {
	nums := ReadInput("small_input.txt")
	fmt.Println("Part1 solution: ", Part1(nums, 80))
	// fmt.Println("Part2 solution: ", Part2Attempt1(nums, 80))
	fmt.Println("Part2 attempt 2 solution: ", Part2Attempt2(nums, 255))

	// a := make(map[int]int)
	// a[2] = 3
	// UpdateMap(a)
	// fmt.Println(a)
}

/*

helper(5)
helper(4) helper(4)
helper(3) helper(3)
helper(2) helper(2)
helper(1) helper(1)


*/
