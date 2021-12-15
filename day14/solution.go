package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Data struct {
	template string
	rules    map[string]string
}

type PairAndSteps struct {
	pair  string
	steps int
}

func ReadInput(file_name string) Data {
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatalf("Could not open file: %s", file_name)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var data Data
	rules := make(map[string]string)
	i := 0

	for scanner.Scan() {
		cur_line := scanner.Text()
		if i == 0 {
			data.template = cur_line
			i++
			continue
		}
		if len(cur_line) == 0 {
			continue
		}
		cur_line_split := strings.Split(cur_line, " -> ")
		rules[cur_line_split[0]] = cur_line_split[1]
	}
	data.rules = rules
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

func FindMaxElementMinusMinElement(template string) int {
	counts := make(map[string]int)

	for _, c := range template {
		counts[string(c)]++
	}
	max_count := math.MinInt32
	min_count := math.MaxInt32

	for _, v := range counts {
		max_count = max(max_count, v)
		min_count = min(min_count, v)
	}
	return max_count - min_count
}

func Part1(data Data, num_steps int) int {
	template := data.template
	rules := data.rules

	for i := 0; i < num_steps; i++ {
		last_idx := 0
		new_template := ""
		for j := 0; j < len(template)-1; j++ {
			if val, ok := rules[string(template[j])+string(template[j+1])]; ok {
				new_value := ""
				if last_idx != 0 && last_idx == j {
					new_value = val + string(template[j+1])
					last_idx = j + 1
				} else {
					new_value = string(template[j]) + val + string(template[j+1])
					last_idx = j + 1
				}
				new_template += new_value
			}
		}
		template = new_template
	}

	return FindMaxElementMinusMinElement(template)
}

func MergeCounts(m1, m2 map[string]int) map[string]int {
	for ia, va := range m1 {
		m2[ia] += va
	}
	return m2
}

func CopyMap(old map[string]int) map[string]int {
	new_map := make(map[string]int)
	for k, v := range old {
		new_map[k] = v
	}
	return new_map
}

func Part2Helper(num_steps int, cur_pair string, rules map[string]string, seen map[PairAndSteps]map[string]int) map[string]int {
	if val, ok := seen[PairAndSteps{cur_pair, num_steps}]; ok {
		tmp_val := CopyMap(val)
		return tmp_val
	}
	val, ok := rules[cur_pair]
	cur_count := make(map[string]int)
	if ok {
		cur_count[val]++
		if num_steps != 1 {
			new_count := MergeCounts(Part2Helper(num_steps-1, string(cur_pair[0])+val, rules, seen),
				Part2Helper(num_steps-1, val+string(cur_pair[1]), rules, seen))

			new_count = MergeCounts(new_count, cur_count)
			tmp_copy := CopyMap(new_count)
			seen[PairAndSteps{cur_pair, num_steps}] = tmp_copy
			return new_count
		}
	}
	return cur_count
}

func Part2(data Data, num_steps int) int {
	starting_count := make(map[string]int)
	for _, c := range data.template {
		starting_count[string(c)]++
	}
	seen := make(map[PairAndSteps]map[string]int)
	final_count := make(map[string]int)
	for i := 0; i < len(data.template)-1; i++ {
		final_count = MergeCounts(final_count, Part2Helper(num_steps, string(data.template[i])+string(data.template[i+1]), data.rules, seen))
	}
	final_count = MergeCounts(final_count, starting_count)
	fmt.Println(final_count)
	max_count := math.MinInt64
	min_count := math.MaxInt64

	for _, v := range final_count {
		max_count = max(max_count, v)
		min_count = min(min_count, v)
	}
	return max_count - min_count
}

func main() {
	data := ReadInput("input.txt")
	fmt.Println(Part1(data, 10))
	fmt.Println(Part2(data, 40))
}
