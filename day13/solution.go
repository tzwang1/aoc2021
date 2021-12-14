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

type Point struct {
	x int
	y int
}

type Fold struct {
	axis  string
	value int
}

type Data struct {
	data  []Point
	folds []Fold
}

func ReadInput(file_name string) Data {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatalf("Could not open file: %s", file_name)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var data Data
	// Read coords
	for scanner.Scan() {
		cur_line := scanner.Text()
		if len(cur_line) == 0 {
			break
		}
		cur_line_split := strings.Split(cur_line, ",")
		x, err := strconv.Atoi(cur_line_split[0])
		if err != nil {
			log.Fatalf("Could not convert string: %s to int", x)
		}
		y, err := strconv.Atoi(cur_line_split[1])
		if err != nil {
			log.Fatalf("Could not convert string: %s to int", cur_line_split[1])
		}
		data.data = append(data.data, Point{x, y})
	}

	// Read folds
	for scanner.Scan() {
		cur_line := scanner.Text()
		cur_line_split := strings.Split(cur_line, " ")
		fold := strings.Split(cur_line_split[2], "=")
		value, err := strconv.Atoi(fold[1])
		if err != nil {
			log.Fatalf("Could not convert string: %s to int", fold[1])
		}
		data.folds = append(data.folds, Fold{fold[0], value})
	}
	return data
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Part1(data Data) int {
	point_map := make(map[Point]bool)
	for _, point := range data.data {
		point_map[point] = true
	}

	first_fold := data.folds[0]
	for point := range point_map {
		if first_fold.axis == "x" {
			if point.x < first_fold.value {
				new_x := first_fold.value + first_fold.value - point.x
				delete(point_map, point)
				point_map[Point{new_x, point.y}] = true
			}
		} else {
			if point.y > first_fold.value {
				new_y := first_fold.value + first_fold.value - point.y
				delete(point_map, point)
				point_map[Point{point.x, new_y}] = true
			}
		}
	}
	num_points := 0
	for _, value := range point_map {
		if value {
			num_points++
		}
	}
	return num_points
}

func WriteGridToFile(grid [][]string) {
	f, err := os.Create("part2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, grid_row := range grid {
		cur_row := ""
		for _, c := range grid_row {
			if string(c) == "#" {
				cur_row += string(c)
			} else {
				cur_row += " "
			}
		}
		fmt.Fprintln(f, cur_row)
	}
	f.Close()
}

func Part2(data Data) {
	point_map := make(map[Point]bool)
	for _, point := range data.data {
		point_map[point] = true
	}

	for _, fold := range data.folds {
		for point := range point_map {
			if fold.axis == "x" {
				if point.x > fold.value {
					new_x := fold.value - point.x + fold.value
					delete(point_map, point)
					point_map[Point{new_x, point.y}] = true
				}
			} else {
				if point.y > fold.value {
					new_y := fold.value - point.y + fold.value
					delete(point_map, point)
					point_map[Point{point.x, new_y}] = true
				}
			}
		}
	}

	max_x, max_y := 0, 0
	min_x, min_y := math.MaxInt32, math.MaxInt32
	for point, value := range point_map {
		if value {
			max_x = max(point.x, max_x)
			max_y = max(point.y, max_y)
			min_x = min(point.x, min_x)
			min_y = min(point.y, min_y)
		}
	}

	grid := make([][]string, max_y+1)
	for i := 0; i < max_y+1; i++ {
		grid[i] = make([]string, max_x+1)
	}
	for point := range point_map {
		grid[point.y][point.x] = "#"
	}
	WriteGridToFile(grid)
}

func main() {
	data1 := ReadInput("small_input.txt")
	data2 := ReadInput("input.txt")
	fmt.Println(Part1(data1))
	Part2(data2)
}
