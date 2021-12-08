package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	command string
	value   int
}

func ReadInput() []Command {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Could not open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var commands []Command
	for scanner.Scan() {
		cur_line := scanner.Text()
		command_and_num := strings.Split(cur_line, " ")
		num, err := strconv.Atoi(command_and_num[1])
		if err != nil {
			log.Fatalf("Could not convert input line: %s to int", cur_line)
		}
		commands = append(commands, Command{command_and_num[0], num})
	}
	return commands
}

func Part1(commands []Command) int {
	x, y := 0, 0
	for _, command := range commands {
		if command.command == "forward" {
			x += command.value
		} else if command.command == "up" {
			y -= command.value
		} else {
			y += command.value
		}
	}
	return x * y
}

func Part2(commands []Command) int {
	horizontal, depth, aim := 0, 0, 0

	for _, command := range commands {
		if command.command == "forward" {
			horizontal += command.value
			depth += aim * command.value
		} else if command.command == "up" {
			aim -= command.value
		} else {
			aim += command.value
		}
	}
	return horizontal * depth
}

func main() {
	commands := ReadInput()
	fmt.Println(Part1(commands))
	fmt.Println(Part2(commands))

}
