package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

func findMaxMinNum(nums []int) (int, int) {
	min := math.MaxInt32
	max := math.MinInt32

	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return max, min
}

func CalculateFuelCostPart1(nums []int, horizontal_pos int) int {
	fuel_cost := 0
	for _, num := range nums {
		fuel_cost += int(math.Abs(float64(horizontal_pos - num)))
	}
	return fuel_cost
}

func Part1(nums []int) int {
	max, min := findMaxMinNum(nums)
	min_fuel_cost := math.MaxInt32

	for min < max {
		mid := (min + max) / 2
		fuel_cost_left := CalculateFuelCostPart1(nums, (min+mid)/2)
		fuel_cost_right := CalculateFuelCostPart1(nums, (mid+max)/2)
		if fuel_cost_left < fuel_cost_right {
			min_fuel_cost = int(math.Min(float64(fuel_cost_left), float64(min_fuel_cost)))
			max = mid
		} else {
			min_fuel_cost = int(math.Min(float64(fuel_cost_right), float64(min_fuel_cost)))
			min = mid + 1
		}
	}
	return min_fuel_cost
}

func CalculateFuelCostPart2(nums []int, horizontal_pos int) int {
	fuel_cost := 0
	for _, num := range nums {
		diff := int(math.Abs(float64(horizontal_pos - num)))
		fuel_cost += (diff * (diff + 1)) / 2
	}
	return fuel_cost
}

func Part2(nums []int) int {
	max, min := findMaxMinNum(nums)
	min_fuel_cost := math.MaxInt32

	for min < max {
		mid := (min + max) / 2
		fuel_cost_left := CalculateFuelCostPart2(nums, (min+mid)/2)
		fuel_cost_right := CalculateFuelCostPart2(nums, (mid+max)/2)
		if fuel_cost_left < fuel_cost_right {
			min_fuel_cost = int(math.Min(float64(fuel_cost_left), float64(min_fuel_cost)))
			max = mid
		} else {
			min_fuel_cost = int(math.Min(float64(fuel_cost_right), float64(min_fuel_cost)))
			min = mid + 1
		}
	}
	return min_fuel_cost
}

func main() {
	nums := ReadInput("input.txt")
	fmt.Println(Part1(nums))
	fmt.Println(Part2(nums))
}
