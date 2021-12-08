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

type Segment struct {
	start []int
	end   []int
}

type Point struct {
	x int
	y int
}

func convertStringToInt(num_str string) int {
	num, err := strconv.Atoi(num_str)
	if err != nil {
		log.Fatalf("Could not convert string: %s to int", num_str)
	}
	return num
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

func ReadInput(file_name string) []Segment {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatalf("Could not open file: %s", file_name)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var segments []Segment
	for scanner.Scan() {
		cur_line := scanner.Text()
		cur_line_split := strings.Split(cur_line, " -> ")
		start_pos := strings.Split(cur_line_split[0], ",")
		end_pos := strings.Split(cur_line_split[1], ",")
		var cur_segment Segment
		cur_segment.start = append(cur_segment.start, convertStringToInt(start_pos[0]))
		cur_segment.start = append(cur_segment.start, convertStringToInt(start_pos[1]))
		cur_segment.end = append(cur_segment.end, convertStringToInt(end_pos[0]))
		cur_segment.end = append(cur_segment.end, convertStringToInt(end_pos[1]))
		segments = append(segments, cur_segment)
	}
	return segments
}

func Part1(segments []Segment) int {
	num_overlap := 0
	grid := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]int, 1000)
	}
	for _, segment := range segments {
		if segment.start[0] == segment.end[0] {
			i := segment.start[0]
			for j := min(segment.start[1], segment.end[1]); j <= max(segment.start[1], segment.end[1]); j++ {
				grid[i][j] += 1
			}
		}
		if segment.start[1] == segment.end[1] {
			j := segment.start[1]
			for i := min(segment.start[0], segment.end[0]); i <= max(segment.start[0], segment.end[0]); i++ {
				grid[i][j] += 1
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] >= 2 {
				num_overlap++
			}
		}
	}
	return num_overlap
}

func Is45Diagonal(start, end []int) bool {
	return math.Abs(float64(start[0]-end[0])) == math.Abs(float64(start[1])-float64(end[1]))
}

func minPointX(p1, p2 []int) []int {
	if p1[0] < p2[0] {
		return p1
	}
	return p2
}

func maxPointX(p1, p2 []int) []int {
	if p1[0] > p2[0] {
		return p1
	}
	return p2
}

func CalculatePointsConvered(p1, p2 []int) []Point {
	points := []Point{}
	start := minPointX(p1, p2)
	end := maxPointX(p1, p2)
	if start[1] < end[1] {
		for i := 0; start[0]+i <= end[0]; i++ {
			points = append(points, Point{start[0] + i, start[1] + i})
		}
	} else if start[1] > end[1] {
		for i := 0; start[0]+i <= end[0]; i++ {
			points = append(points, Point{start[0] + i, start[1] - i})
		}
	}
	return points
}

func Part2(segments []Segment) int {
	num_overlap := 0
	grid := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]int, 1000)
	}
	for _, segment := range segments {
		if segment.start[0] == segment.end[0] {
			i := segment.start[0]
			for j := min(segment.start[1], segment.end[1]); j <= max(segment.start[1], segment.end[1]); j++ {
				grid[i][j] += 1
			}
		}
		if segment.start[1] == segment.end[1] {
			j := segment.start[1]
			for i := min(segment.start[0], segment.end[0]); i <= max(segment.start[0], segment.end[0]); i++ {
				grid[i][j] += 1
			}
		}

		if Is45Diagonal(segment.start, segment.end) {
			points := CalculatePointsConvered(segment.start, segment.end)
			for _, point := range points {
				grid[point.x][point.y]++
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] >= 2 {
				num_overlap++
			}
		}
	}
	return num_overlap
}

func main() {
	segments := ReadInput("input.txt")
	fmt.Println(Part1(segments))
	fmt.Println(Part2(segments))
	p1 := []int{9, 7}
	p2 := []int{5, 3}
	fmt.Println(CalculatePointsConvered(p1, p2))
}
