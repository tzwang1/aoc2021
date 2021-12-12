package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadInput(file_name string) map[string][]string {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatalf("Could not open file: %s", file_name)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	graph := make(map[string][]string)
	for scanner.Scan() {
		cur_line := scanner.Text()
		rooms := strings.Split(cur_line, "-")
		graph[rooms[0]] = append(graph[rooms[0]], rooms[1])
		graph[rooms[1]] = append(graph[rooms[1]], rooms[0])
	}
	return graph
}

func FindNumPathsPart1Helper(cur_room string, graph map[string][]string, seen map[string]bool) int {
	if cur_room == "end" {
		return 1
	}
	if val, ok := seen[cur_room]; ok && val {
		return 0
	}
	if cur_room == strings.ToLower(cur_room) {
		seen[cur_room] = true
	}
	num_paths := 0
	for _, next_room := range graph[cur_room] {
		num_paths += FindNumPathsPart1Helper(next_room, graph, seen)
	}
	seen[cur_room] = false
	return num_paths
}

func FindNumPathsPart1(graph map[string][]string) int {
	seen := make(map[string]bool)
	return FindNumPathsPart1Helper("start", graph, seen)

}

func FindNumPathsPart2Helper(cur_room string, graph map[string][]string, seen map[string]int, visited_twice bool) int {
	// fmt.Println("Cur_room: ", cur_room)
	// fmt.Println("Seen: ", seen)
	if cur_room == "end" {
		return 1
	}
	if val, ok := seen[cur_room]; ok {
		if cur_room == "start" {
			return 0
		}
		if (val == 1 && visited_twice) || val == 2 {
			return 0
		}
		if val == 1 && !visited_twice {
			visited_twice = true
		}
	}
	if cur_room == strings.ToLower(cur_room) {
		seen[cur_room] += 1
	}
	num_paths := 0
	for _, next_room := range graph[cur_room] {
		num_paths += FindNumPathsPart2Helper(next_room, graph, seen, visited_twice)
	}
	seen[cur_room] -= 1
	return num_paths
}

func FindNumPathsPart2(graph map[string][]string) int {
	seen := make(map[string]int)
	visited_twice := false
	return FindNumPathsPart2Helper("start", graph, seen, visited_twice)
}

func main() {
	graph1 := ReadInput("input.txt")
	fmt.Println(FindNumPathsPart1(graph1))
	graph2 := ReadInput("input.txt")
	fmt.Println(FindNumPathsPart2(graph2))
}
