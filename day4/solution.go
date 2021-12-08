package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	board [][]int
}

func ReadInput(filename string) ([]int, []Board) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Could not open file")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	nums_str := ""
	var boards []Board
	var cur_board Board
	var board_row []int
	for scanner.Scan() {
		cur_line := scanner.Text()
		if len(nums_str) == 0 {
			nums_str = cur_line
		} else {
			if len(cur_line) != 0 {
				cur_line_split := strings.Split(cur_line, " ")
				for _, cur_num_str := range cur_line_split {
					if len(cur_num_str) == 0 {
						continue
					}
					cur_num, err := strconv.Atoi(cur_num_str)
					if err != nil {
						log.Fatalf("Could not convert string: %s to int", cur_num_str)
					}
					board_row = append(board_row, cur_num)
				}
				cur_board.board = append(cur_board.board, board_row)
				board_row = []int{}
			} else {
				if len(cur_board.board) != 0 {
					boards = append(boards, cur_board)
					cur_board = Board{}
				}
			}
		}
	}
	nums_str_split := strings.Split(nums_str, ",")
	var nums []int
	for _, num_str := range nums_str_split {
		num, err := strconv.Atoi(num_str)
		if err != nil {
			log.Fatalf("Could not convert string: %s to int", num_str)
		}
		nums = append(nums, num)
	}
	return nums, boards
}

func Part1(nums []int, boards []Board) int {
	for _, num := range nums {
		for _, board := range boards {
			for i := 0; i < len(board.board); i++ {
				for j := 0; j < len(board.board[i]); j++ {
					if board.board[i][j] == num && board.board[i][j] > 0 {
						board.board[i][j] *= -1
						bingo_row := true
						bingo_col := true
						for row := 0; row < len(board.board); row++ {
							if board.board[row][j] > 0 {
								bingo_row = false
								break
							}
						}
						for col := 0; col < len(board.board[i]); col++ {
							if board.board[i][col] > 0 {
								bingo_col = false
								break
							}
						}
						if bingo_row || bingo_col {
							return calculateScore(num, board.board)
						}
					}
				}
			}
		}
	}
	return 0
}

func Part2(nums []int, boards []Board) int {
	winning_boards := make(map[int]int)
	var winning_board_order []int
	for _, num := range nums {
		for board_idx, board := range boards {
			for i := 0; i < len(board.board); i++ {
				for j := 0; j < len(board.board[i]); j++ {
					if board.board[i][j] == num && board.board[i][j] > 0 {
						board.board[i][j] *= -1
						bingo_row := true
						bingo_col := true
						for row := 0; row < len(board.board); row++ {
							if board.board[row][j] > 0 {
								bingo_row = false
								break
							}
						}
						for col := 0; col < len(board.board[i]); col++ {
							if board.board[i][col] > 0 {
								bingo_col = false
								break
							}
						}
						if bingo_row || bingo_col {
							score := calculateScore(num, board.board)
							if _, ok := winning_boards[board_idx]; !ok {
								winning_boards[board_idx] = score
								winning_board_order = append(winning_board_order, score)
							}
						}
					}
				}
			}
		}
	}
	return winning_board_order[len(winning_board_order)-1]
}

func calculateScore(num int, board [][]int) int {
	boardSum := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] > 0 {
				boardSum += board[i][j]
			}
		}
	}
	return boardSum * num
}

func main() {
	nums, boards := ReadInput("input.txt")
	fmt.Println(Part1(nums, boards))
	fmt.Println(Part2(nums, boards))
}
