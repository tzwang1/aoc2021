package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type Point struct {
	row int
	col int
}

func ReadInput(file_name string) [][]int {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatalf("Could not open file: %s", file_name)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var nums [][]int
	for scanner.Scan() {
		cur_line := scanner.Text()
		var cur_row []int
		for _, num_str := range cur_line {
			num, err := strconv.Atoi(string(num_str))
			if err != nil {
				log.Fatalf("Could not convert string: %s to int", string(num_str))
			}
			cur_row = append(cur_row, num)
		}
		nums = append(nums, cur_row)
	}
	return nums
}

func IsLowPoint(row int, col int, nums [][]int) bool {
	if row-1 >= 0 && nums[row-1][col] <= nums[row][col] {
		return false
	}
	if row+1 < len(nums) && nums[row+1][col] <= nums[row][col] {
		return false
	}
	if col-1 >= 0 && nums[row][col-1] <= nums[row][col] {
		return false
	}
	if col+1 < len(nums[row]) && nums[row][col+1] <= nums[row][col] {
		return false
	}
	return true
}

func Part1(nums [][]int) int {
	risk_sum := 0
	for row := range nums {
		for col := range nums[row] {
			if IsLowPoint(row, col, nums) {
				risk_sum += nums[row][col] + 1
			}
		}
	}
	return risk_sum
}

func IsPointValid(point Point, nums [][]int) bool {
	return point.row >= 0 && point.row < len(nums) && point.col >= 0 && point.col < len(nums[point.row])
}

func FindBasinSize(row int, col int, nums [][]int) int {
	queue := make([]Point, 0)
	queue = append(queue, Point{row, col})
	basin_size := 0
	seen := make(map[Point]bool)
	for len(queue) != 0 {
		cur_point := queue[0]
		queue = queue[1:]
		if _, ok := seen[cur_point]; ok {
			continue
		}
		seen[cur_point] = true
		if IsPointValid(cur_point, nums) && nums[cur_point.row][cur_point.col] != 9 {
			basin_size++
			queue = append(queue, Point{cur_point.row - 1, cur_point.col})
			queue = append(queue, Point{cur_point.row + 1, cur_point.col})
			queue = append(queue, Point{cur_point.row, cur_point.col - 1})
			queue = append(queue, Point{cur_point.row, cur_point.col + 1})
		}
	}
	return basin_size
}

func Part2(nums [][]int) int {
	max_basin_sizes := []int{math.MinInt32, math.MinInt32, math.MinInt32}
	for row := range nums {
		for col := range nums[row] {
			if IsLowPoint(row, col, nums) {
				basin_size := FindBasinSize(row, col, nums)
				if basin_size > max_basin_sizes[0] {
					max_basin_sizes[2] = max_basin_sizes[1]
					max_basin_sizes[1] = max_basin_sizes[0]
					max_basin_sizes[0] = basin_size
				} else if basin_size > max_basin_sizes[1] {
					max_basin_sizes[2] = max_basin_sizes[1]
					max_basin_sizes[1] = basin_size
				} else if basin_size > max_basin_sizes[2] {
					max_basin_sizes[2] = basin_size
				}
			}
		}
	}
	return max_basin_sizes[0] * max_basin_sizes[1] * max_basin_sizes[2]
}

func main() {
	nums := ReadInput("input.txt")
	fmt.Println(Part1(nums))
	fmt.Println(Part2(nums))
}
