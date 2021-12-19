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

type TargetArea struct {
	x_min int
	x_max int
	y_min int
	y_max int
}

type Result struct {
	hits_target bool
	max_y       int
}

func ReadInput(file_name string) TargetArea {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatalf("Could not open file: %s", file_name)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	cur_line := scanner.Text()
	cur_line_split := strings.Split(cur_line, " ")
	var target_area TargetArea
	x_values := strings.Split(strings.TrimRight(strings.Split(cur_line_split[2], "=")[1], ","), "..")
	y_values := strings.Split(strings.Split(cur_line_split[3], "=")[1], "..")
	x_min, err := strconv.Atoi(x_values[0])
	if err != nil {
		log.Fatalf("Could not convert x_min string: %s to int\n", x_values[0])
	}
	target_area.x_min = x_min

	x_max, err := strconv.Atoi(x_values[1])
	if err != nil {
		log.Fatalf("Could not convert x_max string: %s to int\n", x_values[1])
	}
	target_area.x_max = x_max

	y_min, err := strconv.Atoi(y_values[0])
	if err != nil {
		log.Fatalf("Could not convert y_min string: %s to int\n", y_values[0])
	}
	target_area.y_min = y_min

	y_max, err := strconv.Atoi(y_values[1])
	if err != nil {
		log.Fatalf("Could not convert y_max string: %s to int\n", y_values[1])
	}
	target_area.y_max = y_max

	return target_area
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TrajectoryHitsTarget(x int, y int, target_area TargetArea) Result {
	// fmt.Printf("Testing if initial velocity of x: %d and y: %d hits target\n", x, y)
	cur_x, cur_y := 0, 0
	var result Result
	result.hits_target = false
	max_y := math.MinInt64
	for {
		cur_x += x
		cur_y += y

		// fmt.Println("Cur_x: ", cur_x)
		// fmt.Println("Cur_y: ", cur_y)

		if x == 0 && cur_y < target_area.y_min {
			break
		}

		if x > 0 {
			x--
		} else if x < 0 {
			x++
		}
		y--
		if cur_x >= target_area.x_min && cur_x <= target_area.x_max && cur_y >= target_area.y_min && cur_y <= target_area.y_max {
			// fmt.Println("Hit target!")
			result.hits_target = true
		}
		max_y = Max(max_y, cur_y)
	}
	result.max_y = max_y
	// if result.hits_target {
	// 	fmt.Println("Hit target!")
	// }
	return result
}

func Part1(target_area TargetArea) int {
	y := target_area.y_min
	max_y := 0
	for y < 1000 {
		for x := 0; x <= target_area.x_max; x++ {
			result := TrajectoryHitsTarget(x, y, target_area)
			if result.hits_target {
				max_y = result.max_y
			}
		}
		y++
	}
	return max_y
}

func Part2(target_area TargetArea) int {
	y := target_area.y_min
	num_hits := 0
	for y < 1000 {
		for x := target_area.y_min; x <= target_area.x_max; x++ {
			result := TrajectoryHitsTarget(x, y, target_area)
			if result.hits_target {
				num_hits++
			}
		}
		y++
	}
	return num_hits
}

func main() {
	target_area := ReadInput("input.txt")
	fmt.Println(target_area)
	fmt.Println(Part1(target_area))
	fmt.Println(Part2(target_area))
}
