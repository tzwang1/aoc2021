package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type PointWithCost struct {
	i    int
	j    int
	cost int
}

type Point struct {
	i int
	j int
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

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func IsValid(point Point, grid [][]int) bool {
	return point.i >= 0 && point.i < len(grid) && point.j >= 0 && point.j < len(grid[point.i])
}

func GetNeighbours(point PointWithCost) []Point {
	neighbours := []Point{}
	neighbours = append(neighbours, Point{point.i - 1, point.j})
	neighbours = append(neighbours, Point{point.i, point.j - 1})
	neighbours = append(neighbours, Point{point.i + 1, point.j})
	neighbours = append(neighbours, Point{point.i, point.j + 1})
	return neighbours
}

func Part1(grid [][]int) int {
	return FindShortedPath(grid)
}

func FindShortedPath(grid [][]int) int {
	queue := []PointWithCost{}
	queue = append(queue, PointWithCost{0, 0, 0})
	visited := make(map[Point]bool)
	min_dist := make(map[Point]int)

	for len(queue) != 0 {
		cur_point := queue[0]
		// fmt.Println(queue)
		queue = queue[1:]
		cur_point_without_cost := Point{cur_point.i, cur_point.j}
		if _, ok := visited[cur_point_without_cost]; ok {
			continue
		}
		cur_neighbours_min_cost_order := []PointWithCost{}
		for _, neighbour := range GetNeighbours(cur_point) {
			if IsValid(neighbour, grid) {
				if _, ok := visited[neighbour]; ok {
					continue
				}
				val, ok := min_dist[neighbour]
				if !ok {
					min_dist[neighbour] = min_dist[cur_point_without_cost] + grid[neighbour.i][neighbour.j]
				} else {
					if val < min_dist[cur_point_without_cost]+grid[neighbour.i][neighbour.j] {
						min_dist[neighbour] = val
					}
				}
				cur_neighbours_min_cost_order = append(cur_neighbours_min_cost_order, PointWithCost{neighbour.i, neighbour.j, min_dist[neighbour]})
			}
		}
		visited[cur_point_without_cost] = true
		queue = append(queue, cur_neighbours_min_cost_order...)
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].cost < queue[j].cost
		})
	}
	return min_dist[Point{len(grid) - 1, len(grid[0]) - 1}]
}

func ExpandGrid(grid [][]int) [][]int {
	num_rows := len(grid)
	num_cols := len(grid[0])
	full_grid := make([][]int, num_rows*5)
	for i := range full_grid {
		full_grid[i] = make([]int, num_cols*5)
	}

	for i := 0; i < len(full_grid); i++ {
		for j := 0; j < len(full_grid[i]); j++ {
			if i < num_rows && j < num_cols {
				full_grid[i][j] = grid[i][j]
			} else if i >= num_rows {
				new_value := full_grid[i-num_rows][j] + 1
				if new_value > 9 {
					new_value = 1
				}
				full_grid[i][j] = new_value
			} else if j >= num_cols {
				new_value := full_grid[i][j-num_cols] + 1
				if new_value > 9 {
					new_value = 1
				}
				full_grid[i][j] = new_value
			}
		}
	}
	return full_grid
}

func Part2(grid [][]int) int {
	full_grid := ExpandGrid(grid)
	return FindShortedPath(full_grid)
}

func FindGridDiff(grid1 [][]int, grid2 [][]int) {
	for i := 0; i < len(grid1); i++ {
		for j := 0; j < len(grid1[i]); j++ {
			if grid1[i][j] != grid2[i][j] {
				fmt.Printf("Grid1 and Grid2 are not equal at row: %d and col: %d.\n\tGrid1 has value: %d\n\tGrid2 has value: %d\n", i, j, grid1[i][j], grid2[i][j])
			}
		}
	}
}

func main() {
	grid := ReadInput("input.txt")
	fmt.Println(Part1(grid))
	fmt.Println(Part2(grid))
}
