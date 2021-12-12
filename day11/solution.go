package main

import (
	"bufio"
	"fmt"
	"log"
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

	var grid [][]int
	for scanner.Scan() {
		cur_line := scanner.Text()
		cur_row := []int{}
		for _, num_str := range cur_line {
			num, err := strconv.Atoi(string(num_str))
			if err != nil {
				log.Fatalf("Could not convert string: %s to int", string(num_str))
			}
			cur_row = append(cur_row, num)
		}
		grid = append(grid, cur_row)
	}
	return grid
}

func GetNeighbours(cur_point Point) []Point {
	neighbours := []Point{}
	row := cur_point.row
	col := cur_point.col
	neighbours = append(neighbours, Point{row - 1, col - 1})
	neighbours = append(neighbours, Point{row - 1, col})
	neighbours = append(neighbours, Point{row - 1, col + 1})
	neighbours = append(neighbours, Point{row, col - 1})
	neighbours = append(neighbours, Point{row, col + 1})
	neighbours = append(neighbours, Point{row + 1, col - 1})
	neighbours = append(neighbours, Point{row + 1, col})
	neighbours = append(neighbours, Point{row + 1, col + 1})

	return neighbours
}

func IsValid(row int, col int, grid [][]int) bool {
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[row])
}

func IncreaseGridValues(grid [][]int) [][]int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j]++
		}
	}
	return grid
}

func PerformFlashes(row int, col int, grid [][]int, flashed map[Point]bool) ([][]int, int) {
	queue := []Point{}
	queue = append(queue, Point{row, col})
	num_flashes := 0

	for len(queue) != 0 {
		cur_point := queue[0]
		queue = queue[1:]

		if _, ok := flashed[cur_point]; ok {
			continue
		}
		flashed[cur_point] = true
		num_flashes++
		neighbours := GetNeighbours(cur_point)

		for _, n := range neighbours {
			if IsValid(n.row, n.col, grid) {
				grid[n.row][n.col] += 1
				if grid[n.row][n.col] > 9 {
					queue = append(queue, Point{n.row, n.col})
				}
			}
		}
	}
	return grid, num_flashes
}

func Part1(grid [][]int, total_steps int) int {
	num_flashes := 0
	total_rows := len(grid)
	total_cols := len(grid[0])
	for cur_step := 0; cur_step < total_steps; cur_step++ {
		// if cur_step%10 == 0 {
		// }
		grid = IncreaseGridValues(grid)
		// for _, cur_row := range grid {
		// 	fmt.Println(cur_row)
		// }
		// fmt.Println("_________________________")
		flashed := make(map[Point]bool)
		for i := 0; i < total_rows; i++ {
			for j := 0; j < total_cols; j++ {
				if grid[i][j] >= 10 {
					var cur_flashes int
					grid, cur_flashes = PerformFlashes(i, j, grid, flashed)
					// fmt.Printf("\tNum flashes: %d at row: %d and col: %d\n", cur_flashes, i, j)
					num_flashes += cur_flashes
				}
			}
		}
		// fmt.Printf("Num flashes: %d at step: %d\n", num_flashes, cur_step)
		for i := 0; i < total_rows; i++ {
			for j := 0; j < total_cols; j++ {
				if grid[i][j] > 9 {
					grid[i][j] = 0
				}
			}
		}
	}
	return num_flashes
}

func Part2(grid [][]int) int {
	num_flashes := 0
	total_rows := len(grid)
	total_cols := len(grid[0])
	cur_step := 0
	for true {
		// if cur_step%10 == 0 {
		// }
		grid = IncreaseGridValues(grid)
		// for _, cur_row := range grid {
		// 	fmt.Println(cur_row)
		// }
		// fmt.Println("_________________________")
		flashed := make(map[Point]bool)
		for i := 0; i < total_rows; i++ {
			for j := 0; j < total_cols; j++ {
				if grid[i][j] >= 10 {
					var cur_flashes int
					grid, cur_flashes = PerformFlashes(i, j, grid, flashed)
					// fmt.Printf("\tNum flashes: %d at row: %d and col: %d\n", cur_flashes, i, j)
					num_flashes += cur_flashes
				}
			}
		}
		// fmt.Printf("Num flashes: %d at step: %d\n", num_flashes, cur_step)
		for i := 0; i < total_rows; i++ {
			for j := 0; j < total_cols; j++ {
				if grid[i][j] > 9 {
					grid[i][j] = 0
				}
			}
		}
		cur_step++
		if len(flashed) == 100 {
			fmt.Println(grid)
			return cur_step
		}
	}
	return -1
}

func main() {
	grid1 := ReadInput("input.txt")
	grid2 := ReadInput("input.txt")
	fmt.Println(Part1(grid1, 100))
	fmt.Println(Part2(grid2))
}
